/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-23 15:37:48
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/publishaction.go
 * @Description: PublishAction服务
 */
package handler

import (
	"context"
	"log"
	"net"
	"os"
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/models"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/ffmpeg"
	"simple-DY/DY-srvs/video-srv/utils/jwt"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type publishactionserver struct {
	pb.UnimplementedPublishActionServer
}

func (s *publishactionserver) PublishAction(ctx context.Context, in *pb.DouyinPublishActionRequest) (*pb.DouyinPublishActionResponse, error) {

	// 构建返回的响应
	publishActionResponse := pb.DouyinPublishActionResponse{}

	// 没有携带Token信息
	if len(in.Token) == 0 {
		publishActionResponse.StatusCode = -1
		publishActionResponse.StatusMsg = "没有携带Token信息！"
		zap.L().Error("没有携带Token信息！无法上传视频！")
		return &publishActionResponse, nil
	}

	// 从Token中读取携带的id信息
	tokenId, _ := jwt.ParseToken(strings.Fields(in.Token)[1])

	// 数据库查询和更新的模板
	user := models.Users{}

	// 根据姓名查找数据库中的用户信息
	global.DB.Select("name").Where("id = ?", tokenId.Id).Find(&user)

	// 如果这个用户不存在，则不能返回信息
	if user.Name == "" {
		publishActionResponse.StatusCode = 2
		publishActionResponse.StatusMsg = "用户不存在！"
		zap.L().Error("用户不存在！无法上传视频！")
		return &publishActionResponse, nil
	}

	// 判断文件夹路径是否存在并创建
	videoPath := global.GlobalConfig.StaticPath + global.GlobalConfig.VideoPath + tokenId.Id + "/"
	imagePath := global.GlobalConfig.StaticPath + global.GlobalConfig.ImagePath + tokenId.Id + "/"
	_, err := os.Stat(videoPath)
	if os.IsNotExist(err) {
		videoerr := os.Mkdir(videoPath, 0666)
		imageerr := os.Mkdir(imagePath, 0666)
		if videoerr != nil || imageerr != nil {
			zap.L().Error("创建文件夹失败！错误信息为：" + videoerr.Error())
		}
	} else if err != nil {
		zap.L().Error("判断文件夹失败！错误信息为：" + err.Error())
	}

	// 生成唯一的文件名称
	videoName := uuid.NewV4().String()
	videoPath = videoPath + videoName + ".mp4"
	imagePath = imagePath + videoName + ".jpg"

	// 将字节流写入视频文件
	err = os.WriteFile(videoPath, []byte(in.Data), 0666)
	if err != nil {
		zap.L().Error("无法写入文件！错误信息为：" + err.Error())
		return &publishActionResponse, nil
	}

	// 截取视频文件的第一帧作为封面
	ffmpeg.ExtractFirstFrame(videoPath, imagePath)

	authorId, _ := strconv.ParseInt(tokenId.Id, 10, 64)

	// 向数据库中插入数据
	videoInfo := models.Videos{
		AuthorId:    authorId,
		FileName:    videoName,
		VideoSuffix: ".mp4",
		PublishTime: time.Now().Unix(),
		Title:       in.Title,
	}
	global.DB.Create(&videoInfo)

	publishActionResponse = pb.DouyinPublishActionResponse{
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
