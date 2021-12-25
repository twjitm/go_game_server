package message

import (
	"fmt"
)

type Animal interface {
	Eat()
	Drink()
}

type Cat struct {
	name string
	sex  string
}

type Dog struct {
}

func (cat *Cat) Eat() {
	fmt.Println("a cat eat name=" + cat.name)
}

func (cat *Cat) Drink() {
	fmt.Println("a cat drink name =" + cat.name)
}

func (c *Cat)Rename() {
	c.name = "小布丁"
}

func Active() {
	//var cat = Cat{name: "小宝贝"}
	//cat.Eat()
	//cat.Rename()
	//cat.Drink()

	//fmt.Println(&cat)
	//	var nulOb *Null
	//	fmt.Println(nulOb == nil)
	//	fmt.Println(IsNilOb(nulOb))
	search()
}

type Null struct {
}

func IsNilOb(v interface{}) bool {
	return v == nil
}

func search(){
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range hash {
		println(k, v)
	}

}


