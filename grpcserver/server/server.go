package server

import (
	"app/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	Address    string
	gameServer *GameServer
	grpcServer *grpc.Server
}

func (s *Server) GatUserInfo(c context.Context, request *proto.GetUserInfoRequest) (*proto.GetUserInfoReply, error) {
	names := request.Name
	fmt.Print(names)
	user := proto.UserInfo{
		Id:        0,
		Name:      "",
		Email:     "",
		Phone:     "",
		Addresses: nil,
	}
	var list []*proto.UserInfo
	list[0] = &user
	reply := proto.GetUserInfoReply{
		UserList: list,
	}
	return &reply, nil
}

func (s *Server) SendMessage(server proto.Rpc_SendMessageServer) error {
	recvChat := proto.ChatInfo{}
	recvChat.Type = 1
	_ = server.RecvMsg(&recvChat)

	fmt.Println(recvChat.Context.Context)

	sendChat := proto.ChatInfo{
		Id:   0,
		Type: 0,
		Context: &proto.ChatContext{
			Type:    0,
			Context: "hello,im server,go language",
		},
		Time:     0,
		Sender:   0,
		Receiver: 0,
		Topic:    0,
	}
	_ = server.Send(&sendChat)
	sendChat.Context.Context = "hello client,the message send from server"
	_ = server.Send(&sendChat)

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
	proto.RegisterRpcServer(s, &Server{
		Address:    "127.0.0",
		gameServer: &GameServer{},
		grpcServer: s,
	})
	_ = s.Serve(lis)
}
