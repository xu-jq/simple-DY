package router

import (
	"github.com/gin-gonic/gin"
	"simple-DY/DY-api/interact-web/api"
	"simple-DY/DY-api/interact-web/middlewares"
)

func InitCommentRouter(Router *gin.RouterGroup) {
	msgRouter := Router.Group("/comment")
	{
		msgRouter.POST("/action/", middlewares.JWTAuth(), api.CommentAction)
		msgRouter.GET("/list/", middlewares.JWTAuth(), api.CommentList)
	}
}
