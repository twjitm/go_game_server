package web

import (
	"fmt"
	"net"
)

func UnixNet() {
	socketFile := "/tmp/go_sock.sock"
	liseten, ero := net.Listen("unix", socketFile)
	if ero != nil {
		fmt.Println("start go unix socket net filed")
		return
	}
	conn, _ := liseten.Accept()
	var b []byte
	_, _ = conn.Read(b)
	//
	defer conn.Close()
}
