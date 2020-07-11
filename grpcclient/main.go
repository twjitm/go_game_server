package grpcclient

import (
	"app/proto/message"
	"context"
	"fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	"google.golang.org/grpc"
	"log"
)

func Client() {

	address := ""
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := message.NewRpcGameServerClient(conn)
	name := []string{"twj", "funplus"}
	request := message.GetUserInfoRequest{
		Name: name,
	}
	resp, _ := client.GatUserInfo(context.Background(), &request)
	infos, _ := resp.Recv()
	fmt.Println(infos)

}
