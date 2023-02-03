/*
 * @Date: 2023-01-19 14:08:05
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-03 16:35:29
 * @FilePath: /simple-DY/DY-api/video-web/api/feed.go
 * @Description: 1.1 视频流接口
 */
package api

import (
	"context"
	"net/http"
	"simple-DY/DY-api/video-web/global"
	"simple-DY/DY-api/video-web/models"
	videopb "simple-DY/DY-api/video-web/proto/video"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 1.1 视频流接口 /douyin/feed/
func Feed(c *gin.Context) {

	// 请求服务1：获取视频
	responseFeed, err := douyinFeed(c.Query("latest_time"))
	if err != nil {
		zap.L().Error("GRPC失败！错误信息：" + err.Error())
	}

	// 处理接收到的数据
	videolistLen := len(responseFeed.GetVideoList())
	videolist := make([]models.Video, videolistLen)

	responseFeedVideoList := responseFeed.GetVideoList()

	var wgFeed sync.WaitGroup

	for idx := 0; idx < videolistLen; idx += 1 {
		wgFeed.Add(1)
		go func(idx int) {
			defer wgFeed.Done()
			authorIdInt64 := responseFeedVideoList[idx].GetAuthor().GetId()
			authorIdString := strconv.FormatInt(authorIdInt64, 10)

			// 请求服务2：User 信息
			id, followCount, followerCount, name, _, _, isFollow := userService(c, authorIdString)

			videoIdInt64 := responseFeedVideoList[idx].GetId()
			videoIdString := strconv.FormatInt(videoIdInt64, 10)

			videolist[idx].Id = videoIdInt64
			videolist[idx].Author = models.User{
				Id:            id,
				Name:          name,
				FollowCount:   followCount,
				FollowerCount: followerCount,
				IsFollow:      isFollow,
			}
			// 请求服务3：Video 信息
			favoriteCount, commentCount, isFavorite := videoService(c, videoIdString)
			videolist[idx].PlayUrl = responseFeedVideoList[idx].GetPlayUrl()
			videolist[idx].CoverUrl = responseFeedVideoList[idx].GetCoverUrl()
			videolist[idx].FavoriteCount = favoriteCount
			videolist[idx].CommentCount = commentCount
			videolist[idx].IsFavorite = isFavorite
			videolist[idx].Title = responseFeedVideoList[idx].GetTitle()
		}(idx)
	}
	wgFeed.Wait()

	// 将接收的服务端响应绑定到结构体上
	feedResponse := models.FeedResponse{
		Res: models.ResponseCodeAndMessage{
			StatusCode: responseFeed.GetStatusCode(),
			StatusMsg:  responseFeed.GetStatusMsg(),
		},
		NextTime:  responseFeed.GetNextTime(),
		VideoList: videolist,
	}

	// 根据不同的返回状态码设置不同的http状态码
	if feedResponse.Res.StatusCode == 0 {
		c.JSON(http.StatusOK, feedResponse)
	} else {
		c.JSON(http.StatusBadRequest, feedResponse)
	}
	zap.L().Info("返回响应成功！")
}

func douyinFeed(latest_time string) (responseFeed *videopb.DouyinFeedResponse, err error) {
	// 将接收的客户端请求参数绑定到结构体上
	latestTime, err := strconv.ParseInt(latest_time, 10, 64)
	if err != nil {
		zap.L().Error("时间戳转换为整数失败！错误信息：" + err.Error())
	}
	feedRequest := models.FeedRequest{
		LatestTime: latestTime,
	}

	// // 与服务器建立GRPC连接
	// conn := InitGRPC(global.GlobalConfig.GRPC.FeedPort)
	// defer conn.Close()

	// zap.L().Info("服务器端口：" + global.GlobalConfig.GRPC.FeedPort)

	// cpb := pb.NewFeedClient(conn)

	// 将接收到的请求通过GRPC转发给服务端并接收响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()
	responseFeed, err = global.FeedSrvClient.Feed(ctx, &videopb.DouyinFeedRequest{
		LatestTime: feedRequest.LatestTime,
	})

	zap.L().Info("通过GRPC接收到的响应：" + responseFeed.String())
	return
}
