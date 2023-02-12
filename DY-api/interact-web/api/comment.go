package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"simple-DY/DY-api/interact-web/global"
	"simple-DY/DY-api/interact-web/proto"
	"strconv"
)

func CommentAction(ctx *gin.Context) {
	token := ctx.Query("token")
	vId := ctx.Query("video_id")
	zap.S().Info("token:", token)
	zap.S().Info("video_id:", vId)
	actionType := ctx.Query("action_type")
	aType, _ := strconv.Atoi(actionType)
	if token == "" || vId == "" || aType == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "参数异常",
		})
		return
	}
	vid, _ := strconv.Atoi(vId)
	if aType == 1 {
		commentText := ctx.Query("comment_text")
		if commentText == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "参数异常",
			})
			return
		}
		resp, err := global.InteractSrvClient.CommentAction(ctx, &proto.DouyinCommentActionRequest{
			Token:       token,
			VideoId:     int64(vid),
			ActionType:  int32(aType),
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
	} else if aType == 2 {
		commentId := ctx.Query("comment_id")
		cId, _ := strconv.Atoi(commentId)
		if cId == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "参数异常",
			})
			return
		}
		resp, err := global.InteractSrvClient.CommentAction(ctx, &proto.DouyinCommentActionRequest{
			Token:      token,
			VideoId:    int64(vid),
			ActionType: int32(aType),
			CommentId:  int64(cId),
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
	token := ctx.Query("token")
	vId := ctx.Query("video_id")
	zap.S().Info("token:", token)
	zap.S().Info("video_id:", vId)
	videoId, _ := strconv.Atoi(vId)
	if token == "" || videoId == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "参数异常",
		})
		return
	}
	resp, err := global.InteractSrvClient.GetCommentList(ctx, &proto.DouyinCommentListRequest{
		Token:   token,
		VideoId: int64(videoId),
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
