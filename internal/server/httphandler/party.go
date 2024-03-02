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

func CreatePartyHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
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
		partyId, err := d.GetPartiesTable(ctx).CreateParty(ctx, userId)
		if err != nil {
			logrus.WithError(err).Error("Error while creating party")
			if e, ok := err.(*def.ClientError); ok {
				util.SendErrorResponse(w, e.Code, e.Message)
				return
			}
			util.SendErrorResponse(w, 500, "error while creating party")
			return
		}
		resp := &def.CreatePartyResponse{
			PartyId: partyId,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			logrus.WithError(err).Error("Error while json encoding response")
			util.SendErrorResponse(w, 500, "error while json encoding response")
			return
		}
		w.WriteHeader(201)
		_, err = w.Write(b)
		if err != nil {
			logrus.Fatal("ERROR WHILE WRITING RESPONSE")
			util.SendErrorResponse(w, 500, "error while writing body")
			return
		}
	}
}
