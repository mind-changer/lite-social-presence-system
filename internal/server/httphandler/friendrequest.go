package httphandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lite-social-presence-system/config"
)

func UpdateFriendRequestsHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := w.Write([]byte("this is friend handling"))
		if err != nil {
			log.Fatal("ERROR WHILE WRITING RESPONSE")
		}
		fmt.Println(c)
	}
}
