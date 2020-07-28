package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultPo struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func (result *ResultPo) pushData(key string, data interface{}) {
	if result.Data == nil {
		result.Data = make(map[string]interface{})
	}
	result.Data[key] = data
}

func NewResultPo() ResultPo {
	result := ResultPo{}
	return result
}

var global *gin.Engine

func start() {
	fmt.Println("<<<---------start http server------------>>>")

	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("login", login)
	}
	static := r.Group("/")
	{
		static.GET("login.html", ginHandler)

	}
	//r.StaticFS("/web/", http.Dir("web/static"))
	r.LoadHTMLGlob("web/static/*")
	global = r
	_ = r.Run("127.0.0.1:7077") // listen and serve on 0.0.0.0:7077
}

func ginHandler(context *gin.Context) {

	context.HTML(http.StatusOK, "login.html", nil)
}

func login(context *gin.Context) {
	result := NewResultPo()

	name := context.DefaultQuery("name", "")

	result.Message = "successful"
	result.Code = 200
	result.pushData("username", name)
	fmt.Println(name)
	context.JSON(200, result)
}


func stop() {

}

func request() {

}

func Init() {
	start()
}
