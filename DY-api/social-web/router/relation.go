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

func InitRelationRouter(Router *gin.RouterGroup) {
	relationRouter := Router.Group("relation")
	{
		relationRouter.GET("/follow/list", api.GetFollowList)
		relationRouter.GET("/follower/list", api.GetFollowerList)
		relationRouter.GET("/friend/list", api.GetFriendList)
		relationRouter.POST("/action", api.RelationAction)

	}
}
