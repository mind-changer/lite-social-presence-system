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

func SendPartyInvitationHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pathVars := mux.Vars(r)
		ownerId := pathVars["user-id"]
		partyId := pathVars["party_id"]
		b, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.WithError(err).Error("Error while reading body")
			util.SendErrorResponse(w, 500, "error whil reading body")
			return
		}
		req := &def.SendPartyInvitationRequest{}
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
		if err := d.GetPartyInvitationsTable(ctx).SendPartyInvitation(ctx, partyId, ownerId, req.UserId); err != nil {
			logrus.WithError(err).Error("Error while sending party invitation")
			util.SendErrorResponse(w, 500, "error while sending party invitation")
			return
		}
		resp := &def.SendPartyInvitationResponse{
			Status: "PARTY_INVITATION_SENT",
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

func AcceptPartyInvitationHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pathVars := mux.Vars(r)
		userId := pathVars["user-id"]
		partyId := pathVars["party_id"]
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		if err := d.GetPartyInvitationsTable(ctx).AcceptPartyInvitation(ctx, partyId, userId); err != nil {
			logrus.WithError(err).Error("Error while accepting party invitation")
			util.SendErrorResponse(w, 500, "error while accepting party invitation")
			return
		}
		resp := &def.AcceptPartyInvitationResponse{
			Status: "PARTY_INVITATION_ACCEPTED",
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

func RejectPartyInvitationHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pathVars := mux.Vars(r)
		userId := pathVars["user-id"]
		partyId := pathVars["party_id"]
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		if err := d.GetPartyInvitationsTable(ctx).RejectPartyInvitation(ctx, partyId, userId); err != nil {
			logrus.WithError(err).Error("Error while rejecting party invitation")
			util.SendErrorResponse(w, 500, "error while rejecting party invitation")
			return
		}
		resp := &def.RejectPartyInvitationResponse{
			Status: "PARTY_INVITATION_REJECTED",
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
