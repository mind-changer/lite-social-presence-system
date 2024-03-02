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

func UpdateUserStatusHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
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
		req := &def.UpdateUserStatusRequest{}
		err = json.Unmarshal(b, req)
		if err != nil {
			logrus.WithError(err).Error("Error while json decoding request")
			util.SendErrorResponse(w, 500, "error while json decoding request")
			return
		}
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		if err := d.GetUsersTable(ctx).UpdateUserStatus(ctx, userId, req.UserStatus); err != nil {
			logrus.WithError(err).Error("Error while updating user status")
			if e, ok := err.(*def.ClientError); ok {
				util.SendErrorResponse(w, e.Code, e.Message)
				return
			}
			util.SendErrorResponse(w, 500, "error while updating user status")
			return
		}
		resp := &def.UpdateUserStatusResponse{
			Status: "USER_STATUS_UPDATED",
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
			util.SendErrorResponse(w, 500, "error while writing body")
			return
		}
	}
}
