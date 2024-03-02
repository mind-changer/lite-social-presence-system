package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/lite-social-presence-system/internal/def"
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
	err := f.conn.QueryRow(ctx, query, userId1, userId2).Scan(&exists)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if friend")
		return false, err
	}
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
		logrus.WithError(err).Error("Error while checking if friend exists")
		return err
	}
	if !userExists {
		logrus.WithError(err).Error("User doesnt exist")
		return def.CreateClientError(400, "user doesnt exist")
	}
	if !friendExists {
		logrus.WithError(err).Error("Friend doesnt exist")
		return def.CreateClientError(400, "friend doesnt exist")
	}
	friendReqExists, err := f.db.GetFriendRequestsTable(ctx).FriendRequestExists(ctx, friendId, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if friend req exists")
		return err
	}
	if !friendReqExists {
		logrus.WithError(err).Error("Friend request doesnt exist")
		return def.CreateClientError(400, "friend req doesnt exist")
	}
	exists, err := f.IsFriend(ctx, userId, friendId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if friend  exists")
		return err
	}
	if exists {
		logrus.WithError(err).Error("Friend already exists")
		return def.CreateClientError(409, "friend already exists")
	}
	insertSql := `
	insert into friends(user_id,friend_id)
	values
	($1,$2),
	($2,$1);
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
		return def.CreateClientError(400, "user doesnt exist")
	}
	friendExists, err := usersTable.UserExists(ctx, friendId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if friend exists")
		return err
	}
	if !friendExists {
		logrus.WithError(err).Error("Friend doesnt exist")
		return def.CreateClientError(400, "friend doesnt exist")
	}
	deleteSql := `
	DELETE FROM friends
    WHERE (user_id=$1 AND friend_id=$2) OR (user_id=$2 AND friend_id=$1);
	`
	_, err = f.conn.Exec(ctx, deleteSql, userId, friendId)
	if err != nil {
		logrus.WithError(err).Error("Error while deleting friend")
		return err
	}
	return nil
}
