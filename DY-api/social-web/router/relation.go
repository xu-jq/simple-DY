/*
 * @Date: 2023-01-28 20:53:39
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-28 21:01:39
 * @FilePath: /simple-DY/DY-api/social-web/router/relation.go
 * @Description:
 */
package router

import (
	"github.com/gin-gonic/gin"
	"simple-DY/DY-api/social-web/api"
	"simple-DY/DY-api/social-web/middlewares"
)

func InitRelationRouter(Router *gin.RouterGroup) {
	relationRouter := Router.Group("/relation")

	{
		relationRouter.GET("/follow/list/", middlewares.JWTAuth(), api.GetFollowList)
		relationRouter.GET("/follower/list/", middlewares.JWTAuth(), api.GetFollowerList)
		// friend 就是关注的用户
		relationRouter.GET("/friend/list/", middlewares.JWTAuth(), api.GetFriendList)
		relationRouter.POST("/action/", middlewares.JWTAuth(), api.RelationAction)

	}
}
