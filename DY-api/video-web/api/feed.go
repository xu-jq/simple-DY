/*
 * @Date: 2023-01-19 14:08:05
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-28 21:27:06
 * @FilePath: /simple-DY/DY-api/video-web/api/feed.go
 * @Description: 1.1 视频流接口
 */
package api

import (
	"context"
	"net/http"
	"simple-DY/DY-api/video-web/global"
	"simple-DY/DY-api/video-web/models"
	pb "simple-DY/DY-api/video-web/proto"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 1.1 视频流接口 /douyin/feed/
func Feed(c *gin.Context) {

	responseFeed, err := douyinFeed(c.Query("latest_time"))
	if err != nil {
		zap.L().Error("GRPC失败！错误信息：" + err.Error())
	}

	// 处理接收到的数据
	videolistLen := len(responseFeed.GetVideoList())
	videolist := make([]models.Video, videolistLen)

	responseFeedVideoList := responseFeed.GetVideoList()

	for idx := 0; idx < videolistLen; idx += 1 {

		authorIdInt64 := responseFeedVideoList[idx].GetAuthor().GetId()
		authorIdString := strconv.FormatInt(authorIdInt64, 10)

		// 调用UserInfo服务获取作者名称
		responseUserInfo, err := douyinUser(authorIdString)
		if err != nil {
			zap.L().Error("GRPC失败！错误信息：" + err.Error())
			return
		}
		// Todo 调用其他服务获取关注总数、粉丝总数和UserId与Token解析出来的ID的关注关系
		// Todo 调用其他服务获取视频点赞总数、视频评论总数和Token解析出来的ID是否对这个视频点赞

		videolist[idx].Id = responseFeedVideoList[idx].GetId()
		videolist[idx].Author = models.User{
			Id:            authorIdInt64,
			Name:          responseUserInfo.User.GetName(),
			FollowCount:   1,    // Todo 关注总数
			FollowerCount: 1,    // Todo 粉丝总数
			IsFollow:      true, // Todo 关注关系
		}
		videolist[idx].PlayUrl = responseFeedVideoList[idx].GetPlayUrl()
		videolist[idx].CoverUrl = responseFeedVideoList[idx].GetCoverUrl()
		videolist[idx].FavoriteCount = 1 // Todo 视频点赞总数
		videolist[idx].CommentCount = 1  // Todo 视频评论总数
		videolist[idx].IsFavorite = true // Todo 是否点赞
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

	// 根据不同的返回状态码设置不同的http状态码
	if feedResponse.Res.StatusCode == 0 {
		c.JSON(http.StatusOK, feedResponse)
	} else {
		c.JSON(http.StatusBadRequest, feedResponse)
	}
	zap.L().Info("返回响应成功！")
}

func douyinFeed(latest_time string) (responseFeed *pb.DouyinFeedResponse, err error) {
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
	responseFeed, err = global.FeedSrvClient.Feed(ctx, &pb.DouyinFeedRequest{
		LatestTime: feedRequest.LatestTime,
	})

	zap.L().Info("通过GRPC接收到的响应：" + responseFeed.String())
	return
}
