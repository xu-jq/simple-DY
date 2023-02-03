/*
 * @Date: 2023-01-21 10:01:21
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-03 10:26:42
 * @FilePath: /simple-DY/DY-api/video-web/api/publishaction.go
 * @Description: 1.2.2 投稿接口
 */
package api

import (
	"context"
	"io"
	"net/http"
	"simple-DY/DY-api/video-web/global"
	"simple-DY/DY-api/video-web/models"
	videopb "simple-DY/DY-api/video-web/proto/video"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 1.2.2 投稿接口 /douyin/publish/action/

func PublishAction(c *gin.Context) {

	// 将接收的客户端请求参数绑定到结构体上
	video, _ := c.FormFile("data")
	videoFile, _ := video.Open()
	videoByte, _ := io.ReadAll(videoFile)

	publishActionRequest := models.PublishActionRequest{
		Data:  videoByte,
		Token: c.PostForm("token"),
		Title: c.PostForm("title"),
	}

	// // 与服务器建立GRPC连接
	// conn := InitGRPC(global.GlobalConfig.GRPC.PublishActionPort)
	// defer conn.Close()

	// zap.L().Info("服务器端口：" + global.GlobalConfig.GRPC.PublishActionPort)

	// cpb := pb.NewPublishActionClient(conn)

	// 将接收到的请求通过GRPC转发给服务端并接收响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.FileSecond))
	defer cancel()

	responsePublishAction, err := global.PublishActionSrvClient.PublishAction(ctx, &videopb.DouyinPublishActionRequest{
		Data:  publishActionRequest.Data,
		Token: publishActionRequest.Token,
		Title: publishActionRequest.Title,
	})

	if err != nil {
		zap.L().Error("GRPC失败！错误信息：" + err.Error())
	}

	zap.L().Info("通过GRPC接收到的响应：" + responsePublishAction.String())

	// 将接收的服务端响应绑定到结构体上
	publishActionResponse := models.PublishActionResponse{
		Res: models.ResponseCodeAndMessage{
			StatusCode: responsePublishAction.GetStatusCode(),
			StatusMsg:  responsePublishAction.GetStatusMsg(),
		},
	}

	// 根据不同的返回状态码设置不同的http状态码
	if publishActionResponse.Res.StatusCode == 0 {
		c.JSON(http.StatusOK, publishActionResponse)
	} else {
		c.JSON(http.StatusBadRequest, publishActionResponse)
	}
	zap.L().Info("返回响应成功！")
}
