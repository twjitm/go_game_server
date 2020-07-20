package main

import (
	"app/message"
	"fmt"
)

func init() {
	fmt.Println(" init func")

}
func main() {
	message.FormatJson()
	//thread.GoRun()
	////thread.GoChannelAndSelect()
	//thread.ChannelLock()
	//server.Start()

}
