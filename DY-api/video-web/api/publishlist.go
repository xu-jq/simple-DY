/*
 * @Date: 2023-01-21 10:01:21
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-22 21:30:19
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
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 1.2.1 视频发布列表 /douyin/publish/list/

func PublishList(c *gin.Context) {

	// 将接收的客户端请求参数绑定到结构体上
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		zap.L().Error("用户id转换为整数失败！错误信息为：" + err.Error())
	}
	publishListRequest := models.UserRequest{
		UserId: userId,
		Token:  c.Query("token"),
	}

	// 与服务器建立GRPC连接
	conn := InitGRPC(global.GlobalConfig.GRPCServerPublishListPort)
	defer conn.Close()

	zap.L().Info("服务器端口为：" + global.GlobalConfig.GRPCServerPublishListPort)

	cpb := pb.NewPublishListClient(conn)

	// 将接收到的请求通过GRPC转发给服务端并接收响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	responsePublishList, err := cpb.PublishList(ctx, &pb.DouyinPublishListRequest{
		UserId: publishListRequest.UserId,
		Token:  publishListRequest.Token,
	})
	if err != nil {
		zap.L().Error("GRPC失败！错误信息为：" + err.Error())
	}

	zap.L().Info("通过GRPC接收到的响应为：" + responsePublishList.String())

	// 处理接收到的数据
	videolistLen := len(responsePublishList.GetVideoList())
	videolist := make([]models.Video, videolistLen)

	responsePublishListVideoList := responsePublishList.GetVideoList()

	for idx := 0; idx < videolistLen; idx += 1 {
		videolist[idx].Id = responsePublishListVideoList[idx].GetId()
		videolist[idx].Author = models.User{
			Id:            responsePublishListVideoList[idx].GetAuthor().GetId(),
			Name:          responsePublishListVideoList[idx].GetAuthor().GetName(),
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
		videolist[idx].PlayUrl = responsePublishListVideoList[idx].GetPlayUrl()
		videolist[idx].CoverUrl = responsePublishListVideoList[idx].GetCoverUrl()
		videolist[idx].FavoriteCount = 0
		videolist[idx].CommentCount = 0
		videolist[idx].IsFavorite = false
		videolist[idx].Title = responsePublishListVideoList[idx].GetTitle()
	}

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
