package client

import (
	message "app/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func Client() {

	address := "127.0.0.1:1008611"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := message.NewRpcClient(conn)
	name := []string{"twj", "funplus"}
	request := message.GetUserInfoRequest{
		Name: name,
	}
	resp, _ := client.GatUserInfo(context.Background(), &request)
	infos, _ := resp.Recv()
	fmt.Println(infos)

}
