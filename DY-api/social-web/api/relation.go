/**
* @Author Wang Hui
* @Description
* @Date
**/
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"simple-DY/DY-api/social-web/global"
	"simple-DY/DY-api/social-web/proto"
	"strconv"
)

func RelationAction(ctx *gin.Context) {
	zap.S().Info("-------RelationAction-------")
	// auth中间件解析后，将userId存入ctx中。
	userId, _ := ctx.Get("TokenId")
	id, _ := strconv.Atoi(userId.(string))
	toUserId, _ := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(ctx.Query("action_type"), 10, 64)
	zap.S().Info("接受的参数：", userId, toUserId, actionType)
	_, err := global.SocialSrvClient.RelationAction(ctx, &proto.RelationActionRequest{
		UserId:     int64(id),
		ToUserId:   toUserId,
		ActionType: int32(actionType),
	})
	if err != nil {
		zap.S().Error("RelationAction：", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "执行成功",
	})
}

func GetFollowList(ctx *gin.Context) {
	zap.S().Info("-------GetFollowList-------")
	userID := ctx.Query("user_id")
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
		zap.S().Error("GetFollowList：", err)
		ctx.JSON(http.StatusOK, userList)
		return
	}
	for _, v := range list.UserList {
		userList = append(userList, v)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "执行成功",
		"user_list":   userList,
	})
}

func GetFollowerList(ctx *gin.Context) {
	zap.S().Info("-------GetFollowerList-------")
	userID := ctx.Query("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return
	}
	list, err := global.SocialSrvClient.GetFollowerList(ctx, &proto.FollowerListRequest{UserId: int64(id)})
	if err != nil {
		zap.S().Error("GetFollowerList：", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	userList := make([]interface{}, 0)
	if list == nil {
		ctx.JSON(http.StatusOK, userList)
		return
	}
	for _, v := range list.UserList {
		userList = append(userList, v)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "执行成功",
		"user_list":   userList,
	})
}

func GetFriendList(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return
	}
	list, err := global.SocialSrvClient.GetFriendList(ctx, &proto.GetFriendListRequest{UserId: int64(id)})
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
		userList = append(userList, v)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "执行成功",
		"user_list":   userList,
	})
}
