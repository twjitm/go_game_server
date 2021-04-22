package gorpc

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	fmt.Println("Multiply")
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return fmt.Errorf("error")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	fmt.Println("Divide")
	return nil
}

func StartRpc() {
	arith := new(Arith)
	_ = rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "127.0.0.1:1234")
	if e != nil {
		fmt.Println("listen error:", e)
	}
	_ = http.Serve(l, nil)
}
