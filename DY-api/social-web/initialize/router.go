package initialize

import (
	"github.com/gin-gonic/gin"
	"simple-DY/DY-api/social-web/middlewares"
	"simple-DY/DY-api/social-web/router"

	"net/http"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	// 开启健康检查
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	Router.Use(middlewares.Cors())
	//添加链路追踪
	group := Router.Group("/douyin")
	router.InitMsgRouter(group)
	router.InitRelationRouter(group)
	return Router
}
