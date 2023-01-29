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
	"simple-DY/DY-api/social-web/forms"
	"simple-DY/DY-api/social-web/global"
	"simple-DY/DY-api/social-web/proto"
	"strconv"
)

func MsgChat(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	toUserID := ctx.DefaultQuery("to_user_id", "0")
	id, _ := strconv.Atoi(toUserID)
	zap.S().Info("接受的参数to_user_id：", id)
	msgChat, err := global.SocialSrvClient.MsgChat(ctx, &proto.MsgChatRequest{
		UserId:   userId.(int64),
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
	// 接受数据以及表单验证
	userId, _ := ctx.Get("userId")
	reqForm := forms.MsgActionReq{}
	if err := ctx.ShouldBindJSON(&reqForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("接受的参数：", reqForm)
	_, err := global.SocialSrvClient.MsgAction(ctx, &proto.MsgActionRequest{
		UserId:     userId.(int64),
		ToUserId:   reqForm.ToUserID,
		ActionType: reqForm.ActionType,
		Content:    reqForm.Content,
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
