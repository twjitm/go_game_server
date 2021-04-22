package message

import (
	"fmt"
	"os"
	"sync"
)

func ErrorTest() {
	val, err := os.Open("/data/txt.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val.Name())
	}

}

func DefTest() {
	defer fmt.Println("in main")
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	panic("unknown err")
}

func Defers() error {
	var tasks []Quest
	tasks = append(tasks, Quest{})
	errCh := make(chan error, len(tasks))
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	for i := range tasks {
		go func() {
			defer wg.Done()
			if err := tasks[i].Run(); err != nil {
				errCh <- err
			}
		}()
	}
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}

}

type Quest struct {

}

func (q *Quest) Run() error {

	fmt.Println(1111)
	var err error
	return err
}
