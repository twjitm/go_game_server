package message

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Signals() {
	chSignal := make(chan os.Signal, 1)
	done :=make(chan bool,1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-chSignal:
		fmt.Println(11111)
		signal.Reset(syscall.SIGINT, syscall.SIGTERM)
	case d:=<-done:
		fmt.Print(d)
	}
	fmt.Println("stop world")
	done<-true
}
