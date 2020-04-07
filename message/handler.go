package message

import "fmt"

/**
接口学习
*/
type handler interface {
	handler() int
}

type TcpHandler struct {
}

type UdpHandler struct {

}

func (tcp TcpHandler) handler() int {
	fmt.Println("tcp handler")
	return 0
}

func (upd UdpHandler) handler() int {
	fmt.Println("udp")
	return 0
}


