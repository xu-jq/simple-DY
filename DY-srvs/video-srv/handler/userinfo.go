/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-28 23:21:08
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/userinfo.go
 * @Description: UserInfo服务
 */
package handler

import (
	"context"
	"net"
	"simple-DY/DY-srvs/video-srv/global"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/consul"
	"simple-DY/DY-srvs/video-srv/utils/dao"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type userinfoserver struct {
	pb.UnimplementedUserInfoServer
}

func (s *userinfoserver) UserInfo(ctx context.Context, in *pb.DouyinUserRequest) (*pb.DouyinUserResponse, error) {

	// 构建返回的响应
	userResponse := pb.DouyinUserResponse{}

	// 通过id获取Users表的信息
	user := dao.GetUserById(in.UserId)

	// 如果这个用户不存在，则不能返回信息
	if user.Name == "" {
		zap.L().Error("用户不存在！")
		userResponse.StatusCode = 2
		userResponse.StatusMsg = "用户不存在！"
		return &userResponse, nil
	}

	// 返回响应
	userResponse.StatusMsg = "成功获取用户信息"
	userResponse.User = &pb.User{
		Id:   user.Id,
		Name: user.Name,
	}
	zap.L().Info("成功获取用户信息！")

	return &userResponse, nil
}

func UserInfoService(port string) {
	defer global.Wg.Done()

	s := grpc.NewServer()
	pb.RegisterUserInfoServer(s, &userinfoserver{})
	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息：" + err.Error())
	}
	zap.L().Info("服务器监听地址：" + lis.Addr().String())

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	//服务注册
	register_client := consul.NewRegistryClient(global.GlobalConfig.Consul.Address, global.GlobalConfig.Consul.Port)
	register_client.Register("localhost", port, "UserInfo", "UserInfo")

	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息：" + err.Error())
	}
}
