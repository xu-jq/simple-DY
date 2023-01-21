/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-21 12:24:19
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/userregister.go
 * @Description: UserRegister服务
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

type userregisterserver struct {
	pb.UnimplementedUserRegisterServer
}

func (s *userregisterserver) UserRegister(ctx context.Context, in *pb.DouyinUserRegisterRequest) (*pb.DouyinUserRegisterResponse, error) {

	zap.L().Error("待补充")

	return &pb.DouyinUserRegisterResponse{}, nil
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
