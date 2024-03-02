package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lite-social-presence-system/config"
	"github.com/lite-social-presence-system/internal/db"
	"github.com/lite-social-presence-system/internal/def"
	"github.com/lite-social-presence-system/util"
	"github.com/sirupsen/logrus"
)

func AcceptFriendRequestHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pathVars := mux.Vars(r)
		userId := pathVars["user-id"]
		requesterId := pathVars["requester-id"]
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		if err := d.GetFriendRequestsTable(ctx).AcceptFriendRequest(ctx, userId, requesterId); err != nil {
			logrus.WithError(err).Error("Error while adding friend")
			if e, ok := err.(*def.ClientError); ok {
				util.SendErrorResponse(w, e.Code, e.Message)
				return
			}
			util.SendErrorResponse(w, 500, "error while adding friend")
			return
		}
		resp := &def.AcceptFriendRequestResponse{
			Status: "FRIEND_ADDED",
		}
		b, err := json.Marshal(resp)
		if err != nil {
			logrus.WithError(err).Error("Error while json encoding response")
			util.SendErrorResponse(w, 500, "error while json encoding response")
			return
		}
		_, err = w.Write(b)
		if err != nil {
			logrus.Fatal("ERROR WHILE WRITING RESPONSE")
			util.SendErrorResponse(w, 500, "error while writing body")
			return
		}
	}
}

func RejectFriendRequestHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pathVars := mux.Vars(r)
		userId := pathVars["user-id"]
		requesterId := pathVars["requester-id"]
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		if err := d.GetFriendRequestsTable(ctx).DeleteFriendRequest(ctx, userId, requesterId); err != nil {
			logrus.WithError(err).Error("Error while rejecting friend request")
			if e, ok := err.(*def.ClientError); ok {
				util.SendErrorResponse(w, e.Code, e.Message)
				return
			}
			util.SendErrorResponse(w, 500, "error while rejecting friend request")
			return
		}
		resp := &def.RejectRequestResponse{
			Status: "FRIEND_REQUEST_REJECTED",
		}
		b, err := json.Marshal(resp)
		if err != nil {
			logrus.WithError(err).Error("Error while json encoding response")
			util.SendErrorResponse(w, 500, "error while json encoding response")
			return
		}
		_, err = w.Write(b)
		if err != nil {
			logrus.Fatal("ERROR WHILE WRITING RESPONSE")
			util.SendErrorResponse(w, 500, "error while writing body")
			return
		}
	}
}
