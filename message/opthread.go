package message

import (
	"fmt"
)

/**
go 语言 有关多线程的学习
*/

func Fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

