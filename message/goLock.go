package message

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//sync 包中的类使用和学习

type GoLock struct {
	lock      sync.Mutex
	status    int
	rw        sync.RWMutex
	atomicNum atomic.Value
}

var goLock = GoLock{}

func consume() {
	goLock.lock.Lock()
	goLock.status += 1
	fmt.Println(goLock.status)
	defer goLock.lock.Unlock()
	goLock.rw.RLock()
	fmt.Println("a===", goLock.status)
	defer goLock.rw.RUnlock()
	//只执行一次 的方式
	once := sync.Once{}
	once.Do(func() {
		fmt.Println(1)
	})
}

func atomicOption() {
	goLock.atomicNum.Store(1)
}

//原子变量操作  ASQ CSP
func Test() {
	var value *int32
	var b int32 = 0
	value = &b
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(value, 1)
		}()
	}
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println(*value)
}

type (
	Subscriber chan interface{}         // 订阅者为一个管道
	TopicFunc  func(v interface{}) bool // 主题为一个过滤器
)

//func Getafs() {
//	fs := gatefs.New(vfs.OS("/path"), make(chan bool, 8))
//}
