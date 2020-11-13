package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_game_server/database"
)

var global *gin.Engine

func start() {
	fmt.Println("<<<---------start http server------------>>>")
	r := gin.Default()
	r.Use(Authorize())
	v1 := r.Group("v1")
	{
		v1.GET("login", login)
	}
	global = r
	_ = r.Run("127.0.0.1:7077")
}

func stop() {

}

func request() {

}

func Init() {
	database.SetupMongoDB()
	start()
}
