package chat

import (
	"net"
	"sync"
)

type ChatServer interface {
	Listen(address string) error

	Broadcast(command interface{}) error

	Start()

	Close()
}

type TcpChatServer struct {
	listener net.Listener
	clients  []*client
	mutex    *sync.Mutex
}

type client struct {
	conn   net.Conn
	name   string
	writer *CommandWriter
}
