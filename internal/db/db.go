package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/lite-social-presence-system/config"
)

type DB interface {
	GetFriendsTable(ctx context.Context) Friends
	GetFriendRequestsTable(ctx context.Context) FriendRequests
	GetPartiesTable(ctx context.Context) Parties
	GetPartyInvitationsTable(ctx context.Context) PartyInvitations
	GetUsersTable(ctx context.Context) Users
	GetPartyMembersTable(ctx context.Context) PartyMembers
}

type db struct {
	conn *pgx.Conn
}

var dbMutex sync.Mutex
var dbObject *db

func GetDBObject(ctx context.Context, dbConfig *config.PostgresConfig) (DB, error) {
	var err error
	if dbObject != nil {
		return dbObject, nil
	}
	dbMutex.Lock()
	defer dbMutex.Unlock()
	if dbObject == nil {
		url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
			dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
		fmt.Println("postgres url", url)
		conn, dbErr := pgx.Connect(ctx, url)
		if dbErr != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			return nil, dbErr
		}
		dbObject = &db{conn}
	}
	return dbObject, nil
}

func (d *db) GetFriendsTable(ctx context.Context) Friends {
	if friendsObject != nil {
		return friendsObject
	}
	friendsMutex.Lock()
	defer friendsMutex.Unlock()
	if friendsObject == nil {
		friendsObject = &friends{d, d.conn}
	}
	return friendsObject
}
func (d *db) GetFriendRequestsTable(ctx context.Context) FriendRequests {
	if friendRequestsObject != nil {
		return friendRequestsObject
	}
	friendRequestsMutex.Lock()
	defer friendRequestsMutex.Unlock()
	if friendRequestsObject == nil {
		friendRequestsObject = &friendRequests{d, d.conn}
	}
	return friendRequestsObject
}
func (d *db) GetPartiesTable(ctx context.Context) Parties {
	if partiesObject != nil {
		return partiesObject
	}
	partiesMutex.Lock()
	defer partiesMutex.Unlock()
	if partiesObject == nil {
		partiesObject = &parties{d, d.conn}
	}
	return partiesObject
}
func (d *db) GetPartyInvitationsTable(ctx context.Context) PartyInvitations {
	if friendpartyInvitationsObject != nil {
		return friendpartyInvitationsObject
	}
	partyInvitationsMutex.Lock()
	defer partyInvitationsMutex.Unlock()
	if friendpartyInvitationsObject == nil {
		friendpartyInvitationsObject = &partyInvitations{d, d.conn}
	}
	return friendpartyInvitationsObject
}
func (d *db) GetUsersTable(ctx context.Context) Users {
	if usersObject != nil {
		return usersObject
	}
	usersMutex.Lock()
	defer usersMutex.Unlock()
	if usersObject == nil {
		usersObject = &users{d, d.conn}
	}
	return usersObject
}
func (d *db) GetPartyMembersTable(ctx context.Context) PartyMembers {
	if partyMembersObject != nil {
		return partyMembersObject
	}
	partyMembersMutex.Lock()
	defer partyMembersMutex.Unlock()
	if partyMembersObject == nil {
		partyMembersObject = &partyMembers{d, d.conn}
	}
	return partyMembersObject
}
