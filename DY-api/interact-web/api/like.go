package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"simple-DY/DY-api/interact-web/global"
	"simple-DY/DY-api/interact-web/proto"
	"strconv"
)

func LikeAction(ctx *gin.Context) {
	vid := ctx.Query("video_id")
	aType := ctx.Query("action_type")
	videoId, _ := strconv.Atoi(vid)
	actionType, _ := strconv.Atoi(aType)
	token := ctx.Query("token")
	if videoId == 0 || token == "" || actionType == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "参数异常",
		})
		return
	}
	resp, err := global.InteractSrvClient.FavoriteAction(context.Background(), &proto.DouyinFavoriteActionRequest{
		Token:      ctx.Query("token"),
		VideoId:    int64(videoId),
		ActionType: int32(actionType),
	})
	if err != nil {
		zap.S().Error("LikeAction：", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
	})
}

func LikeList(ctx *gin.Context) {
	uId := ctx.Query("user_id")
	token := ctx.Query("token")
	if token == "" || uId == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "参数异常",
		})
		return
	}
	userId, _ := strconv.Atoi(uId)
	resp, err := global.InteractSrvClient.GetFavoriteList(context.Background(), &proto.DouyinFavoriteListRequest{
		UserId: int64(userId),
		Token:  token,
	})
	if err != nil {
		zap.S().Error("LikeList：", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"video_list":  resp.VideoList,
	})
}
