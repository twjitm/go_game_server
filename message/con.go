package message

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func ContextTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	go SlowOperation(ctx)
	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("goroutine:", runtime.NumGoroutine())
		}
	}()
	//time.Sleep(4 * time.Second)
	select {
	case result := <-result():
		fmt.Println(result)
	}

}

func SlowOperation(ctx context.Context) {
	done := make(chan struct {
		Name string
	})
	go func() { // 模拟慢操作
		dur := time.Duration(rand.Intn(5)+1) * time.Second
		time.Sleep(dur)
		done <- struct{ Name string }{Name: "twj"}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("SlowOperation timeout:", ctx.Err())
	case v := <-done:
		fmt.Println("v==", v)
		fmt.Println("Complete work")
	}
}

func result() chan string {
	res := make(chan string)
	go func() {
		res <- "funlpus"
	}()
	return res
}

type MateData struct {
	Values     []interface{}
	Handler    func([]int) int
	downstream chan interface{}
}

func handl(data chan MateData) {
	for datum := range data {
		hand := datum.Handler
		arr := []int{1, 2, 4}
		if hand(arr) > 0 {
			fmt.Println("handler>>>>")
		}
		datum.downstream <- "downstream data"
	}
}

func HandlerTest() {
	data := make(chan MateData)
	mate := MateData{
		Values: nil,
		Handler: func(ints []int) int {
			num := 0
			for i := 0; i < len(ints); i++ {
				num += ints[i]
			}
			return num
		},
		downstream: make(chan interface{}),
	}
	go func() {
		data <- mate
	}()
	go handl(data)
	select {
	case <-data:
		fmt.Println("data")
		break
	case d := <-mate.downstream:
		fmt.Println( d)
	}
	time.Sleep(time.Duration(2)*time.Second)

}
