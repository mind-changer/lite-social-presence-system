package httphandler

import (
	"encoding/json"
	"fmt"
	"log"
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
		// b, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	logrus.WithError(err).Error("Error while reading body")
		// 	util.SendErrorResponse(w, 500, "error whil reading body")
		// 	return
		// }
		// req := &def.AddFriendRequest{}
		// err = json.Unmarshal(b, req)
		// if err != nil {
		// 	logrus.WithError(err).Error("Error while json decoding request")
		// 	util.SendErrorResponse(w, 500, "error while getting json decoding request")
		// 	return
		// }
		d, err := db.GetDBObject(ctx, cfg.Postgres)
		if err != nil {
			logrus.WithError(err).Error("Error while getting db object")
			util.SendErrorResponse(w, 500, "error while getting db object")
			return
		}
		rows, err := d.GetFriendsTable(ctx).ReadFriends(ctx, userId)
		if err != nil {
			logrus.WithError(err).Error("Error while reading friends")
			util.SendErrorResponse(w, 500, "error while while reading friends")
			return
		}
		logrus.Info("rows ", rows)
		resp := &def.ViewFriendsResponse{
			Friends: []string{"lmao"},
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
		c, err := w.Write([]byte("this is friend handling"))
		if err != nil {
			log.Fatal("ERROR WHILE WRITING RESPONSE")
		}
		fmt.Println(c)
	}
}

func UpdateFriendsHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := w.Write([]byte("this is friend handling"))
		if err != nil {
			log.Fatal("ERROR WHILE WRITING RESPONSE")
		}
		fmt.Println(c)
	}
}
