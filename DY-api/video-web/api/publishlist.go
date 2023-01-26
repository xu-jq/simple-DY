/*
 * @Date: 2023-01-21 10:01:21
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 17:14:57
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

	responsePublishList, err := douyinPublishList(c.Query("user_id"))
	if err != nil {
		zap.L().Error("GRPC失败！错误信息：" + err.Error())
	}

	// Todo 调用其他服务获取关注总数、粉丝总数和UserId与Token解析出来的ID的关注关系

	// 处理接收到的数据
	videolistLen := len(responsePublishList.GetVideoList())
	videolist := make([]models.Video, videolistLen)

	responsePublishListVideoList := responsePublishList.GetVideoList()

	for idx := 0; idx < videolistLen; idx += 1 {

		// Todo 调用其他服务获取视频点赞总数、视频评论总数和Token解析出来的ID是否对这个视频点赞

		videolist[idx].Id = responsePublishListVideoList[idx].GetId()
		videolist[idx].Author = models.User{
			Id:            responsePublishListVideoList[idx].GetAuthor().GetId(),
			Name:          responsePublishListVideoList[idx].GetAuthor().GetName(),
			FollowCount:   1,    // Todo 关注总数
			FollowerCount: 1,    // Todo 粉丝总数
			IsFollow:      true, // Todo 关注关系
		}
		videolist[idx].PlayUrl = responsePublishListVideoList[idx].GetPlayUrl()
		videolist[idx].CoverUrl = responsePublishListVideoList[idx].GetCoverUrl()
		videolist[idx].FavoriteCount = 1 // Todo 视频点赞总数
		videolist[idx].CommentCount = 1  // Todo 视频评论总数
		videolist[idx].IsFavorite = true // Todo 是否点赞
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

func douyinPublishList(user_id string) (responsePublishList *pb.DouyinPublishListResponse, err error) {
	// 将接收的客户端请求参数绑定到结构体上
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		zap.L().Error("用户id转换为整数失败！错误信息：" + err.Error())
	}
	publishListRequest := models.PublishListRequest{
		UserId: userId,
	}

	// 与服务器建立GRPC连接
	conn := InitGRPC(global.GlobalConfig.GRPC.PublishListPort)
	defer conn.Close()

	zap.L().Info("服务器端口：" + global.GlobalConfig.GRPC.PublishListPort)

	cpb := pb.NewPublishListClient(conn)

	// 将接收到的请求通过GRPC转发给服务端并接收响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()
	responsePublishList, err = cpb.PublishList(ctx, &pb.DouyinPublishListRequest{
		UserId: publishListRequest.UserId,
	})
	zap.L().Info("通过GRPC接收到的响应：" + responsePublishList.String())
	return
}
