package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type FriendRequests interface {
	SendFriendRequest(ctx context.Context, userId, requesterId string) error
	AcceptFriendRequest(ctx context.Context, userId, requesterId string) error
	RejectFriendRequest(ctx context.Context, userId, requesterId string) error
	FriendRequestExists(ctx context.Context, userId, requesterId string) (bool, error)
}

type friendRequests struct {
	*db
	conn *pgx.Conn
}

var friendRequestsMutex sync.Mutex
var friendRequestsObject *friendRequests

func (f *friendRequests) SendFriendRequest(ctx context.Context, userId, requesterId string) error {
	usersTable := f.db.GetUsersTable(ctx)
	userExists, err := usersTable.UserExists(ctx, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	requesterExists, err := usersTable.UserExists(ctx, requesterId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !userExists {
		logrus.WithError(err).Error("User doesnt exist")
		return fmt.Errorf("user doesnt exist")
	}
	if !requesterExists {
		logrus.WithError(err).Error("Friend doesnt exist")
		return fmt.Errorf("user doesnt exist")
	}
	insertSql := `
	insert into friend_requests(user_id,requester_id) 
	values($1,$2);
	`
	_, err = f.conn.Exec(ctx, insertSql, userId, requesterId)
	if err != nil {
		logrus.WithError(err).Error("Error while inserting friend")
		return err
	}
	return nil
}

func (f *friendRequests) FriendRequestExists(ctx context.Context, userId, requesterId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM friend_requests WHERE user_id=$1 AND reqeuster_id=$2)`
	exists := false
	err := f.conn.QueryRow(ctx, query, userId, requesterId).Scan(&exists)
	if err != nil {
		logrus.WithError(err).Error("Error while getting user friends")
		return false, err
	}
	logrus.Info("query,id ", query, userId)
	logrus.Info(exists)
	return exists, nil
}

func (f *friendRequests) AcceptFriendRequest(ctx context.Context, userId, requesterId string) error {
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
	requesterExists, err := usersTable.UserExists(ctx, requesterId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !requesterExists {
		logrus.WithError(err).Error("Friend doesnt exist")
		return fmt.Errorf("user doesnt exist")
	}
	friendReqExists, err := f.db.GetFriendRequestsTable(ctx).FriendRequestExists(ctx, userId, requesterId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !friendReqExists {
		logrus.WithError(err).Error("Friend request doesnt exist")
		return fmt.Errorf("user doesnt exist")
	}
	insertSql := `
	insert into friends(user_id,friend_id) 
	values($1,$2),
	values($2,$1);
	`
	_, err = f.conn.Exec(ctx, insertSql, userId, requesterId)
	if err != nil {
		logrus.WithError(err).Error("Error while inserting friend")
		return err
	}
	return nil
}

func (f *friendRequests) RejectFriendRequest(ctx context.Context, userId, requesterId string) error {

	friendReqExists, err := f.db.GetFriendRequestsTable(ctx).FriendRequestExists(ctx, userId, requesterId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !friendReqExists {
		logrus.WithError(err).Error("Friend request doesnt exist")
		return fmt.Errorf("user doesnt exist")
	}
	insertSql := `
	DELETE from friend_requests
	WHERE user_id=$1 AND requester_id=$2;
	`
	_, err = f.conn.Exec(ctx, insertSql, userId, requesterId)
	if err != nil {
		logrus.WithError(err).Error("Error while inserting friend")
		return err
	}
	return nil
}
