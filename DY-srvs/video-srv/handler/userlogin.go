/*
 * @Date: 2023-01-20 14:46:54
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-02 15:25:28
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/userlogin.go
 * @Description: UserLogin服务
 */
package handler

import (
	"context"
	pb "simple-DY/DY-srvs/video-srv/proto"
	"simple-DY/DY-srvs/video-srv/utils/dao"
	"simple-DY/DY-srvs/video-srv/utils/md5salt"

	"go.uber.org/zap"
)

type Userloginserver struct {
	pb.UnimplementedUserLoginServer
}

func (s *Userloginserver) UserLogin(ctx context.Context, in *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {

	// 构建返回的响应
	userLoginResponse := pb.DouyinUserLoginResponse{}

	// 通过name获取Users表的信息
	user := dao.GetUserByName(in.Username)

	// 如果这个用户不存在，则不能登录
	if user.Id == 0 {
		userLoginResponse.StatusCode = 2
		userLoginResponse.StatusMsg = "用户不存在！"
		zap.L().Error("用户不存在！无法登录！用户名称：" + in.Username)
		return &userLoginResponse, nil
	}

	// 将用户密码加密
	password := md5salt.MD5V(in.Password, in.Username, 1)

	// 验证密码是否正确
	if user.Password != password {
		userLoginResponse.StatusCode = 3
		userLoginResponse.StatusMsg = "密码错误！"
		zap.L().Error("密码错误！无法登录！")
		return &userLoginResponse, nil
	}

	// 构建返回的响应
	userLoginResponse.StatusCode = 0
	userLoginResponse.StatusMsg = "登录成功！"
	userLoginResponse.UserId = user.Id
	// userLoginResponse.Token = jwt.GenerateToken(user.Id)
	zap.L().Info("用户名和密码正确！登录成功！")

	return &userLoginResponse, nil
}
