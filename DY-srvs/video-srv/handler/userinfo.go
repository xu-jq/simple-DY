/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 23:18:18
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/userinfo.go
 * @Description: UserInfo服务
 */
package handler

import (
	"context"
	"log"
	"net"
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/models"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/jwt"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type userinfoserver struct {
	pb.UnimplementedUserInfoServer
}

func (s *userinfoserver) UserInfo(ctx context.Context, in *pb.DouyinUserRequest) (*pb.DouyinUserResponse, error) {

	// 构建返回的响应
	userResponse := pb.DouyinUserResponse{}

	// 没有携带Token信息
	if len(in.Token) == 0 {
		userResponse.StatusCode = -1
		userResponse.StatusMsg = "没有携带Token信息！"
		zap.L().Error("没有携带Token信息！无法获取用户信息！")
		return &userResponse, nil
	}

	// 从Token中读取携带的id信息
	tokenId, err := jwt.ParseToken(strings.Fields(in.Token)[1])
	if err != nil || tokenId.Id != strconv.FormatInt(in.UserId, 10) {
		userResponse.StatusCode = 1
		userResponse.StatusMsg = "Token不正确！"
		zap.L().Error("Token不正确！无法获取用户信息！")
		return &userResponse, nil
	}

	// 数据库查询和更新的模板
	user := models.Users{}

	// 根据姓名查找数据库中的用户信息
	global.DB.Where("id = ?", in.UserId).Find(&user)

	// 如果这个用户不存在，则不能返回信息
	if user.Id == 0 {
		userResponse.StatusCode = 2
		userResponse.StatusMsg = "用户不存在！"
		zap.L().Error("用户不存在！无法获取用户信息！")
	} else {
		userResponse.StatusCode = 0
		userResponse.StatusMsg = "成功获取用户信息"
		userResponse.User = &pb.User{
			Id:   user.Id,
			Name: user.Name,
		}
		zap.L().Info("成功获取用户信息！")
	}
	zap.L().Info("返回响应成功！")

	return &userResponse, nil
}

func UserInfoService(port string) {
	defer global.Wg.Done()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		zap.L().Error("无法监听客户端！错误信息：" + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterUserInfoServer(s, &userinfoserver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息：" + err.Error())
	}
}
