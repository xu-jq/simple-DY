/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 14:58:33
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/publishaction.go
 * @Description: PublishAction服务
 */
package handler

import (
	"context"
	"net"
	"os"
	"simple-DY/DY-srvs/video-srv/global"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/backup"
	"simple-DY/DY-srvs/video-srv/utils/dao"
	"simple-DY/DY-srvs/video-srv/utils/ffmpeg"
	"simple-DY/DY-srvs/video-srv/utils/jwt"
	"simple-DY/DY-srvs/video-srv/utils/oss"
	"strconv"
	"strings"
	"time"

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
		publishActionResponse.StatusCode = 4
		publishActionResponse.StatusMsg = "没有携带Token信息！"
		zap.L().Error("没有携带Token信息！无法上传视频！")
		return &publishActionResponse, nil
	}

	// 从Token中读取携带的id信息
	tokenId, _ := jwt.ParseToken(strings.Fields(in.Token)[1])

	// 将token内部的id转换为int类型
	authorId, _ := strconv.ParseInt(tokenId.Id, 10, 64)

	// 根据id查找数据库中的用户信息
	user := dao.GetUserById(authorId)

	// 如果这个用户不存在，则不能返回信息
	if user.Name == "" {
		zap.L().Error("用户不存在！无法上传视频！")
		publishActionResponse.StatusCode = 5
		publishActionResponse.StatusMsg = "用户不存在！"
		return &publishActionResponse, nil
	}

	// 判断用户的文件夹路径是否存在并创建
	fileName, videoStaticFileName, imageStaticFileName, err := backup.GenerateFilePath(tokenId.Id)
	if err != nil {
		zap.L().Error("备份文件夹操作失败！错误信息：" + err.Error())
		publishActionResponse.StatusCode = 6
		publishActionResponse.StatusMsg = "备份文件夹操作失败！"
		return &publishActionResponse, nil
	}
	zap.L().Info("备份文件夹操作成功！错误信息：")

	// 将字节流写入视频文件
	err = os.WriteFile(videoStaticFileName, []byte(in.Data), 0666)
	if err != nil {
		zap.L().Error("无法写入视频文件！错误信息：" + err.Error())
		publishActionResponse.StatusCode = 7
		publishActionResponse.StatusMsg = "无法写入视频文件！"
		return &publishActionResponse, nil
	}
	zap.L().Info("视频文件备份成功！路径：" + videoStaticFileName)

	// 截取视频文件的第一帧作为封面并存储
	err = ffmpeg.ExtractFirstFrame(videoStaticFileName, imageStaticFileName)
	if err != nil {
		zap.L().Error("无法写入图片文件！错误信息：" + err.Error())
		publishActionResponse.StatusCode = 8
		publishActionResponse.StatusMsg = "无法写入图片文件！"
		return &publishActionResponse, nil
	}
	zap.L().Info("图片文件备份成功！路径：" + imageStaticFileName)

	videoOSSFileName := tokenId.Id + global.GlobalConfig.OSS.VideoPath + fileName + global.GlobalConfig.OSS.VideoSuffix
	ImageOSSFileName := tokenId.Id + global.GlobalConfig.OSS.ImagePath + fileName + global.GlobalConfig.OSS.ImageSuffix

	// 上传视频文件
	err = oss.UploadFileToQiniuOSS(videoStaticFileName, videoOSSFileName)
	if err != nil {
		zap.L().Error("无法上传视频文件！错误信息：" + err.Error())
		publishActionResponse.StatusCode = 9
		publishActionResponse.StatusMsg = "无法上传视频文件！"
		return &publishActionResponse, nil
	}
	zap.L().Info("视频文件上传成功！路径：" + videoOSSFileName)

	// 上传图片文件
	err = oss.UploadFileToQiniuOSS(imageStaticFileName, ImageOSSFileName)
	if err != nil {
		zap.L().Error("无法上传图片文件！错误信息：" + err.Error())
		publishActionResponse.StatusCode = 10
		publishActionResponse.StatusMsg = "无法上传图片文件！"
		return &publishActionResponse, nil
	}
	zap.L().Info("图片文件上传成功！路径：" + ImageOSSFileName)

	// 17M文件，13秒发送给请求给服务器，64秒处理完返回响应
	// 31M文件，25秒发送给请求给服务器，121秒处理完返回响应
	// 客户端在发送请求后开始计时，10秒钟内不能返回响应就报网络错误

	// 向数据库中插入数据
	dao.InsertVideo(authorId, fileName, time.Now().Unix(), in.Title)

	publishActionResponse = pb.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "投稿视频上传成功",
	}

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
