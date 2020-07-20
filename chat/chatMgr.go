package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

// SendCommand is used for sending new message from client
type SendCommand struct {
	Message string
}

// NameCommand is used for setting client display name
type NameCommand struct {
	Name string
}

// MessageCommand is used for notifying new messages
type MessageCommand struct {
	Name    string
	Message string
}

type CommandWriter struct {
	writer io.Writer
}

func NewCommandWriter(writer io.Writer) *CommandWriter {

	return &CommandWriter{
		writer: writer,
	}
}
func (w *CommandWriter) writeString(msg string) error {
	_, err := w.writer.Write([]byte(msg))
	return err
}

func (w *CommandWriter) Write(command interface{}) (int, error) {
	// naive implementation ...
	var err error
	switch v := command.(type) {
	case SendCommand:
		err = w.writeString(fmt.Sprintf("SEND %v\n", v.Message))
	case MessageCommand:
		err = w.writeString(fmt.Sprintf("MESSAGE %v %v\n", v.Name, v.Message))
	case NameCommand:
		err = w.writeString(fmt.Sprintf("NAME %v\n", v.Name))
	}
	return 0, err
}

type CommandReader struct {
	reader *bufio.Reader
}

func NewCommandReader(reader io.Reader) *CommandReader {

	return &CommandReader{
		reader: bufio.NewReader(reader),
	}
}

func (r *CommandReader) Read() (interface{}, error) {

	commandName, err := r.reader.ReadString(' ')
	if err != nil {
		return nil, err
	}
	switch commandName {
	case "MESSAGE ":
		user, err := r.reader.ReadString(' ')
		if err != nil {
			return nil, err
		}
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return MessageCommand{
			user[:len(user)-1],
			message[:len(message)-1],
		}, nil
		// similar implementation for other commands
	default:
		log.Printf("Unknown command: %v", commandName)
	}
	return nil, err

}

