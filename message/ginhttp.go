package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func requestHandler() {

}

func start() {
	fmt.Println("启动gin服务器")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run("127.0.0.1:7077") // listen and serve on 0.0.0.0:8080
}

func stop() {

}

func request() {

}

func Init() {
	start()
}
