package io

import (
	"bytes"
	"fmt"
	"time"
)

type MessageHead struct {
	ID     int
	Time   int64
	Length int64
}
type MessageBody struct {
	Body [] byte
}

type handler interface {
	GetHead()
	GetBody()
	Encode(a []int) ([]byte, error)
	Decode()
}

type UdpHandler struct {
	head MessageHead
	body MessageBody
	handler
}

type TcpHandler struct {
	handler
	head MessageHead
	body MessageBody
}

func (tcp *TcpHandler) Decode() {

}

func (tcp *TcpHandler) Encode(a []int) ([]byte, error) {

	var pkg = new(bytes.Buffer)
	pkg.WriteByte(8)
	headLength := 8  //tcp.head.ID
	headLength += 16 //tcp.head.Length
	headLength += 16 //tcp.head.Time
	headLength += len(tcp.body.Body)
	return pkg.Bytes(), error()
}

func Test() {
	a := []int{1}
	b := []byte{1}
	tcpHandler := TcpHandler{
		head: MessageHead{
			ID:     1,
			Time:   time.Now().Unix(),
			Length: 1,
		},
		body: MessageBody{Body: b},
	}
	result, _ := tcpHandler.Encode(a)
	fmt.Println(result)
}
