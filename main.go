package main

import (
	"fmt"
	"os"
)

const PI = float64(3.1415926)
const PI1 float64 = 3.1415926
const c0 = iota

const (
	OMG float64 = 2.1
)

func main() {
	array := [6]int{12, 212, 12, 12, 12, 1}
	modify(array)

	var types string
	fmt.Println("hello world")
	for i := 1; i < len(os.Args); i++ {
		types += os.Args[i]
	}
	var i, j int32
	i, j = j, i
	fmt.Println(PI, OMG)
	var arrays [5] int32
	fmt.Println(arrays)

}

func modify(array [6] int) {
	array[1] = 1
	fmt.Println(array)
}

func gettype() {
	var types int32
	fmt.Println(types)
	str := "hell"
	lenI := len(str)
	fmt.Println(lenI)
}
func getName(age int32, name string) {

}
