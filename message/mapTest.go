package message

import "fmt"

type User struct {
	UName string
	ID    int64
	Age   int
}

func GetUserList() User {

	maps := make(map[int]User)

	maps[1] = User{
		UName: "twj",
		ID:    1,
		Age:   26,
	}
	maps[2] = User{
		UName: "xlp",
		ID:    2,
		Age:   25,
	}

	for k, v := range maps {
		fmt.Println(k)
		fmt.Println(v)
	}
	re, ok := maps[1]
	if ok {
		fmt.Println(re)
	}

	return maps[1]
}

//线性安全操作map
