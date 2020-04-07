package message

import (
	"fmt"
	"os"
)

func ErrorTest() {
	val, err := os.Open("/data/txt.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val.Name())
	}

}
