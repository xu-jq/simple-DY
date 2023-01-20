/*
 * @Date: 2023-01-19 14:08:05
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-20 19:29:03
 * @FilePath: /simple-DY/DY-api/video-web/api/feed.go
 * @Description: 1.1 视频流接口
 */
package api

import (
	"context"
	"net/http"
	"simple-DY/DY-api/video-web/models"
	pb "simple-DY/DY-api/video-web/proto"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 1.1 视频流接口 /douyin/feed/
func Feed(c *gin.Context) {

	// 将接收的客户端请求参数绑定到结构体上
	latestTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		zap.L().Error("时间戳转换为整数失败！错误信息为：" + err.Error())
	}
	feedRequest := models.FeedRequest{
		LatestTime: latestTime,
		Token:      c.Query("token"),
	}

	// 与服务器建立GRPC连接
	conn := InitGRPC()
	defer conn.Close()

	cpb := pb.NewFeedClient(conn)

	// 将接收到的请求通过GRPC转发给服务端并接收响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	responseFeed, err := cpb.Feed(ctx, &pb.DouyinFeedRequest{
		LatestTime: feedRequest.LatestTime,
		Token:      feedRequest.Token,
	})
	if err != nil {
		zap.L().Error("GRPC失败！错误信息为：" + err.Error())
	}

	zap.L().Info("通过GRPC接收到的响应为：" + responseFeed.String())

	// 处理接收到的数据
	videolistLen := len(responseFeed.GetVideoList())
	videolist := make([]models.Video, videolistLen)

	responseFeedVideoList := responseFeed.GetVideoList()

	for idx := 0; idx < videolistLen; idx += 1 {
		videolist[idx].Id = responseFeedVideoList[idx].GetId()
		videolist[idx].Author = models.User{
			Id:            responseFeedVideoList[idx].GetAuthor().GetId(),
			Name:          responseFeedVideoList[idx].GetAuthor().GetName(),
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
		videolist[idx].PlayUrl = responseFeedVideoList[idx].GetPlayUrl()
		videolist[idx].CoverUrl = responseFeedVideoList[idx].GetCoverUrl()
		videolist[idx].FavoriteCount = 0
		videolist[idx].CommentCount = 0
		videolist[idx].IsFavorite = false
		videolist[idx].Title = responseFeedVideoList[idx].GetTitle()
	}

	// 将接收的服务端响应绑定到结构体上
	feedResponse := models.FeedResponse{
		Res: models.ResponseCodeAndMessage{
			StatusCode: responseFeed.GetStatusCode(),
			StatusMsg:  responseFeed.GetStatusMsg(),
		},
		NextTime:  responseFeed.GetNextTime(),
		VideoList: videolist,
	}

	c.JSON(http.StatusOK, feedResponse)
}
