package web

import (
	"github.com/gin-gonic/gin"
	"go_game_server/common"
)

func login(context *gin.Context) {
	result := common.CreateResultPo()
	name := context.DefaultQuery("name", "")
	result.Message = "successful"
	result.Code = 200
	result.PushData("username", name)
	context.Request.Header.Set("session","1111")
	context.JSON(200, result)
}
