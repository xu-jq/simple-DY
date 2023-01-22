/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-22 21:00:36
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/userregister.go
 * @Description: UserRegister服务
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

type userregisterserver struct {
	pb.UnimplementedUserRegisterServer
}

func (s *userregisterserver) UserRegister(ctx context.Context, in *pb.DouyinUserRegisterRequest) (*pb.DouyinUserRegisterResponse, error) {

	// 构建返回的响应
	userRegisterResponse := pb.DouyinUserRegisterResponse{}

	// 数据库查询和更新的模板
	user := models.Users{}

	// 根据姓名查找数据库中的用户信息
	global.DB.Where("name = ?", in.Username).Find(&user)

	// 如果这个用户已经存在，则不能注册
	if user.Id != 0 {
		userRegisterResponse.StatusCode = 1
		userRegisterResponse.StatusMsg = "用户已经存在！"
		zap.L().Error("用户已经存在！无法注册！已经存在的用户名称为：" + user.Name)
	} else {
		// 将用户密码加密
		password := md5salt.MD5V(in.Password, in.Username, 1)

		// 将用户信息插入数据库
		user.Name = in.Username
		user.Password = password
		global.DB.Create(&user)

		// 构建返回的响应
		userRegisterResponse.StatusCode = 0
		userRegisterResponse.StatusMsg = "注册成功！"
		userRegisterResponse.UserId = user.Id
		userRegisterResponse.Token = jwt.GenerateToken(user.Id)
		zap.L().Info("注册成功！")
	}
	zap.L().Info("返回响应成功！")

	return &userRegisterResponse, nil
}

func UserRegisterService(port string) {
	defer global.Wg.Done()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息为：" + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterUserRegisterServer(s, &userregisterserver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息为：" + err.Error())
	}
}
