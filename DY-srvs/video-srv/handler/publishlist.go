/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-05 19:23:58
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/publishlist.go
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

func (s *Videoserver) PublishList(ctx context.Context, in *pb.DouyinPublishListRequest) (*pb.DouyinPublishListResponse, error) {

	// 构建返回的响应
	publishListResponse := pb.DouyinPublishListResponse{}

	// 根据id查找数据库中的用户信息
	user := dao.GetUserById(in.UserId)

	// 如果这个用户不存在，则不能返回信息
	if user.Name == "" {
		publishListResponse.StatusCode = 2
		publishListResponse.StatusMsg = "用户不存在！"
		zap.L().Error("用户不存在！无法获取用户投稿的视频！")
		return &publishListResponse, nil
	}

	// 查询作者视频
	result := dao.GetAuthorVideos(in.UserId)
	zap.L().Info("作者投稿视频查询完成！")

	videolistLen := len(result)

	publishListResponse = pb.DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  "作者投稿视频查询成功",
		VideoList:  make([]*pb.Video, videolistLen),
	}

	urlprefix := global.GlobalConfig.OSS.Address + strconv.FormatInt(in.UserId, 10)

	for idx := 0; idx < videolistLen; idx += 1 {
		filename := result[idx]["file_name"].(string)
		publishListResponse.VideoList[idx] = &pb.Video{
			Id: result[idx]["id"].(int64),
			Author: &pb.User{
				Id:   in.UserId,
				Name: user.Name,
			},
			PlayUrl:  urlprefix + global.GlobalConfig.OSS.VideoPath + filename + global.GlobalConfig.OSS.VideoSuffix,
			CoverUrl: urlprefix + global.GlobalConfig.OSS.ImagePath + filename + global.GlobalConfig.OSS.ImageSuffix,
			Title:    result[idx]["title"].(string),
		}
	}

	return &publishListResponse, nil
}
