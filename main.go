package main

import (
	"bufio"
	"fmt"
	"go_game_server/arithmetic"
	"go_game_server/message"
	"net/http"
	"os"
)

const PI = float64(3.1415926)
const PI1 float64 = 3.1415926
const c0 = iota

const (
	OMG float64 = 2.1
)

type Feed struct {
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Address string `json:"address"`
}

func main() {

	//message.TestFunc()
	//message.ExeCommand()
	//conf.GetRedisConfig()
	//server.Init()
	//web.Init()
	//server.Start()
	//time.Sleep(1000)
	////go grpcclient.Client()
	//wg.Wait()
	//user:=database.UserInfo{
	//	ID:       1,
	//	Name:     "twjitm",
	//	Birthday: "2020",
	//}
	//database.SaveUser(user)
	//
	//message.GetUserList()
	//input()
	//message.CreateThread()
	// message.MapTest()
	// message.SliceTest()
	//message.PointerTest()
	//message.ErrorTest()

	//array := [6]int{12, 212, 12, 12, 12, 1}
	//
	//modify(array)
	//forTest(array)
	//switchTest(2)
	//switchTest(3)
	//switchTest(0)
	//message.Run("twjitm")
	//var types string
	//fmt.Println("hello world")
	//for i := 1; i < len(os.Args); i++ {
	//	types += os.Args[i]
	//}
	//var i, j int32
	//i, j = j, i
	//fmt.Println(PI, OMG)
	//var arrays [5] int32
	//fmt.Println(arrays)
	//
	//searchTeam("twj")
	//request()

	//var array = [3]int{12, 212, 12}
	//var tcp = message.TcpHandler{
	//	Cmd:   12,
	//	Ctime: 12,
	//	MType: 12,
	//	Head:  array,
	//	Body:  nil,
	//}
	//
	//message.Builder(tcp)
	//var udp = message.UdpHandler{
	//	Cmd:  c0,
	//	Head: array,
	//	Body: nil,
	//}
	//message.Builder(&udp)
	//
	//message.GetBody(&udp)

	//web.UnixNet()
	//cluster.EtcdClient.Get("twjitm/*", func(data string) {
	//	fmt.Println(data)
	//})
	//

	//-----------
	//cluster.EtcdClient.Watcher("twjitm/*", func(event *clientv3.Event) {
	//	//
	//})
	//message.TestContrived()
	//for  {
	//	time.Sleep(time.Second*1)
	//}
	//--------------
	// data:=[]int{0,0,1,1,2,3,4}
	//
	//  arithmetic.DuplicatesArray(data)
	//dayN()
	message.TestInterface()

	//images.Image()
   //message.ContextTest()
	message.HandlerTest()
}

func modify(array [6]int) {
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

func getName(age int32, name string) int32 {

	return age
}

func searchTeam(search string) {
	fmt.Println(search)
}

/**
获取键盘输入
*/
func input() {
	var port = 0
	println("请输入端口号")
	_, _ = fmt.Scan(&port)
	println("输入的参数为", port)

	println("请输入一个字符串")
	var reader = bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	println(str)
}

func switchTest(caseVal int32) {
	switch caseVal {
	case 1:
		println("is 1")
		break
	case 2:
		println("is 2")
		break
	case 3:
		fallthrough
	case 4:
		println("is 3 or 4")
		break
	default:
		println("is default", caseVal)
		break
	}
}

func forTest(array [6]int) {
	for key, value := range array {

		println(key)
		println(value)
	}
}

func request() {
	http.HandleFunc("/", dispatch)

	fmt.Println("服务器即将开启，访问地址 http://localhost:7008")

	err := http.ListenAndServe(":7008", nil)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)

	}
}

/**

 */
func dispatch(response http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(response, "这是一个开始")
}


func dayN(){
	searchdata:=[]int{1,4,6,7,8,9,11,13,14,16,18,19}
	arithmetic.HalfSearch(searchdata,18)
}
