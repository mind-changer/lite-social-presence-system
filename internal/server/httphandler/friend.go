package httphandler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lite-social-presence-system/config"
	"github.com/lite-social-presence-system/internal/db"
	"github.com/lite-social-presence-system/internal/def"
	"github.com/lite-social-presence-system/util"
	"github.com/sirupsen/logrus"
)

func ViewFriendsHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pathVars := mux.Vars(r)
		userId := pathVars["user-id"]
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		friendIds, err := d.GetFriendsTable(ctx).GetFriends(ctx, userId)
		if err != nil {
			logrus.WithError(err).Error("Error while reading friends")
			util.SendErrorResponse(w, 500, "error while while reading friends")
			return
		}
		logrus.Info("rows ", friendIds)
		resp := &def.ViewFriendsResponse{
			Friends: friendIds,
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
			return
		}
		//NOTE: can introduce pagination in future
	}
}

func AddFriendHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pathVars := mux.Vars(r)
		userId := pathVars["user-id"]
		b, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.WithError(err).Error("Error while reading body")
			util.SendErrorResponse(w, 500, "error whil reading body")
			return
		}
		req := &def.AddFriendRequest{}
		err = json.Unmarshal(b, req)
		if err != nil {
			logrus.WithError(err).Error("Error while json decoding request")
			util.SendErrorResponse(w, 500, "error while getting json decoding request")
			return
		}
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		if err := d.GetFriendRequestsTable(ctx).SendFriendRequest(ctx, req.User, userId); err != nil {
			logrus.WithError(err).Error("Error while reading friends")
			util.SendErrorResponse(w, 500, "error while while reading friends")
			return
		}
		resp := &def.AddFriendResponse{
			Status: "FRIEND_REQUEST_SENT",
		}
		b, err = json.Marshal(resp)
		if err != nil {
			logrus.WithError(err).Error("Error while json encoding response")
			util.SendErrorResponse(w, 500, "error while json encoding response")
			return
		}
		_, err = w.Write(b)
		if err != nil {
			logrus.Fatal("ERROR WHILE WRITING RESPONSE")
			return
		}
	}
}

func RemoveFriendHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pathVars := mux.Vars(r)
		userId := pathVars["user-id"]
		friendId := pathVars["friend-id"]
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		if err := d.GetFriendsTable(ctx).RemoveFriend(ctx, userId, friendId); err != nil {
			logrus.WithError(err).Error("Error while reading friends")
			util.SendErrorResponse(w, 500, "error while while reading friends")
			return
		}
		resp := &def.RemoveFriendResponse{
			Status: "FRIEND_REMOVED",
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
			return
		}
	}
}
