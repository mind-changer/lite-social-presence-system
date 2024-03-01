package grpchandler

import (
	"context"
	"time"

	"github.com/lite-social-presence-system/config"
	"github.com/lite-social-presence-system/internal/db"
	onlinestatuspb "github.com/lite-social-presence-system/pb/onlinestatus"
	"github.com/sirupsen/logrus"
)

func GetPartyMembers(in *onlinestatuspb.PartyMembersRequest, stream onlinestatuspb.PartyService_GetPartyMembersServer, cfg *config.Config) error {
	timer := time.NewTicker(2 * time.Second)
	ctx := stream.Context()
	d, err := db.GetDBObject(ctx, cfg.Postgres)
	if err != nil {
		logrus.WithError(err).Error("Error while getting db object")
		return err
	}
	if err := sendPartyMembers(ctx, in, stream, d); err != nil {
		logrus.WithError(err).Error("Error while sending party members")
		return err
	}
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			if err := sendPartyMembers(ctx, in, stream, d); err != nil {
				logrus.WithError(err).Error("Error while sending party members")
				return err
			}
		}
	}
}

func sendPartyMembers(ctx context.Context, in *onlinestatuspb.PartyMembersRequest, stream onlinestatuspb.PartyService_GetPartyMembersServer, d db.DB) error {
	members, err := d.GetPartyMembersTable(ctx).GetPartyMembers(ctx, in.PartyId)
	if err != nil {
		logrus.WithError(err).Error("Error while getting party members")
		return err
	}
	if err := stream.Send(&onlinestatuspb.PartyMembersResponse{Members: members}); err != nil {
		logrus.WithError(err).Error("Error while sending response")
		return err
	}
	return nil
}
