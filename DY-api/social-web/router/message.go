/**
* @Author Wang Hui
* @Description
* @Date
**/
package router

import (
	"github.com/gin-gonic/gin"
	"simple-DY/DY-api/social-web/api"
	"simple-DY/DY-api/social-web/middlewares"
)

func InitMsgRouter(Router *gin.RouterGroup) {
	msgRouter := Router.Group("/message")
	{
		msgRouter.GET("/chat", middlewares.JWTAuth(), api.MsgChat)
		msgRouter.POST("/action/", middlewares.JWTAuth(), api.MsgAction)
	}
}
