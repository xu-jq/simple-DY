package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"simple-DY/DY-api/interact-web/global"
	"simple-DY/DY-api/interact-web/proto"
)

func CommentAction(ctx *gin.Context) {
	token := ctx.GetString("token")
	vId := ctx.GetInt64("video_id")
	actionType := ctx.GetInt("action_type")
	if token == "" || vId == 0 || actionType == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "参数异常",
		})
		return
	}
	if actionType == 1 {
		commentText := ctx.GetString("comment_text")
		if commentText == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "参数异常",
			})
			return
		}
		resp, err := global.InteractSrvClient.CommentAction(ctx, &proto.DouyinCommentActionRequest{
			Token:       token,
			VideoId:     vId,
			ActionType:  int32(actionType),
			CommentText: commentText,
		})
		if err != nil {
			zap.S().Error("CommentAction：", err)
			HandleGrpcErrorToHttp(err, ctx)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": resp.StatusCode,
			"status_msg":  resp.StatusMsg,
			"comment":     resp.Comment,
		})
	} else if actionType == 2 {
		commentId := ctx.GetInt64("comment_id")
		if commentId == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "参数异常",
			})
			return
		}
		resp, err := global.InteractSrvClient.CommentAction(ctx, &proto.DouyinCommentActionRequest{
			Token:      token,
			VideoId:    vId,
			ActionType: int32(actionType),
			CommentId:  commentId,
		})
		if err != nil {
			zap.S().Error("CommentAction：", err)
			HandleGrpcErrorToHttp(err, ctx)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": resp.StatusCode,
			"status_msg":  resp.StatusMsg,
			"comment":     resp.Comment,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "参数异常",
		})
		return
	}
}

func CommentList(ctx *gin.Context) {
	token := ctx.GetString("token")
	vId := ctx.GetInt64("video_id")
	if token == "" || vId == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "参数异常",
		})
		return
	}
	resp, err := global.InteractSrvClient.GetCommentList(ctx, &proto.DouyinCommentListRequest{
		Token:   token,
		VideoId: vId,
	})
	if err != nil {
		zap.S().Error("CommentList：", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code":  resp.StatusCode,
		"status_msg":   resp.StatusMsg,
		"comment_list": resp.CommentList,
	})
}
