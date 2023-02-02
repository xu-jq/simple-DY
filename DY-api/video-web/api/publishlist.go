/*
 * @Date: 2023-01-21 10:01:21
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-02 19:26:08
 * @FilePath: /simple-DY/DY-api/video-web/api/publishlist.go
 * @Description: 1.2.1 视频发布列表
 */
package api

import (
	"context"
	"net/http"
	"simple-DY/DY-api/video-web/global"
	"simple-DY/DY-api/video-web/models"
	pb "simple-DY/DY-api/video-web/proto"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 1.2.1 视频发布列表 /douyin/publish/list/

func PublishList(c *gin.Context) {

	idString := c.Query("user_id")

	responsePublishList, err := douyinPublishList(idString)
	if err != nil {
		zap.L().Error("GRPC失败！错误信息：" + err.Error())
	}

	// 处理接收到的数据
	videolistLen := len(responsePublishList.GetVideoList())
	videolist := make([]models.Video, videolistLen)

	responsePublishListVideoList := responsePublishList.GetVideoList()

	// User 信息（因为查询的是发布列表，所以就一个用户信息）
	id, followCount, followerCount, name, _, _, isFollow := userService(c, idString)

	var wgPublishList sync.WaitGroup

	for idx := 0; idx < videolistLen; idx += 1 {
		wgPublishList.Add(1)
		go func(idx int) {
			defer wgPublishList.Done()
			videoIdInt64 := responsePublishListVideoList[idx].GetId()
			videoIdString := strconv.FormatInt(videoIdInt64, 10)
			videolist[idx].Id = videoIdInt64
			videolist[idx].Author = models.User{
				Id:            id,
				Name:          name,
				FollowCount:   followCount,
				FollowerCount: followerCount,
				IsFollow:      isFollow,
			}
			// Video 信息
			favoriteCount, commentCount, isFavorite := videoService(c, videoIdString)
			videolist[idx].PlayUrl = responsePublishListVideoList[idx].GetPlayUrl()
			videolist[idx].CoverUrl = responsePublishListVideoList[idx].GetCoverUrl()
			videolist[idx].FavoriteCount = favoriteCount
			videolist[idx].CommentCount = commentCount
			videolist[idx].IsFavorite = isFavorite
			videolist[idx].Title = responsePublishListVideoList[idx].GetTitle()
		}(idx)
	}
	wgPublishList.Wait()

	// 将接收的服务端响应绑定到结构体上
	publishListResponse := models.PublishListResponse{
		Res: models.ResponseCodeAndMessage{
			StatusCode: responsePublishList.GetStatusCode(),
			StatusMsg:  responsePublishList.GetStatusMsg(),
		},
		VideoList: videolist,
	}

	// 根据不同的返回状态码设置不同的http状态码
	if publishListResponse.Res.StatusCode == 0 {
		c.JSON(http.StatusOK, publishListResponse)
	} else {
		c.JSON(http.StatusBadRequest, publishListResponse)
	}
	zap.L().Info("返回响应成功！")
}

func douyinPublishList(user_id string) (responsePublishList *pb.DouyinPublishListResponse, err error) {
	// 将接收的客户端请求参数绑定到结构体上
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		zap.L().Error("用户id转换为整数失败！错误信息：" + err.Error())
	}
	publishListRequest := models.PublishListRequest{
		UserId: userId,
	}

	// // 与服务器建立GRPC连接
	// conn := InitGRPC(global.GlobalConfig.GRPC.PublishListPort)
	// defer conn.Close()

	// zap.L().Info("服务器端口：" + global.GlobalConfig.GRPC.PublishListPort)

	// cpb := pb.NewPublishListClient(conn)

	// 将接收到的请求通过GRPC转发给服务端并接收响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()
	responsePublishList, err = global.PublishListSrvClient.PublishList(ctx, &pb.DouyinPublishListRequest{
		UserId: publishListRequest.UserId,
	})
	zap.L().Info("通过GRPC接收到的响应：" + responsePublishList.String())
	return
}
