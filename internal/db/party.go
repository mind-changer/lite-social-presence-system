package db

import (
	"sync"

	"github.com/jackc/pgx/v5"
)

type Parties interface {
	// GetConn(ctx context.Context) (*pgx.Conn, error)
}

type parties struct {
	conn *pgx.Conn
}

var partiesMutex sync.Mutex
var partiesObject *parties
