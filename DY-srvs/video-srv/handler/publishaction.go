/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-22 21:37:14
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/publishaction.go
 * @Description: PublishAction服务
 */
package handler

import (
	"context"
	"fmt"
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

	zap.L().Error("接收到的响应如下")
	fmt.Println(in.Title)
	fmt.Println(in.Token)
	fmt.Println(in.Data)
	zap.L().Error("待补充")

	publishActionResponse := pb.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "作者投稿视频上传成功",
	}

	zap.L().Info("返回响应成功！")

	return &publishActionResponse, nil
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
