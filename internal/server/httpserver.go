package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lite-social-presence-system/config"
	"github.com/lite-social-presence-system/internal/server/httphandler"
	"github.com/sirupsen/logrus"
)

func registerAPIHandlers(r *mux.Router, cfg *config.Config) {
	//Update user status
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/status", httphandler.UpdateUserStatusHandler(cfg)).Methods("PUT")

	//VIEW friend
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/friends", httphandler.ViewFriendsHandler(cfg)).Methods("GET")
	//ADD friend
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/friends", httphandler.AddFriendHandler(cfg)).Methods("POST")
	//REMOVE friend
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/friends/{friend-id}", httphandler.RemoveFriendHandler(cfg)).Methods("DELETE")

	//ACCEPT friend request
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/friend-requests/{requester-id}", httphandler.AcceptFriendRequestHandler(cfg)).Methods("PATCH")
	//REJECT friend request
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/friend-requests/{requester-id}", httphandler.RejectFriendRequestHandler(cfg)).Methods("DELETE")

	//CREATE party
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/parties", httphandler.CreatePartyHandler(cfg)).Methods("POST")

	//SEND party invitation
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/parties/{party-id}/member-invitations", httphandler.SendPartyInvitationHandler(cfg)).Methods("POST")
	//ACCEPT party invitation
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/party-invitations/{party-id}", httphandler.AcceptPartyInvitationHandler(cfg)).Methods("PATCH")
	//REJECT party invitation
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/party-invitations/{party-id}", httphandler.RejectPartyInvitationHandler(cfg)).Methods("DELETE")

	//KICK party member
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/parties/{party-id}/members/{member-id}", httphandler.KickPartyMemberHandler(cfg)).Methods("DELETE")

	//LEAVE party
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/joined-parties/current/{party-id}", httphandler.LeavePartyHandler(cfg)).Methods("DELETE")
}

func RunHTTPServer(cfg *config.Config) {

	r := mux.NewRouter()
	registerAPIHandlers(r, cfg)
	httpServer := http.Server{
		Addr:         "0.0.0.0:80",
		Handler:      r,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	logrus.Info("Starting http server")
	if err := httpServer.ListenAndServe(); err != nil {
		logrus.Fatal("SERVER QUIT ERROR", err)
	}
}
