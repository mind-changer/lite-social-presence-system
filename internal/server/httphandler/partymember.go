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

func KickPartyMemberHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pathVars := mux.Vars(r)
		ownerId := pathVars["user-id"]
		partyId := pathVars["party-id"]
		memberId := pathVars["member-id"]
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		if err := d.GetPartyMembersTable(ctx).KickPartyMember(ctx, partyId, ownerId, memberId); err != nil {
			logrus.WithError(err).Error("Error while kicking party member")
			util.SendErrorResponse(w, 500, "error while kicking party member")
			return
		}
		resp := &def.KickPartyMemberResponse{
			Status: "PARTY_MEMBER_REMOVED",
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

func LeavePartyHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pathVars := mux.Vars(r)
		partyId := pathVars["party-id"]
		memberId := pathVars["member-id"]
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		if err := d.GetPartyMembersTable(ctx).LeaveParty(ctx, partyId, memberId); err != nil {
			logrus.WithError(err).Error("Error while leaving party")
			util.SendErrorResponse(w, 500, "error while leaving party")
			return
		}
		resp := &def.LeavePartyResponse{
			Status: "PARTY_MEMBER_REMOVED",
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
