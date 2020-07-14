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

func (s *Server) GatUserInfo(request *message.GetUserInfoRequest, server message.Rpc_GatUserInfoServer) error {
	names := request.Name
	fmt.Print(names)
	_ = server.Send(&message.GetUserInfoReply{
		UserList: nil,
	})
	return nil
}

func (s *Server) SendMessage(c context.Context, in *message.ChatInfo) (*message.ChatInfo, error) {
	return nil, nil
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
