/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 17:25:32
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/feed.go
 * @Description: Feed服务
 */
package handler

import (
	"context"
	"net"
	"simple-DY/DY-srvs/video-srv/global"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/dao"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type feedserver struct {
	pb.UnimplementedFeedServer
}

func (s *feedserver) Feed(ctx context.Context, in *pb.DouyinFeedRequest) (*pb.DouyinFeedResponse, error) {

	// 从数据库中获取指定条件的视频
	result, latestTimeStamp := dao.GetFeedVideos(in.LatestTime, 30)

	zap.L().Info("获取视频流成功！")

	videolistLen := len(result)

	feedResponse := pb.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "获取视频流成功",
		NextTime:   latestTimeStamp["t"].(int64) * 1000,
		VideoList:  make([]*pb.Video, videolistLen),
	}

	urlprefix := global.GlobalConfig.OSS.Address

	for idx := 0; idx < videolistLen; idx += 1 {
		authorId := result[idx]["author_id"].(int64)
		authorIdString := strconv.FormatInt(authorId, 10)
		feedResponse.VideoList[idx] = &pb.Video{
			Id: result[idx]["id"].(int64),
			Author: &pb.User{
				Id: authorId,
			},
			PlayUrl:  urlprefix + authorIdString + global.GlobalConfig.OSS.VideoPath + result[idx]["file_name"].(string) + global.GlobalConfig.OSS.VideoSuffix,
			CoverUrl: urlprefix + authorIdString + global.GlobalConfig.OSS.ImagePath + result[idx]["file_name"].(string) + global.GlobalConfig.OSS.ImageSuffix,
			Title:    result[idx]["title"].(string),
		}
	}

	zap.L().Info("返回响应成功！")

	return &feedResponse, nil
}

func FeedService(port string) {
	defer global.Wg.Done()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息：" + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterFeedServer(s, &feedserver{})
	zap.L().Info("服务器监听地址：" + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息：" + err.Error())
	}
}
