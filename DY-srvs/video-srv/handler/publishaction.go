/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 18:04:25
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/publishaction.go
 * @Description: PublishAction服务
 */
package handler

import (
	"context"
	"net"
	"os"
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/models"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/ffmpeg"
	"simple-DY/DY-srvs/video-srv/utils/jwt"
	"simple-DY/DY-srvs/video-srv/utils/oss"
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

	// 将视频文件和图片文件存储在本地进行备份

	// 判断用户的文件夹路径是否存在并创建

	// 用户文件夹路径
	userPath := global.GlobalConfig.StaticBackup.StaticPath + tokenId.Id
	// 用户视频与图片的路径
	videoStaticPath := userPath + global.GlobalConfig.StaticBackup.VideoPath
	imageStaticPath := userPath + global.GlobalConfig.StaticBackup.ImagePath

	_, err := os.Stat(userPath)
	if os.IsNotExist(err) {
		videoerr := os.MkdirAll(videoStaticPath, 0666)
		imageerr := os.Mkdir(imageStaticPath, 0666)
		if videoerr != nil || imageerr != nil {
			zap.L().Error("创建文件夹失败！错误信息：" + videoerr.Error())
		}
	} else if err != nil {
		zap.L().Error("判断文件夹失败！错误信息：" + err.Error())
	}

	// 生成文件名称
	fileName := uuid.NewV4().String()

	// 组装完整的文件名称
	videoStaticFileName := videoStaticPath + fileName + global.GlobalConfig.StaticBackup.VideoSuffix
	imageStaticFileName := imageStaticPath + fileName + global.GlobalConfig.StaticBackup.ImageSuffix

	// 将字节流写入视频文件
	err = os.WriteFile(videoStaticFileName, []byte(in.Data), 0666)
	if err != nil {
		zap.L().Error("无法写入视频文件！错误信息：" + err.Error())
		return &publishActionResponse, nil
	}
	zap.L().Info("视频文件备份成功！路径：" + videoStaticFileName)

	// 截取视频文件的第一帧作为封面并存储
	err = ffmpeg.ExtractFirstFrame(videoStaticFileName, imageStaticFileName)
	if err != nil {
		zap.L().Error("无法写入图片文件！错误信息：" + err.Error())
		return &publishActionResponse, nil
	}
	zap.L().Info("图片文件备份成功！路径：" + imageStaticFileName)

	videoOSSFileName := tokenId.Id + global.GlobalConfig.OSS.VideoPath + fileName + global.GlobalConfig.OSS.VideoSuffix
	ImageOSSFileName := tokenId.Id + global.GlobalConfig.OSS.ImagePath + fileName + global.GlobalConfig.OSS.ImageSuffix

	// 上传视频文件
	err = oss.UploadFileToQiniuOSS(videoStaticFileName, videoOSSFileName)
	if err != nil {
		zap.L().Error("无法上传视频文件！错误信息：" + err.Error())
		return &publishActionResponse, nil
	}
	zap.L().Info("视频文件上传成功！路径：" + videoOSSFileName)

	// 上传图片文件
	err = oss.UploadFileToQiniuOSS(imageStaticFileName, ImageOSSFileName)
	if err != nil {
		zap.L().Error("无法上传图片文件！错误信息：" + err.Error())
		return &publishActionResponse, nil
	}
	zap.L().Info("图片文件上传成功！路径：" + ImageOSSFileName)

	// 17M文件，13秒发送给请求给服务器，64秒处理完返回响应
	// 31M文件，25秒发送给请求给服务器，121秒处理完返回响应
	// 客户端在发送请求后开始计时，10秒钟内不能返回响应就报网络错误

	authorId, _ := strconv.ParseInt(tokenId.Id, 10, 64)

	// 向数据库中插入数据
	videoInfo := models.Videos{
		AuthorId:    authorId,
		FileName:    fileName,
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
		zap.L().Error("无法监听客户端！错误信息：" + err.Error())
	}
	s := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*global.GlobalConfig.GRPC.GRPCMsgSize.LargeMB), grpc.MaxSendMsgSize(1024*1024*global.GlobalConfig.GRPC.GRPCMsgSize.LargeMB))
	pb.RegisterPublishActionServer(s, &publishactionserver{})
	zap.L().Info("服务器监听地址：" + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息：" + err.Error())
	}
}
