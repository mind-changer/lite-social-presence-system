package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Parties interface {
	CreateParty(ctx context.Context, ownerId string) (string, error)
	PartyExists(ctx context.Context, partyId string) (bool, error)
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
		return "", fmt.Errorf("user doesnt exist")
	}

	insertSql := `
	insert into parties(id,owner_id) 
	values($1,$2);
	`
	partyId := uuid.NewString()
	_, err = p.conn.Exec(ctx, insertSql, partyId, ownerId)
	if err != nil {
		logrus.WithError(err).Error("Error while inserting friend")
		return "", err
	}
	return partyId, nil
}

func (p *parties) PartyExists(ctx context.Context, partyId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM parties WHERE id=$1)`
	exists := false
	err := p.conn.QueryRow(ctx, query, partyId).Scan(&exists)
	if err != nil {
		logrus.WithError(err).Error("Error while getting user friends")
		return false, err
	}
	logrus.Info("query,id ", query, partyId)
	logrus.Info(exists)
	return exists, nil
}
