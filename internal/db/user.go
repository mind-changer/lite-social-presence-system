package db

import (
	"sync"

	"github.com/jackc/pgx/v5"
)

type Users interface {
	// GetConn(ctx context.Context) (*pgx.Conn, error)
}

type users struct {
	conn *pgx.Conn
}

var usersMutex sync.Mutex
var usersObject *users
