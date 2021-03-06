package client

import (
	"context"
	"fmt"
	"go_game_server/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Client() {

	address := "127.0.0.1:10086"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewRpcClient(conn)
	//name := []string{"twj", "funplus"}
	//request := message.GetUserInfoRequest{
	//	Name: name,
	//}
	//resp, _ := client.GatUserInfo(context.Background(), &request)
	//infos := resp.UserList
	//fmt.Println(infos)
	chat := proto.ChatInfo{
		Id:   0,
		Type: 0,
		Context: &proto.ChatContext{
			Type:    0,
			Context: "hello ,i am client",
		},
		Time:     time.Now().Unix(),
		Sender:   0,
		Receiver: 0,
		Topic:    0,
	}
	sender, _ := client.SendMessage(context.Background())
	_ = sender.Send(&chat)
	for {
		result, err := sender.Recv()

		if err != nil {
			return
		}
		fmt.Println(result.String())
	}



}
