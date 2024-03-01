package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Friends interface {
	GetFriends(ctx context.Context, userId string) ([]string, error)
	AddFriend(ctx context.Context, userId, friendId string) error
	RemoveFriend(ctx context.Context, userId, friendId string) error
	IsFriend(ctx context.Context, userId1, userId2 string) (bool, error)
}

type friends struct {
	*db
	conn *pgx.Conn
}

var friendsMutex sync.Mutex
var friendsObject *friends

func (f *friends) GetFriends(ctx context.Context, userId string) ([]string, error) {
	query := fmt.Sprintf(`select friend_id from friends where user_id='%s'`, userId)
	logrus.Info("query ", query)
	rows, err := f.conn.Query(ctx, query)
	if err != nil {
		logrus.WithError(err).Error("Error while getting user friends")
		return nil, err
	}
	friends := make([]string, 0)
	for rows.Next() {
		friendId := ""
		if err := rows.Scan(&friendId); err != nil {
			logrus.WithError(err).Error("Error while decoding friends")
			return nil, err
		}
		friends = append(friends, friendId)
	}
	logrus.Info(friends)
	defer rows.Close()
	return friends, nil
}

func (f *friends) IsFriend(ctx context.Context, userId1, userId2 string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM friends WHERE user_id=$1 AND friend_id=$2)`
	exists := false
	err := f.conn.QueryRow(ctx, userId1, userId2).Scan(&exists)
	if err != nil {
		logrus.WithError(err).Error("Error while getting user friends")
		return false, err
	}
	logrus.Info("query,id1,id2 ", query, userId1, userId2)
	logrus.Info(exists)
	return exists, nil
}

func (f *friends) AddFriend(ctx context.Context, userId, friendId string) error {
	usersTable := f.db.GetUsersTable(ctx)
	userExists, err := usersTable.UserExists(ctx, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	friendExists, err := usersTable.UserExists(ctx, friendId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !userExists {
		logrus.WithError(err).Error("User doesnt exist")
		return fmt.Errorf("user doesnt exist")
	}
	if !friendExists {
		logrus.WithError(err).Error("Friend doesnt exist")
		return fmt.Errorf("user doesnt exist")
	}
	insertSql := `
	insert into friends(user_id,friend_id) 
	values($1,$2),
	values($2,$1);
	`
	_, err = f.conn.Exec(ctx, insertSql, userId, friendId)
	if err != nil {
		logrus.WithError(err).Error("Error while inserting friend")
		return err
	}
	return nil
}

func (f *friends) RemoveFriend(ctx context.Context, userId, friendId string) error {
	usersTable := f.db.GetUsersTable(ctx)
	userExists, err := usersTable.UserExists(ctx, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !userExists {
		logrus.WithError(err).Error("User doesnt exist")
		return fmt.Errorf("user doesnt exist")
	}
	friendExists, err := usersTable.UserExists(ctx, friendId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !friendExists {
		logrus.WithError(err).Error("Friend doesnt exist")
		return fmt.Errorf("user doesnt exist")
	}
	deleteSql := `
	DELETE FROM friends
    WHERE (user_id=$1 AND friend_id=$2) OR (user_id=$2 AND friend_id=$1);
	`
	_, err = f.conn.Exec(ctx, deleteSql, userId, friendId)
	if err != nil {
		logrus.WithError(err).Error("Error while inserting friend")
		return err
	}
	return nil
}
