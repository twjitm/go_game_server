package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_game_server/service/sso"
	"net/http"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		ts := c.Query("ts") // 时间戳
		fmt.Println(ts)
		appid := c.Request.Header.Get("appid")
		token := c.Request.Header.Get("token") // 访问令牌
		if appid == "" || token == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "appid or token error"})
			return
		}
		if sso.CertService.Valid(appid, token) {
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			return
		}
	}
}
