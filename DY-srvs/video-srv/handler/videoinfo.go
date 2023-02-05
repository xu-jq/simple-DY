/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-05 14:28:57
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/videoinfo.go
 * @Description: PublishAction服务
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

type VideoInfoserver struct {
	pb.UnimplementedVideoInfoServer
}

func (s *VideoInfoserver) VideoInfo(ctx context.Context, in *pb.DouyinVideoInfoRequest) (*pb.DouyinVideoInfoResponse, error) {

	// 构建返回的响应
	videoInfoResponse := pb.DouyinVideoInfoResponse{}

	// 根据id查找数据库中的用户信息
	video := dao.GetVideoById(in.VideoId)

	// 如果这个视频不存在，则不能返回信息
	if video.FileName == "" {
		videoInfoResponse.StatusCode = 1
		videoInfoResponse.StatusMsg = "视频不存在！"
		zap.L().Error("视频不存在！无法获取视频！")
		return &videoInfoResponse, nil
	}

	urlprefix := global.GlobalConfig.OSS.Address + strconv.FormatInt(video.AuthorId, 10)

	videoInfoResponse = pb.DouyinVideoInfoResponse{
		StatusCode: 0,
		StatusMsg:  "视频查询成功",
		VideoList: &pb.Video{
			Id: in.VideoId,
			Author: &pb.User{
				Id: video.AuthorId,
			},
			PlayUrl:  urlprefix + global.GlobalConfig.OSS.VideoPath + video.FileName + global.GlobalConfig.OSS.VideoSuffix,
			CoverUrl: urlprefix + global.GlobalConfig.OSS.ImagePath + video.FileName + global.GlobalConfig.OSS.ImageSuffix,
			Title:    video.Title,
		},
	}

	return &videoInfoResponse, nil
}
