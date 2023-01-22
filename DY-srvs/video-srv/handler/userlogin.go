/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-22 21:00:31
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/userlogin.go
 * @Description: UserLogin服务
 */
package handler

import (
	"context"
	"log"
	"net"
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/models"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/jwt"
	"simple-DY/DY-srvs/video-srv/utils/md5salt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type userloginserver struct {
	pb.UnimplementedUserLoginServer
}

func (s *userloginserver) UserLogin(ctx context.Context, in *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {

	// 构建返回的响应
	userLoginResponse := pb.DouyinUserLoginResponse{}

	// 数据库查询和更新的模板
	user := models.Users{}

	// 根据姓名查找数据库中的用户信息
	global.DB.Where("name = ?", in.Username).Find(&user)

	// 如果这个用户不存在，则不能登录
	if user.Id == 0 {
		userLoginResponse.StatusCode = 1
		userLoginResponse.StatusMsg = "用户不存在！"
		zap.L().Error("用户不存在！无法登录！用户名称为：" + user.Name)
	} else {
		// 将用户密码加密
		password := md5salt.MD5V(in.Password, in.Username, 1)

		// 验证密码是否正确
		if user.Password != password {
			userLoginResponse.StatusCode = 1
			userLoginResponse.StatusMsg = "密码错误！"
			zap.L().Error("密码错误！无法登录！")
		} else {
			// 构建返回的响应
			userLoginResponse.StatusCode = 0
			userLoginResponse.StatusMsg = "登录成功！"
			userLoginResponse.UserId = user.Id
			userLoginResponse.Token = jwt.GenerateToken(user.Id)
			zap.L().Info("用户名和密码正确！登录成功！")
		}
	}
	zap.L().Info("返回响应成功！")

	return &userLoginResponse, nil
}

func UserLoginService(port string) {
	defer global.Wg.Done()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息为：" + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterUserLoginServer(s, &userloginserver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息为：" + err.Error())
	}
}
