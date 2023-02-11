package api

import (
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
	resp, err := global.InteractSrvClient.FavoriteAction(ctx, &proto.DouyinFavoriteActionRequest{
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
	uId := ctx.GetInt64("user_id")
	token := ctx.GetString("token")
	if token == "" || uId == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "参数异常",
		})
		return
	}
	resp, err := global.InteractSrvClient.GetFavoriteList(ctx, &proto.DouyinFavoriteListRequest{
		UserId: uId,
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