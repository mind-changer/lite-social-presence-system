package grpchandler

import (
	"log"
	"time"

	onlinestatuspb "github.com/lite-social-presence-system/pb/onlinestatus"
)

func GetPartyMembers(in *onlinestatuspb.PartyMembersRequest, stream onlinestatuspb.PartyService_GetPartyMembersServer) error {
	log.Println("start user status server")
	log.Println(in.PartyId)
	timer := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			//get party memebers
			// hwStats, err := s.GetStats()
			// if err != nil {
			// 	log.Println(err.Error())
			// }
			// Send the Hardware stats on the stream
			if err := stream.Send(&onlinestatuspb.PartyMembersResponse{Members: []string{"A", "B"}}); err != nil {
				log.Println(err.Error())
			}
		}
	}
}
