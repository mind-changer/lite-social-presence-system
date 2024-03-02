package db

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/lite-social-presence-system/internal/def"
	"github.com/sirupsen/logrus"
)

type Users interface {
	GetUserStatus(ctx context.Context, userId string) (string, error)
	UpdateUserStatus(ctx context.Context, userId, status string) error
	UserExists(ctx context.Context, userId string) (bool, error)
}

type users struct {
	*db
	conn *pgx.Conn
}

var usersMutex sync.Mutex
var usersObject *users

func (u *users) GetUserStatus(ctx context.Context, userId string) (string, error) {
	query := "SELECT status FROM users WHERE id=$1"
	status := ""
	err := u.conn.QueryRow(ctx, "SELECT status FROM users WHERE id=$1", userId).Scan(&status)
	if err != nil {
		logrus.WithError(err).Error("Error while getting user status")
		return "", err
	}
	logrus.Info("query,id ", query, userId)
	return status, nil
}

func (u *users) UserExists(ctx context.Context, userId string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)"
	exists := false
	err := u.conn.QueryRow(ctx, query, userId).Scan(&exists)
	if err != nil {
		logrus.WithError(err).Error("Error while getting user friends")
		return false, err
	}
	logrus.Info("query,id ", query, userId)
	logrus.Info(exists)
	return exists, nil
}

func (u *users) UpdateUserStatus(ctx context.Context, userId, status string) error {
	userExists, err := u.db.GetUsersTable(ctx).UserExists(ctx, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !userExists {
		logrus.WithError(err).Error("User doesnt exist")
		return def.CreateClientError(400, "user doesnt exist")
	}

	updateSql := "UPDATE users SET status=$1 WHERE id=$2"
	if _, err := u.conn.Exec(ctx, updateSql, status, userId); err != nil {
		logrus.WithError(err).Error("Error while inserting friend")
		return err
	}
	return nil
}
