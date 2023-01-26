/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 11:10:30
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/userinfo.go
 * @Description: UserInfo服务
 */
package handler

import (
	"context"
	"net"
	"simple-DY/DY-srvs/video-srv/global"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/dao"
	"simple-DY/DY-srvs/video-srv/utils/jwt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type userinfoserver struct {
	pb.UnimplementedUserInfoServer
}

func (s *userinfoserver) UserInfo(ctx context.Context, in *pb.DouyinUserRequest) (*pb.DouyinUserResponse, error) {

	// 构建返回的响应
	userResponse := pb.DouyinUserResponse{}

	// 将token解析并与id作比较
	statuscode := jwt.GetAndJudgeIdByToken(in.Token, in.UserId)

	userResponse.StatusCode = statuscode

	if statuscode != 0 {
		if statuscode == 4 {
			userResponse.StatusMsg = "没有携带Token信息！"
		} else if statuscode == 5 {
			userResponse.StatusMsg = "Token不正确！"
		}
		return &userResponse, nil
	}

	// Todo：并行处理三个查询

	// 通过id获取Users表的信息
	user := dao.GetUserById(in.UserId)
	// 查询关注总数
	followcount := dao.CountFollow(in.UserId)
	// 查询粉丝总数
	followercount := dao.CountFollower(in.UserId)

	// 返回响应
	userResponse.StatusMsg = "成功获取用户信息"
	userResponse.User = &pb.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   followcount,
		FollowerCount: followercount,
	}
	zap.L().Info("成功获取用户信息！")

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
	zap.L().Info("服务器监听地址：" + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		zap.L().Error("无法提供服务！错误信息：" + err.Error())
	}
}
