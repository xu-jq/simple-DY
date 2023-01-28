/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-28 23:21:32
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/userregister.go
 * @Description: UserRegister服务
 */
package handler

import (
	"context"
	"net"
	"simple-DY/DY-srvs/video-srv/global"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/consul"
	"simple-DY/DY-srvs/video-srv/utils/dao"
	"simple-DY/DY-srvs/video-srv/utils/jwt"
	"simple-DY/DY-srvs/video-srv/utils/md5salt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
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

	s := grpc.NewServer()
	pb.RegisterUserRegisterServer(s, &userregisterserver{})
	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息：" + err.Error())
	}
	zap.L().Info("服务器监听地址：" + lis.Addr().String())

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	//服务注册
	register_client := consul.NewRegistryClient(global.GlobalConfig.Consul.Address, global.GlobalConfig.Consul.Port)
	register_client.Register("localhost", port, "UserRegister", "UserRegister")

	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息：" + err.Error())
	}
}
