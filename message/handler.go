package message

import "fmt"

/**
接口学习
*/
type handler interface {
	builder() int
	getBody() []byte
}

type TcpHandler struct {
	Cmd   int
	Ctime int64
	MType int
	Head  [3]int
	Body  []byte
}

type UdpHandler struct {
	Cmd  int64
	Head [3]int
	Body []byte
}

func (tcp TcpHandler) getBody() []byte {
	fmt.Println("tcp get Body")
	return nil
}

func (udp UdpHandler) getBody() []byte {
	fmt.Println("udp get body")
	return nil
}

func (tcp TcpHandler) builder() int {
	var cmd = tcp.Cmd
	fmt.Println("tcp handler")
	fmt.Println(cmd)
	return 0
}

func (udp *UdpHandler) builder() int {
	fmt.Println("udp handler")
	return 0
}

func Builder(handler handler) {
	handler.builder()
}
func GetBody(handler handler) {
	handler.getBody()
}
