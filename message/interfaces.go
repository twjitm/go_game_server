package message

import (
	"fmt"
)

type Notification interface {
	notification() error
}

type Logic struct {
	channel int16
}
type Submit struct {
}

func (l *Logic) notification() error {

	fmt.Println("logic channel", l.channel)
	return nil

}
func (s Submit) notification() error {
	fmt.Println("submit")

	return nil
}

func send(notification Notification) {
	_ = notification.notification()

}

func TestInterface() {
	v := duration(1).format(1)
	fmt.Print(v)

}

type duration int

func (d duration) format(val int) string {

	fmt.Print(val)
	return "dsss"
}
