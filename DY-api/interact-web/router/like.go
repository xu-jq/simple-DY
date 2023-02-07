package router

import (
	"github.com/gin-gonic/gin"
	"simple-DY/DY-api/interact-web/api"
	"simple-DY/DY-api/interact-web/middlewares"
)

func InitLikeRouter(Router *gin.RouterGroup) {
	msgRouter := Router.Group("/favorite")
	{
		msgRouter.POST("/action/", middlewares.JWTAuth(), api.LikeAction)
		msgRouter.GET("/list/", middlewares.JWTAuth(), api.LikeList)
	}
}
