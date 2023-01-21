/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-21 12:23:15
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/publishaction.go
 * @Description: PublishAction服务
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

type publishactionserver struct {
	pb.UnimplementedPublishActionServer
}

func (s *publishactionserver) PublishAction(ctx context.Context, in *pb.DouyinPublishActionRequest) (*pb.DouyinPublishActionResponse, error) {

	zap.L().Error("待补充")

	return &pb.DouyinPublishActionResponse{}, nil
}

func PublishActionService(port string) {
	defer global.Wg.Done()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息为：" + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterPublishActionServer(s, &publishactionserver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息为：" + err.Error())
	}
}
