/*
 * @Date: 2023-01-28 20:53:39
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-28 21:01:39
 * @FilePath: /simple-DY/DY-api/social-web/router/relation.go
 * @Description:
 */
/**
* @Author Wang Hui
* @Description
* @Date
**/
package router

import (
	"simple-DY/DY-api/social-web/api"

	"github.com/gin-gonic/gin"
)

func InitRelationRouter(Router *gin.RouterGroup) {
	relationRouter := Router.Group("relation")
	{
		relationRouter.GET("/follow/list", api.GetFollowList)
		relationRouter.GET("/follower/list", api.GetFollowerList)
		relationRouter.GET("/friend/list", api.GetFriendList)
		relationRouter.POST("/action", api.RelationAction)

	}
}
