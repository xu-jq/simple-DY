package initialize

import (
	"github.com/gin-gonic/gin"
	"simple-DY/DY-api/interact-web/middlewares"

	"net/http"
)

func Routers() *gin.Engine {
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
	return Router
}
