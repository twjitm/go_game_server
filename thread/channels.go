package thread

import (
	"fmt"
	"time"
)

func GoRun() {
	//cpu := runtime.NumCPU()
	//for i := 0; i <= 8; i++ {
	//	go func() {
	//	}()
	//}
	//fmt.Printf("%d\n", cpu)
	//time.Sleep(time.Second * 12)

	var chanint = make(chan string, 10)

	go func() {
		for i := 1; i < 100; i++ {
			fmt.Println("杀了产品")
			if i == 50 {
				fmt.Println("close will close")
				close(chanint)
				break
			}
			chanint <- "产品"
		}
	}()

	go func() {

		for v := range chanint {
			fmt.Println("救活了" + v)
		}

		//for {
		//	read, ok := <-chanint
		//	if !ok {
		//		fmt.Println("channel is close")
		//		break
		//	}
		//	fmt.Println("救活了" + read)
		//}
	}()

	time.Sleep(time.Second * 5)
}

//channel 读写分离
func GoRunPre() {
	//var readOnly <-chan int = make(chan int, 100)
	// readOnly <- 100 //只读不可写
	//var writeOnly chan<- int = make(chan int, 10)
	// <- writeOnly   //只写不可读
}

/**
channel 作为参数传递信息
*/
type Channel struct {
	status string
}

func GoRunSelect(chans chan int32, name string, myChan chan *Channel) {

	for {
		chans <- 10
		close(chans)
		break
	}
	myChan <- &Channel{
		status: "111",
	}
}
func GoChannelAndSelect() {

	chan1 := make(chan string, 1)
	chan2 := make(chan string, 2)

	chan2 <- "hello go"
	chan2 <- "hello go"
	chan1 <- "hello world"

	select {
	case msg2 := <-chan2:
		fmt.Println("1::::" + msg2)
		break
	case msg1 := <-chan1:
		fmt.Println(msg1)
	default:
		fmt.Println("No data received.")
	}

	time.Sleep(100 * time.Millisecond)
}

func increment(ch chan bool, x *int) {
	select {
	case <-ch:
		*x = *x + 1
	}
}
func ChannelLock() {
	// 注意要设置容量为 1 的缓冲信道
	pipline := make(chan bool, 1)
	var x int
	for i := 0; i < 1000; i++ {
		pipline <- true
		go increment(pipline, &x)
	}
	// 确保所有的协程都已完成
	// 以后会介绍一种更合适的方法（Mutex），这里暂时使用sleep
	time.Sleep(3000)
	fmt.Println("x 的值：", x)
}
