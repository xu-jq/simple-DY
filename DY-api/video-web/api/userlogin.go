/*
 * @Date: 2023-01-21 10:01:21
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-05 19:15:53
 * @FilePath: /simple-DY/DY-api/video-web/api/userlogin.go
 * @Description: 1.3.3 用户登录
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

// 1.3.3 用户登录 /douyin/user/login/

func UserLogin(c *gin.Context) {
	userLoginRequest := models.UserRegisterLoginRequest{
		UserName: c.Query("username"),
		Password: c.Query("password"),
	}

	// // 与服务器建立GRPC连接
	// conn := InitGRPC(global.GlobalConfig.GRPC.UserLoginPort)
	// defer conn.Close()

	// zap.L().Info("服务器端口：" + global.GlobalConfig.GRPC.UserLoginPort)

	// cpb := pb.NewUserLoginClient(conn)

	// 将接收到的请求通过GRPC转发给服务端并接收响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()

	responseUserLogin, err := global.VideoServiceClient.UserLogin(ctx, &pb.DouyinUserLoginRequest{
		Username: userLoginRequest.UserName,
		Password: userLoginRequest.Password,
	})
	if err != nil {
		zap.L().Error("GRPC失败！错误信息：" + err.Error())
		return
	}

	zap.L().Info("通过GRPC接收到的响应：" + responseUserLogin.String())

	// 将接收的服务端响应绑定到结构体上
	userLoginResponse := models.UserRegisterLoginResponse{
		Res: models.ResponseCodeAndMessage{
			StatusCode: responseUserLogin.GetStatusCode(),
			StatusMsg:  responseUserLogin.GetStatusMsg(),
		},
		UserId: responseUserLogin.GetUserId(),
		Token:  middlewares.GenerateToken(responseUserLogin.GetUserId()),
	}

	// 根据不同的返回状态码设置不同的http状态码
	if userLoginResponse.Res.StatusCode == 0 {
		c.JSON(http.StatusOK, userLoginResponse)
	} else {
		c.JSON(http.StatusBadRequest, userLoginResponse)
	}
	zap.L().Info("返回响应成功！")
}
