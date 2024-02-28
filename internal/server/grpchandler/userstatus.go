package grpchandler

import (
	"log"
	"time"

	onlinestatuspb "github.com/lite-social-presence-system/pb/onlinestatus"
)

func GetUserStatus(in *onlinestatuspb.UserStatusRequest, stream onlinestatuspb.UserStatusService_GetUserStatusServer) error {
	log.Println("start user status server")
	log.Println(in.UserId)
	timer := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			//get user status
			// hwStats, err := s.GetStats()
			// if err != nil {
			// 	log.Println(err.Error())
			// }
			// Send the Hardware stats on the stream
			if err := stream.Send(&onlinestatuspb.UserStatusResponse{Status: "ONLINE"}); err != nil {
				log.Println(err.Error())
			}
		}
	}
}
