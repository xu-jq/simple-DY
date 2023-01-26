/*
 * @Date: 2023-01-21 10:01:21
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 17:15:28
 * @FilePath: /simple-DY/DY-api/video-web/api/userinfo.go
 * @Description: 1.3.1 用户信息
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

// 1.3.1 用户信息

func UserInfo(c *gin.Context) {

	responseUserInfo, err := douyinUser(c.Query("user_id"))
	if err != nil {
		zap.L().Error("GRPC失败！错误信息：" + err.Error())
		return
	}

	// Todo 调用其他服务获取关注总数、粉丝总数和UserId与Token解析出来的ID的关注关系

	// 将接收的服务端响应绑定到结构体上
	userResponse := models.UserResponse{
		Res: models.ResponseCodeAndMessage{
			StatusCode: responseUserInfo.GetStatusCode(),
			StatusMsg:  responseUserInfo.GetStatusMsg(),
		},
		User: models.User{
			Id:            responseUserInfo.User.GetId(),
			Name:          responseUserInfo.User.GetName(),
			FollowCount:   1,    // Todo 关注总数
			FollowerCount: 1,    // Todo 粉丝总数
			IsFollow:      true, // Todo 关注关系
		},
	}

	// 根据不同的返回状态码设置不同的http状态码
	if userResponse.Res.StatusCode == 0 {
		c.JSON(http.StatusOK, userResponse)
	} else {
		c.JSON(http.StatusBadRequest, userResponse)
	}
	zap.L().Info("返回响应成功！")
}

func douyinUser(user_id string) (responseUserInfo *pb.DouyinUserResponse, err error) {
	// 将接收的客户端请求参数绑定到结构体上
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		zap.L().Error("用户id转换为整数失败！错误信息：" + err.Error())
		return
	}
	userRequest := models.UserRequest{
		UserId: userId,
	}

	// 与服务器建立GRPC连接
	conn := InitGRPC(global.GlobalConfig.GRPC.UserInfoPort)
	defer conn.Close()

	zap.L().Info("服务器端口：" + global.GlobalConfig.GRPC.UserInfoPort)

	cpb := pb.NewUserInfoClient(conn)

	// 将接收到的请求通过GRPC转发给服务端并接收响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()
	responseUserInfo, err = cpb.UserInfo(ctx, &pb.DouyinUserRequest{
		UserId: userRequest.UserId,
	})
	zap.L().Info("通过GRPC接收到的响应：" + responseUserInfo.String())
	return
}
