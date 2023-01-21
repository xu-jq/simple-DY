/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-21 12:23:47
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/userinfo.go
 * @Description: UserInfo服务
 */
package handler

import (
	"context"
	"log"
	"net"
	"simple-DY/DY-srvs/video-srv/global"
	pb "simple-DY/DY-srvs/video-srv/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type userinfoserver struct {
	pb.UnimplementedUserInfoServer
}

func (s *userinfoserver) UserInfo(ctx context.Context, in *pb.DouyinUserRequest) (*pb.DouyinUserResponse, error) {

	zap.L().Error("待补充")

	return &pb.DouyinUserResponse{}, nil
}

func UserInfoService(port string) {
	defer global.Wg.Done()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息为：" + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterUserInfoServer(s, &userinfoserver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息为：" + err.Error())
	}
}
