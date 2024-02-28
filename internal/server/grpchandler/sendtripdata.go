package grpchandler

import (
	"io"
	"log"

	onlinestatuspb "github.com/lite-social-presence-system/pb/onlinestatus"
)

func SendTripData(srv onlinestatuspb.TripService_SendTripDataServer) error {
	log.Println("start new server")
	ctx := srv.Context()
	for {

		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
		resp := &onlinestatuspb.TripSummaryResponse{
			Distance: req.Latitude,
			Charge:   10,
		}
		if err := srv.Send(resp); err != nil {
			log.Printf("send error %v", err)
		}
	}

}
