/**
* @Author Wang Hui
* @Description
* @Date
**/
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"simple-DY/DY-api/social-web/global"
	"simple-DY/DY-api/social-web/proto"
	"strconv"
)

func MsgChat(ctx *gin.Context) {
	zap.S().Info("-------MsgChat-------")
	userId := ctx.Query("user_id")
	toUserID := ctx.DefaultQuery("to_user_id", "0")
	id, _ := strconv.Atoi(toUserID)
	zap.S().Info("接受的参数to_user_id：", id)
	userid, _ := strconv.Atoi(userId)
	msgChat, err := global.SocialSrvClient.MsgChat(ctx, &proto.MsgChatRequest{
		UserId:   int64(userid),
		ToUserId: int64(id),
	})
	if err != nil {
		zap.S().Error("MsgChat：", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	msgList := make([]interface{}, 0)
	for _, v := range msgChat.MessageList {
		msgList = append(msgList, v)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code":  0,
		"status_msg":   "执行成功",
		"message_list": msgList,
	})

}

func MsgAction(ctx *gin.Context) {
	zap.S().Info("-------MsgAction-------")
	// 接受数据以及表单验证
	userId, err1 := strconv.ParseInt(ctx.GetString("userId"), 10, 64)
	toUserId, err2 := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	actionType, err3 := strconv.ParseInt(ctx.Query("action_type"), 10, 64)
	content := ctx.Query("action_type")
	// fmt.Println(userId, toUserId, actionType)
	// 传入参数格式有问题。
	if nil != err1 || nil != err2 || nil != err3 || actionType < 1 || actionType > 2 {
		fmt.Printf("fail")
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "用户id格式错误",
		})
		return
	}
	zap.S().Info("接受的参数：", userId, toUserId, actionType)
	_, err := global.SocialSrvClient.MsgAction(ctx, &proto.MsgActionRequest{
		UserId:     userId,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
		Content:    content,
	})
	if err != nil {
		zap.S().Error("MsgAction：", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "执行成功",
	})
}
