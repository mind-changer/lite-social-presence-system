package server

import (
	"net"

	"github.com/lite-social-presence-system/config"
	"github.com/lite-social-presence-system/internal/server/grpchandler"
	onlinestatuspb "github.com/lite-social-presence-system/pb/onlinestatus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type userStatusServer struct {
	onlinestatuspb.UnimplementedUserStatusServiceServer
	cfg *config.Config
}
type partyServer struct {
	onlinestatuspb.UnimplementedPartyServiceServer
	cfg *config.Config
}

func (s *userStatusServer) GetUserStatus(in *onlinestatuspb.UserStatusRequest, stream onlinestatuspb.UserStatusService_GetUserStatusServer) error {
	logrus.Println("get user status called")
	return grpchandler.GetUserStatus(in, stream, s.cfg)
}

func (s *partyServer) GetPartyMembers(in *onlinestatuspb.PartyMembersRequest, stream onlinestatuspb.PartyService_GetPartyMembersServer) error {
	logrus.Println("get party members called")
	return grpchandler.GetPartyMembers(in, stream, s.cfg)
}

func registerGrpcServices(s *grpc.Server, cfg *config.Config) {
	onlinestatuspb.RegisterUserStatusServiceServer(s, &userStatusServer{cfg: cfg})
	onlinestatuspb.RegisterPartyServiceServer(s, &partyServer{cfg: cfg})
}

func RunGRPCServer(cfg *config.Config) {
	listener, err := net.Listen("tcp", ":81")
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	registerGrpcServices(s, cfg)
	logrus.Println("Starting grpc server")
	if err := s.Serve(listener); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
