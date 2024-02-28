package db

import (
	"sync"

	"github.com/jackc/pgx/v5"
)

type FriendRequests interface {
	// GetConn(ctx context.Context) (*pgx.Conn, error)
}

type friendRequests struct {
	conn *pgx.Conn
}

var friendRequestsMutex sync.Mutex
var friendRequestsObject *friendRequests
