/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 16:23:20
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/publishlist.go
 * @Description: PublishAction服务
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

type publishlistserver struct {
	pb.UnimplementedPublishListServer
}

func (s *publishlistserver) PublishList(ctx context.Context, in *pb.DouyinPublishListRequest) (*pb.DouyinPublishListResponse, error) {

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

func PublishListService(port string) {
	defer global.Wg.Done()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息：" + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterPublishListServer(s, &publishlistserver{})
	zap.L().Info("服务器监听地址：" + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息：" + err.Error())
	}
}
