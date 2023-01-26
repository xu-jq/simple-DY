/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 10:58:24
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/userregister.go
 * @Description: UserRegister服务
 */
package handler

import (
	"context"
	"net"
	"simple-DY/DY-srvs/video-srv/global"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/dao"
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

	// 通过name获取Users表的信息
	user := dao.GetUserByName(in.Username)

	// 如果这个用户已经存在，则不能注册
	if user.Id != 0 {
		userRegisterResponse.StatusCode = 1
		userRegisterResponse.StatusMsg = "用户已经存在！"
		zap.L().Error("用户已经存在！无法注册！已经存在的用户名称：" + user.Name)
		return &userRegisterResponse, nil
	}

	// 将用户密码加密
	password := md5salt.MD5V(in.Password, in.Username, 1)

	// 将用户信息插入数据库
	id := dao.InsertUser(in.Username, password)

	// 构建返回的响应
	userRegisterResponse.StatusCode = 0
	userRegisterResponse.StatusMsg = "注册成功！"
	userRegisterResponse.UserId = id
	userRegisterResponse.Token = jwt.GenerateToken(id)

	zap.L().Info("注册成功！")

	return &userRegisterResponse, nil
}

func UserRegisterService(port string) {
	defer global.Wg.Done()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息：" + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterUserRegisterServer(s, &userregisterserver{})
	zap.L().Info("服务器监听地址：" + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息：" + err.Error())
	}
}
