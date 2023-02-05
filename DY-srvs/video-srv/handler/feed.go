/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-05 19:22:35
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/feed.go
 * @Description: Feed服务
 */
package handler

import (
	"context"
	"simple-DY/DY-srvs/video-srv/global"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/dao"
	"strconv"

	"go.uber.org/zap"
)

func (s *Videoserver) Feed(ctx context.Context, in *pb.DouyinFeedRequest) (*pb.DouyinFeedResponse, error) {

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
