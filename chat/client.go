package chat

import (
	"fmt"
	"io"
	"net"
)

type ChatClient interface {
	Dial(address string) error
	Send(command interface{}) error
	SendMessage(message string) error
	SetName(name string) error
	Start()
	Close()
	Incoming() chan MessageCommand
}

type TcpChatClient struct {
	conn      net.Conn
	cmdReader *CommandReader
	cmdWriter *CommandWriter
	name      string
	incoming  chan MessageCommand
}

func NewClient() *TcpChatClient {
	return &TcpChatClient{
		incoming: make(chan MessageCommand),
	}
}

func (c *TcpChatClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)

	if err == nil {
		c.conn = conn
	}

	c.cmdReader = NewCommandReader(conn)
	c.cmdWriter = NewCommandWriter(conn)

	return err
}

func (c *TcpChatClient) Start() {
	for {
		cmd, err := c.cmdReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Read error %v", err)
		}

		if cmd != nil {
			switch v := cmd.(type) {
			case MessageCommand:
				c.incoming <- v
			default:
				fmt.Printf("Unknown command: %v", v)
			}
		}
	}
}

func (c *TcpChatClient) Close() {
	_ = c.conn.Close()
}

func (c *TcpChatClient) Incoming() chan MessageCommand {
	return c.incoming
}

func (c *TcpChatClient) Send(command interface{}) error {
	_, err := c.cmdWriter.Write(command)
	return err
}

func (c *TcpChatClient) SetName(name string) error {
	return c.Send(NameCommand{name})
}

func (c *TcpChatClient) SendMessage(message string) error {
	return c.Send(SendCommand{
		Message: message,
	})
}
