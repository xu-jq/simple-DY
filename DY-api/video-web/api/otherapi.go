/*
 * @Date: 2023-02-02 16:32:06
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-03 10:26:14
 * @FilePath: /simple-DY/DY-api/video-web/api/otherapi.go
 * @Description: 调用其他接口
 */
package api

import (
	"context"
	"simple-DY/DY-api/video-web/global"
	socialpb "simple-DY/DY-api/video-web/proto/social"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// 获取关注总数
func douyinGetFollowList(user_id string) (responseGetFollowList *socialpb.GetFollowListResponse, err error) {
	// 将接收的客户端请求参数绑定到结构体上
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		zap.L().Error("用户id转换为整数失败！错误信息：" + err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()
	responseGetFollowList, err = global.SocialServiceClient.GetFollowList(ctx, &socialpb.GetFollowListRequest{
		UserId: userId,
	})
	zap.L().Info("通过GRPC接收到的响应：" + responseGetFollowList.String())
	return
}

// 获取粉丝总数
func douyinFollowerList(user_id string) (responseFollowerList *socialpb.FollowerListResponse, err error) {
	// 将接收的客户端请求参数绑定到结构体上
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		zap.L().Error("用户id转换为整数失败！错误信息：" + err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()
	responseFollowerList, err = global.SocialServiceClient.GetFollowerList(ctx, &socialpb.FollowerListRequest{
		UserId: userId,
	})
	zap.L().Info("通过GRPC接收到的响应：" + responseFollowerList.String())
	return
}
