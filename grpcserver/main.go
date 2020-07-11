package grpcserver

import (
	"app/proto/message"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	Rpc *message.RpcGameServerServer
}

func (s *Server) GatUserInfo(request *message.GetUserInfoRequest, server message.RpcGameServer_GatUserInfoServer) error {
	names := request.Name
	fmt.Print(names)
	_ = server.Send(&message.UserInfo{
		Id:        1,
		Name:      "twj",
		Email:     "1029718215@qq.com",
		Phone:     "15733280862",
		Addresses: nil,
	})
	return nil
}

func Start() {
	address := "127.0.0.1:10086"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()

	message.RegisterRpcGameServerServer(s, &Server{})
	_ = s.Serve(lis)
}
