package message

import (
	json2 "encoding/json"
	"fmt"
	"strconv"
)

/**
数据结构的一些练习
*/
//--------------------map
/**
map是Go中的内置类型，它将一个值与一个键关联起来。可以使用相应的键检索值。

Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的，也是引用类型
*/

func MapTest() {
	var mapEntry = make(map[string]string)
	mapEntry["name"] = "唐文江"
	mapEntry["go"] = "go开发工程师"
	mapEntry["java"] = "java开发工程师"

	for country := range mapEntry {
		fmt.Println("current of", country, "is", mapEntry[country])
	}
	delete(mapEntry, "age")
	//判断元素是否存在
	val, ok := mapEntry["name"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("false")
	}
	fmt.Println(len(mapEntry))

	newMapEntry := mapEntry
	newMapEntry["age"] = "26"

	fmt.Println("new map=", newMapEntry)
	fmt.Println("map =", mapEntry)

	var abstractMap = make(map[interface{}]interface{})

	abstractMap["string"] = 1
	abstractMap[1] = 2
	abstractMap[newMapEntry] = newMapEntry
	abstractMap[1] = 2

}

/**
                               切片
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
切片是一种方便、灵活且强大的包装器。切片本身没有任何数据。它们只是对现有数组的引用。
切片与数组相比，不需要设定长度，在[]中不用设定值，相对来说比较自由
从概念上面来说slice像一个结构体，这个结构体包含了三个元素：
指针，指向数组中slice指定的开始位置
长度，即slice的长度
最大长度，也就是slice开始位置到数组的最后位置的长度
*/

func SliceTest() {
	sliceEntry := []int{2, 2, 2, 4, 87, 54, 67}
	fmt.Println("切片长度", len(sliceEntry), cap(sliceEntry))

	sliceEntry = append(sliceEntry, 1, 2)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("sliceEntry[1:] ==", sliceEntry[1:])
}

/**
--------指针 fixme 没理解不会用
指针是存储另一个变量的内存地址的变量。
*/
type first struct {
	a    int
	b    bool
	name string
}

func PointerTest() {
	var dataStruct = first{1, false, "twj"}
	var pointer = &dataStruct
	fmt.Println(dataStruct.b, dataStruct.a, dataStruct.name, &dataStruct)
	fmt.Println(&dataStruct.b, &dataStruct.a, &dataStruct.name, dataStruct)
	fmt.Println(pointer.a, &pointer, (*pointer).a)

}

/**

字符串包：strconv，strings
*/
func Strings() {
	strconv.Itoa(1111)
	bo, _ := strconv.ParseBool("true")
	fmt.Println(bo)
}

type UserInfo struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Address  string  `json:"address"`
	Job      JobInfo `json:"job"`
	JobTitle string  `json:"job_title"`
}
type JobInfo struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type Status struct {
	name string `map`
}

func FormatJson() {
	var json = `{"name":"twj","age":27,"address":"北京市昌平区","job":{"title":"go开发工程师","year":3}}`
	by := []byte(json)
	user := UserInfo{}
	_ = json2.Unmarshal(by, &user)
	fmt.Println(user.Name)

}

//函数类型
type Address func(string) string

func (a Address) ToString(str string) string {

	return "this is=" + a(str)
}

func BeiJing(code string) string {

	return "这个地方是:" + code
}

func ShangHai(code string) string {

	return "这个地方是:" + code
}

func GetAddress(address Address, name string) string {

	return address(name)
}
