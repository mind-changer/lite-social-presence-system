package db

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/lite-social-presence-system/internal/def"
	"github.com/sirupsen/logrus"
)

type PartyMembers interface {
	GetPartyMembers(ctx context.Context, partyId string) ([]string, error)
	KickPartyMember(ctx context.Context, partyId, ownerId, memberId string) error
	LeaveParty(ctx context.Context, partyId, memberId string) error
	AddPartyMember(ctx context.Context, partyId, userId string) error
	PartyMemberExists(ctx context.Context, partyId, userId string) (bool, error)
}

type partyMembers struct {
	*db
	conn *pgx.Conn
}

var partyMembersMutex sync.Mutex
var partyMembersObject *partyMembers

func (f *partyMembers) GetPartyMembers(ctx context.Context, partyId string) ([]string, error) {
	query := `select user_id from party_members where party_id=$1`
	logrus.Info("query ", query)
	rows, err := f.conn.Query(ctx, query, partyId)
	if err != nil {
		logrus.WithError(err).Error("Error while getting party members")
		return nil, err
	}
	partyMembers := make([]string, 0)
	for rows.Next() {
		partyMember := ""
		if err := rows.Scan(&partyMember); err != nil {
			logrus.WithError(err).Error("Error while decoding party members")
			return nil, err
		}
		partyMembers = append(partyMembers, partyMember)
	}
	logrus.Info(partyMembers)
	defer rows.Close()
	return partyMembers, nil
}

func (p *partyMembers) AddPartyMember(ctx context.Context, partyId, userId string) error {
	usersTable := p.db.GetUsersTable(ctx)
	userExists, err := usersTable.UserExists(ctx, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !userExists {
		logrus.WithError(err).Error("User doesnt exist")
		return def.CreateClientError(400, "user doesnt exist")
	}
	partyExists, err := p.db.GetPartiesTable(ctx).PartyExists(ctx, partyId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if party exists")
		return err
	}
	if !partyExists {
		logrus.WithError(err).Error("party doesnt exist")
		return def.CreateClientError(400, "party doesnt exist")
	}
	exists, err := p.PartyMemberExists(ctx, partyId, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if party member exists")
		return err
	}
	if exists {
		logrus.WithError(err).Error("party member already exists")
		return def.CreateClientError(409, "party member already exists")
	}
	insertSql := `
	insert into party_members(party_id,user_id) 
	values($1,$2),
	`
	_, err = p.conn.Exec(ctx, insertSql, partyId, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while inserting party member")
		return err
	}
	return nil
}

func (p *partyMembers) KickPartyMember(ctx context.Context, partyId, ownerId, memberId string) error {
	ownerExists, err := p.db.GetUsersTable(ctx).UserExists(ctx, ownerId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if owner exists")
		return err
	}
	if !ownerExists {
		logrus.WithError(err).Error("Owner doesnt exist")
		return def.CreateClientError(400, "owner doesnt exist")
	}
	isOwner, err := p.db.GetPartiesTable(ctx).IsOwner(ctx, partyId, ownerId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if owner exists")
		return err
	}
	if !isOwner {
		logrus.WithError(err).Error("Access denied. Requires owner access")
		return def.CreateClientError(400, "access denied. requires owner access")
	}

	if err = p.LeaveParty(ctx, partyId, memberId); err != nil {
		logrus.WithError(err).Error("Error while leaving party")
		return err
	}
	return nil
}

func (p *partyMembers) LeaveParty(ctx context.Context, partyId, userId string) error {
	usersTable := p.db.GetUsersTable(ctx)
	userExists, err := usersTable.UserExists(ctx, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !userExists {
		logrus.WithError(err).Error("User doesnt exist")
		return def.CreateClientError(400, "user doesnt exist")
	}
	partyExists, err := p.db.GetPartiesTable(ctx).PartyExists(ctx, partyId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !partyExists {
		logrus.WithError(err).Error("Friend doesnt exist")
		return def.CreateClientError(400, "user doesnt exist")
	}
	deleteSql := `
	DELETE FROM party_members
	WHERE party_id=$1 AND user_id=$2;	
	`
	_, err = p.conn.Exec(ctx, deleteSql, partyId, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while inserting friend")
		return err
	}
	return nil
}

func (p *partyMembers) PartyMemberExists(ctx context.Context, partyId, userId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM party_members WHERE party_id=$1 AND user_id=$2)`
	exists := false
	err := p.conn.QueryRow(ctx, query, partyId, userId).Scan(&exists)
	if err != nil {
		logrus.WithError(err).Error("Error while getting user friends")
		return false, err
	}
	logrus.Info("query,id ", query, userId)
	logrus.Info(exists)
	return exists, nil
}
