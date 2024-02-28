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
	//VIEW AND REMOVE friends
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/friends", httphandler.ViewFriendsHandler(cfg)).Methods("GET")
	//ADD friend
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/friends", httphandler.AddFriendHandler(cfg)).Methods("POST")
	//REMOVE friend
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/friends/{friend-id}", httphandler.UpdateFriendsHandler(cfg)).Methods("PATCH")

	//ACCEPT, REJECT friend requests
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/friend-requests/{user-id}", httphandler.UpdateFriendRequestsHandler(cfg)).Methods("PATCH")

	//CREATE party
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/parties", httphandler.CreatePartyHandler(cfg)).Methods("POST")
	//LEAVE, KICK party
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/parties/{party-id}", httphandler.UpdatePartyHandler(cfg)).Methods("PATCH")
	//SEND party invitation
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/parties/{party-id}/member-invitations", httphandler.SendPartyMemberInvitationsHandler(cfg)).Methods("POST")

	//JOIN or REJECT party invitation
	r.HandleFunc("/lite-social-presence-system/users/{user-id}/party-invitations/{party-id}", httphandler.UpdatePartyInvitationsHandler(cfg)).Methods("PATCH")
}

func RunHTTPServer(cfg *config.Config) {

	r := mux.NewRouter()
	registerAPIHandlers(r, cfg)
	httpServer := http.Server{
		Addr:         "127.0.0.1:80",
		Handler:      r,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	logrus.Info("Starting http server")
	if err := httpServer.ListenAndServe(); err != nil {
		logrus.Fatal("SERVER QUIT ERROR", err)
	}
}
