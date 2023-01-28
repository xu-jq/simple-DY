/**
* @Author Wang Hui
* @Description
* @Date
**/
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-DY/DY-api/social-web/global"
	"simple-DY/DY-api/social-web/proto"
	"strconv"
)

func RelationAction(ctx *gin.Context) {

}
func GetFollowList(ctx *gin.Context) {
	userID := ctx.DefaultQuery("userID", "0")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return
	}
	list, err := global.SocialSrvClient.GetFollowList(ctx, &proto.GetFollowListRequest{UserId: int64(id)})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	userList := make([]interface{}, 0)
	if list == nil {
		ctx.JSON(http.StatusOK, userList)
		return
	}
	for _, v := range list.UserList {
		userList = append(userList, v.Id)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "执行成功",
		"user_list":   userList,
	})
}

func GetFollowerList(ctx *gin.Context) {

}
func GetFriendList(ctx *gin.Context) {

}
