package chat

import (
	"fmt"
	"io"
	"net"
	"sync"
)

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

func (s *TcpChatServer) Listen(address string) error {
	l, err := net.Listen("tcp", address)
	if err == nil {
		s.listener = l
	}
	fmt.Printf("Listening on %v", address)
	return err
}

func (s *TcpChatServer) Close() {
	_ = s.listener.Close()
}

func (s *TcpChatServer) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Print(err)
		} else {
			client := s.accept(conn)
			go s.serve(client)
		}
	}
}

func (s *TcpChatServer) accept(conn net.Conn) *client {
	fmt.Printf("Accepting connection from %v, total clients: %v\n", conn.RemoteAddr().String(), len(s.clients)+1)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	client := &client{
		conn:   conn,
		writer: NewCommandWriter(conn),
	}
	s.clients = append(s.clients, client)
	return client
}

func (s *TcpChatServer) remove(client *client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// remove the connections from clients array
	for i, check := range s.clients {
		if check == client {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
		}
	}
	fmt.Printf("Closing connection from %v", client.conn.RemoteAddr().String())
	_ = client.conn.Close()
}

func (s *TcpChatServer) serve(client *client) {
	cmdReader := NewCommandReader(client.conn)
	defer s.remove(client)
	for {
		cmd, err := cmdReader.Read()
		if err != nil && err != io.EOF {
			fmt.Printf("Read error: %v", err)
		}
		if cmd != nil {
			switch v := cmd.(type) {
			case SendCommand:
				go s.Broadcast(MessageCommand{
					Message: v.Message,
					Name:    client.name,
				})
			case NameCommand:
				client.name = v.Name
			}
		}
		if err == io.EOF {
			break
		}
	}
}

func (s *TcpChatServer) Broadcast(command interface{}) error {
	for _, client := range s.clients {
		_, _ = client.writer.Write(command)
	}
	return nil
}

var Address = "127.0.0.1:3333"

func StartServer() *TcpChatServer {
	s := &TcpChatServer{
		mutex: &sync.Mutex{},
	}
	_ = s.Listen(Address)
	s.Start()
	return s
}
