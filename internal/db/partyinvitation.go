package db

import (
	"sync"

	"github.com/jackc/pgx/v5"
)

type PartyInvitations interface {
	// GetConn(ctx context.Context) (*pgx.Conn, error)
}

type partyInvitations struct {
	conn *pgx.Conn
}

var partyInvitationsMutex sync.Mutex
var friendpartyInvitationsObject *partyInvitations
