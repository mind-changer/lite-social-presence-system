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
	ReadFriends(ctx context.Context, userId string) ([]def.Friend, error)
}

type friends struct {
	conn *pgx.Conn
}

var friendsMutex sync.Mutex
var friendsObject *friends

func (f *friends) ReadFriends(ctx context.Context, userId string) ([]def.Friend, error) {
	query := fmt.Sprintf(`select * from friends where user_id='%s'`, userId)
	logrus.Info("query ", query)
	rows, err := f.conn.Query(ctx, query)
	if err != nil {
		logrus.WithError(err).Error("Error while getting user friends")
		return nil, err
	}
	friends := make([]def.Friend, 0)
	for rows.Next() {
		row := def.Friend{}
		if err := rows.Scan(&row.A, &row.B); err != nil {
			logrus.WithError(err).Error("Error while decoding friends")
			return nil, err
		}
		friends = append(friends, row)
	}
	logrus.Info(friends)
	defer rows.Close()
	return friends, nil
}
