/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 16:04:02
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/feed.go
 * @Description: Feed服务
 */
package handler

import (
	"context"
	"log"
	"net"
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/models"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"strconv"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type feedserver struct {
	pb.UnimplementedFeedServer
}

func (s *feedserver) Feed(ctx context.Context, in *pb.DouyinFeedRequest) (*pb.DouyinFeedResponse, error) {

	// 获取请求的时间戳
	timeStamp := in.GetLatestTime() / 1000
	zap.L().Info("此请求的时间是：" + time.Unix(timeStamp, 0).Format(global.GlobalConfig.Time.TimeFormat))

	// 查询前30个视频
	videoQuery := global.DB.Model(&models.Videos{}).Where("videos.publish_time < " + strconv.FormatInt(timeStamp, 10)).Order("publish_time DESC").Limit(30)

	// 查询前30个视频的最早时间
	latestTimeStamp := map[string]interface{}{}
	global.DB.Table("(?) as u", videoQuery).Select("publish_time as t").Order("publish_time ASC").Limit(1).Find(&latestTimeStamp)

	// 数据库中没有更早的视频，就直接使用当前的时间戳替换
	if len(latestTimeStamp) == 0 {
		latestTimeStamp["t"] = time.Now().Unix()
	}

	// 将查询出来的最多30个Video与User进行连接，拼接出作者的名称
	result := []map[string]interface{}{}
	global.DB.Table("(?) as v", videoQuery).Select("v.id, v.author_id, v.file_name, v.publish_time as t, v.title, users.name").Joins("left join users on v.author_id = users.id").Scan(&result)

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
				Id:   authorId,
				Name: result[idx]["name"].(string),
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
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息：" + err.Error())
	}
}
