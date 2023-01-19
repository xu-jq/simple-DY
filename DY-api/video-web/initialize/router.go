/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-19 17:02:04
 * @FilePath: /simple-DY/DY-api/video-web/initialize/router.go
 * @Description:
 */
package initialize

import (
	"simple-DY/DY-api/video-web/api"
	"simple-DY/DY-api/video-web/middlewares"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Routers(debug bool) *gin.Engine {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	Router := gin.Default()

	// 开启健康检查
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 1. 基础接口
	DouyinRouter := Router.Group("/douyin")

	// 1.1 视频流接口
	// 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
	DouyinRouter.GET(
		"/feed/",
		api.Feed,
	)

	// // 1.2 视频上传及发布
	// PublishRouter := DouyinRouter.Group("/publish")
	// {
	// 	// 1.2.1 视频发布列表
	// 	// 用户的视频发布列表，直接列出用户所有投稿过的视频
	// 	PublishRouter.GET(
	// 		"/list/",
	// 		jwt.Auth(),
	// 		api.PublishList,
	// 	)

	// 	// 1.2.2 投稿接口
	// 	// 登录用户选择视频上传
	// 	PublishRouter.POST(
	// 		"/action/",
	// 		jwt.AuthBody(),
	// 		api.PublishAction,
	// 	)
	// }

	// // 1.3 用户操作
	// UserRouter := DouyinRouter.Group("/user")
	// {
	// 	// 1.3.1 用户信息
	// 	// 获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数
	// 	UserRouter.GET("/", jwt.Auth(), api.User)

	// 	// 1.3.2 用户注册
	// 	// 新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token
	// 	UserRouter.POST("/register/", api.UserRegister)

	// 	// 1.3.3 用户登录
	// 	// 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token
	// 	UserRouter.POST("/login/", api.UserLogin)

	// }

	//配置跨域
	Router.Use(middlewares.Cors())

	//添加链路追踪
	return Router
}
