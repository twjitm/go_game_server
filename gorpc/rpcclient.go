package gorpc

import (
	"fmt"
	"net/rpc"
)

func ConnectionRpc() {
	client, e := rpc.DialHTTP("tcp", "127.0.0.1:1234")

	if e != nil {
		fmt.Println("connect error", e)
		return
	}
	arg := &Args{
		A: 2,
		B: 4,
	}
	var rep int
	cc := client.Go("Arith.Multiply", arg, rep, nil)
	re := <-cc.Done
	fmt.Println(re)

	e = client.Call("Arith.Multiply", arg, &rep)
	if e != nil {
		fmt.Println("Arith.Multiply err", e)
	}
}
