/*
 * @Date: 2023-01-21 10:01:21
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-02 15:38:28
 * @FilePath: /simple-DY/DY-api/video-web/api/userregister.go
 * @Description: 1.3.2 用户注册
 */
package api

import (
	"context"
	"net/http"
	"simple-DY/DY-api/video-web/global"
	"simple-DY/DY-api/video-web/middlewares"
	"simple-DY/DY-api/video-web/models"
	pb "simple-DY/DY-api/video-web/proto"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 1.3.2 用户注册 /douyin/user/register/

func UserRegister(c *gin.Context) {
	userRegisterRequest := models.UserRegisterLoginRequest{
		UserName: c.Query("username"),
		Password: c.Query("password"),
	}

	// // 与服务器建立GRPC连接
	// conn := InitGRPC(global.GlobalConfig.GRPC.UserRegisterPort)
	// defer conn.Close()

	// zap.L().Info("服务器端口：" + global.GlobalConfig.GRPC.UserRegisterPort)

	// cpb := pb.NewUserRegisterClient(conn)

	// 将接收到的请求通过GRPC转发给服务端并接收响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()

	responseUserRegister, err := global.UserRegisterSrvClient.UserRegister(ctx, &pb.DouyinUserRegisterRequest{
		Username: userRegisterRequest.UserName,
		Password: userRegisterRequest.Password,
	})
	if err != nil {
		zap.L().Error("GRPC失败！错误信息：" + err.Error())
		return
	}

	zap.L().Info("通过GRPC接收到的响应：" + responseUserRegister.String())

	// 将接收的服务端响应绑定到结构体上
	userRegisterResponse := models.UserRegisterLoginResponse{
		Res: models.ResponseCodeAndMessage{
			StatusCode: responseUserRegister.GetStatusCode(),
			StatusMsg:  responseUserRegister.GetStatusMsg(),
		},
		UserId: responseUserRegister.GetUserId(),
		Token:  middlewares.GenerateToken(responseUserRegister.GetUserId()),
	}

	// 根据不同的返回状态码设置不同的http状态码
	if userRegisterResponse.Res.StatusCode == 0 {
		c.JSON(http.StatusOK, userRegisterResponse)
	} else {
		c.JSON(http.StatusBadRequest, userRegisterResponse)
	}
	zap.L().Info("返回响应成功！")
}
