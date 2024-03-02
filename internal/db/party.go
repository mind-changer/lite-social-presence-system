package db

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lite-social-presence-system/internal/def"
	"github.com/sirupsen/logrus"
)

type Parties interface {
	CreateParty(ctx context.Context, ownerId string) (string, error)
	PartyExists(ctx context.Context, partyId string) (bool, error)
	IsOwner(ctx context.Context, partyId, ownerId string) (bool, error)
}

type parties struct {
	*db
	conn *pgx.Conn
}

var partiesMutex sync.Mutex
var partiesObject *parties

func (p *parties) CreateParty(ctx context.Context, ownerId string) (string, error) {
	usersTable := p.db.GetUsersTable(ctx)
	userExists, err := usersTable.UserExists(ctx, ownerId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return "", err
	}
	if !userExists {
		logrus.WithError(err).Error("User doesnt exist")
		return "", def.CreateClientError(400, "user doesnt exist")
	}

	insertSql := `
	insert into parties(id,owner_id) 
	values($1,$2);
	`
	partyId := uuid.NewString()
	_, err = p.conn.Exec(ctx, insertSql, partyId, ownerId)
	if err != nil {
		logrus.WithError(err).Error("Error while inserting party")
		return "", err
	}
	return partyId, nil
}

func (p *parties) PartyExists(ctx context.Context, partyId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM parties WHERE id=$1)`
	exists := false
	err := p.conn.QueryRow(ctx, query, partyId).Scan(&exists)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if party exists")
		return false, err
	}
	logrus.Info(exists)
	return exists, nil
}

func (p *parties) IsOwner(ctx context.Context, partyId, ownerId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM parties WHERE id=$1 AND owner_id=$2)`
	exists := false
	err := p.conn.QueryRow(ctx, query, partyId, ownerId).Scan(&exists)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if valid owner")
		return false, err
	}
	logrus.Info(exists)
	return exists, nil
}
