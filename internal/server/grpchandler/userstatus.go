package grpchandler

import (
	"context"
	"time"

	"github.com/lite-social-presence-system/config"
	"github.com/lite-social-presence-system/internal/db"
	onlinestatuspb "github.com/lite-social-presence-system/pb/onlinestatus"
	"github.com/sirupsen/logrus"
)

func GetUserStatus(in *onlinestatuspb.UserStatusRequest, stream onlinestatuspb.UserStatusService_GetUserStatusServer, cfg *config.Config) error {
	timer := time.NewTicker(5 * time.Second)
	ctx := stream.Context()
	d, err := db.GetDBObject(ctx, cfg.Postgres)
	if err != nil {
		logrus.WithError(err).Error("Error while getting db object")
		return err
	}
	if err := sendUserStatus(ctx, in, stream, d); err != nil {
		logrus.WithError(err).Error("Error while sending user status")
		return err
	}
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			if err := sendUserStatus(ctx, in, stream, d); err != nil {
				logrus.WithError(err).Error("Error while sending user status")
				return err
			}
		}
	}
}

func sendUserStatus(ctx context.Context, in *onlinestatuspb.UserStatusRequest, stream onlinestatuspb.UserStatusService_GetUserStatusServer, d db.DB) error {
	status, err := d.GetUsersTable(ctx).GetUserStatus(ctx, in.UserId)
	if err != nil {
		logrus.WithError(err).Error("Error while getting user status")
		return err
	}
	if err := stream.Send(&onlinestatuspb.UserStatusResponse{Status: status}); err != nil {
		logrus.WithError(err).Error("Error while sending response")
		return err
	}
	return nil
}
