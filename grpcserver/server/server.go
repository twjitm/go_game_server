package server

import (
	message "app/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	Rpc *message.RpcServer
}

func (s *Server) GatUserInfo(c context.Context, request *message.GetUserInfoRequest) (*message.GetUserInfoReply, error) {
	names := request.Name
	fmt.Print(names)
	return nil, nil
}

func (s *Server) SendMessage(server message.Rpc_SendMessageServer) error {
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
	message.RegisterRpcServer(s, &Server{})
	_ = s.Serve(lis)
}
