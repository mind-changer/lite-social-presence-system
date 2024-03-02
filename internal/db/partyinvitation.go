package db

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/lite-social-presence-system/internal/def"
	"github.com/sirupsen/logrus"
)

type PartyInvitations interface {
	SendPartyInvitation(ctx context.Context, partyId, ownerId, userId string) error
	AcceptPartyInvitation(ctx context.Context, partyId, userId string) error
	DeletePartyInvitation(ctx context.Context, partyId, userId string) error
	PartyInvitationExists(ctx context.Context, partyId, userId string) (bool, error)
}

type partyInvitations struct {
	*db
	conn *pgx.Conn
}

var partyInvitationsMutex sync.Mutex
var friendpartyInvitationsObject *partyInvitations

func (p *partyInvitations) SendPartyInvitation(ctx context.Context, partyId, ownerId, userId string) error {
	usersTable := p.db.GetUsersTable(ctx)
	userExists, err := usersTable.UserExists(ctx, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !userExists {
		logrus.WithError(err).Error("User doesnt exist")
		return def.CreateClientError(404, "user doesnt exist")
	}
	ownerExists, err := usersTable.UserExists(ctx, ownerId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if owner exists")
		return err
	}
	if !ownerExists {
		logrus.WithError(err).Error("Owner doesnt exist")
		return def.CreateClientError(404, "owner doesnt exist")
	}
	partyExists, err := p.db.GetPartiesTable(ctx).PartyExists(ctx, partyId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if party exists")
		return err
	}
	if !partyExists {
		logrus.WithError(err).Error("Party doesnt exist")
		return def.CreateClientError(404, "party doesnt exist")
	}
	isFriend, err := p.db.GetFriendsTable(ctx).IsFriend(ctx, ownerId, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if friend exists")
		return err
	}
	if !isFriend {
		logrus.WithError(err).Error("Friend doesnt exist")
		return def.CreateClientError(404, "friend doesnt exist")
	}
	exists, err := p.PartyInvitationExists(ctx, partyId, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if party invitation exists")
		return err
	}
	if exists {
		logrus.WithError(err).Error("party invitation already exists")
		return def.CreateClientError(409, "party invitation already exists")
	}

	insertSql := `
	insert into party_invitations(party_id,user_id) 
	values($1,$2);
	`
	_, err = p.conn.Exec(ctx, insertSql, partyId, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while inserting party invitation")
		return err
	}
	return nil
}

func (p *partyInvitations) PartyInvitationExists(ctx context.Context, partyId, userId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM party_invitations WHERE party_id=$1 AND user_id=$2)`
	exists := false
	err := p.conn.QueryRow(ctx, query, partyId, userId).Scan(&exists)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if party invitation exists")
		return false, err
	}
	logrus.Info(exists)
	return exists, nil
}

func (p *partyInvitations) AcceptPartyInvitation(ctx context.Context, partyId, userId string) error {
	partyInvExists, err := p.PartyInvitationExists(ctx, partyId, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if party invitation exists")
		return err
	}
	if !partyInvExists {
		logrus.WithError(err).Error("Party invitation doesnt exist")
		return def.CreateClientError(404, "party invitation doesnt exist")
	}
	if err := p.db.GetPartyMembersTable(ctx).AddPartyMember(ctx, partyId, userId); err != nil {
		logrus.WithError(err).Error("Error while inserting party member")
		return err
	}
	if err := p.DeletePartyInvitation(ctx, partyId, userId); err != nil {
		logrus.WithError(err).Error("Error while deleting party invitation")
		return err
	}
	return nil
}

func (p *partyInvitations) DeletePartyInvitation(ctx context.Context, partyId, userId string) error {

	partyInvExists, err := p.PartyInvitationExists(ctx, partyId, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if party invitation exists")
		return err
	}
	if !partyInvExists {
		logrus.WithError(err).Error("Party invitation doesnt exist")
		return def.CreateClientError(404, "party invitation  doesnt exist")
	}
	userExists, err := p.db.GetUsersTable(ctx).UserExists(ctx, userId)
	if err != nil {
		logrus.WithError(err).Error("Error while checking if user exists")
		return err
	}
	if !userExists {
		logrus.WithError(err).Error("User doesnt exist")
		return def.CreateClientError(404, "user doesnt exist")
	}
	insertSql := `
	DELETE from party_invitations
	WHERE user_id=$1 AND party_id=$2;
	`
	_, err = p.conn.Exec(ctx, insertSql, userId, partyId)
	if err != nil {
		logrus.WithError(err).Error("Error while deleting party invitation")
		return err
	}
	return nil
}
