package main

import (
	"app/thread"
	"fmt"
)

func init() {
	fmt.Println(" init func")

}
func main() {
	//thread.GoRun()
	//thread.GoChannelAndSelect()
	thread.ChannelLock()
	//server.Start()

}
