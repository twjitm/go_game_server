package message

import (
	"fmt"
	"os/exec"
)

//go 执行脚本
func ExeCommand() {

	exe := exec.Command("/bin/bash", "-c", "ps aux|grep java")
	result, ero := exe.Output()
	if ero != nil {
		fmt.Println(ero.Error())
	}
	s := string(result)
	fmt.Println(s)
}
