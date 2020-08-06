package message

import (
	"context"
	"encoding/json"
	"fmt"
)

func Function() {

}

//函数参数为函数的函数
func FuncFunc(data int, toData func(content context.Context, info *UserInfo)) {
	user := UserInfo{
		Name: "go",
		Age:  data,
	}
	toData(context.TODO(), &user)
}

func TestFunc() {
	var jsonStr = ""
	FuncFunc(10, func(content context.Context, info *UserInfo) {
		e, err := json.Marshal(info)
		if err != nil {
			fmt.Println(err)
			return
		}
		jsonStr = string(e)
	})
	fmt.Println(jsonStr)

}
