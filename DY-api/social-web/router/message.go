/**
* @Author Wang Hui
* @Description
* @Date
**/
package router

import (
	"github.com/gin-gonic/gin"
	"simple-DY/DY-api/social-web/api"
)

func InitMsgRouter(Router *gin.RouterGroup) {
	msgRouter := Router.Group("/message")
	{
		msgRouter.GET("/chat", api.MsgChat)
		msgRouter.POST("/action", api.MsgAction)
	}
}
