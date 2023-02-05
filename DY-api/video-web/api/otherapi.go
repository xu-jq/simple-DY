/*
 * @Date: 2023-02-02 16:32:06
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-05 18:33:30
 * @FilePath: /simple-DY/DY-api/video-web/api/otherapi.go
 * @Description: 调用其他接口
 */
package api

import (
	"context"
	"simple-DY/DY-api/video-web/global"

	pb "simple-DY/DY-api/video-web/proto"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// 获取关注总数
func douyinGetFollowList(user_id string) (responseGetFollowList *pb.GetFollowListResponse, err error) {
	// 将接收的客户端请求参数绑定到结构体上
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		zap.L().Error("用户id转换为整数失败！错误信息：" + err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()
	responseGetFollowList, err = global.SocialServiceClient.GetFollowList(ctx, &pb.GetFollowListRequest{
		UserId: userId,
	})
	zap.L().Info("通过GRPC接收到的响应：" + responseGetFollowList.String())
	return
}

// 获取粉丝总数
func douyinFollowerList(user_id string) (responseFollowerList *pb.FollowerListResponse, err error) {
	// 将接收的客户端请求参数绑定到结构体上
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		zap.L().Error("用户id转换为整数失败！错误信息：" + err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()
	responseFollowerList, err = global.SocialServiceClient.GetFollowerList(ctx, &pb.FollowerListRequest{
		UserId: userId,
	})
	zap.L().Info("通过GRPC接收到的响应：" + responseFollowerList.String())
	return
}

// 获取视频点赞信息
func douyinLikeVideo(video_id string) (responseLikeVideo *pb.DouyinLikeVideoResponse, err error) {
	// 将接收的客户端请求参数绑定到结构体上
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		zap.L().Error("用户id转换为整数失败！错误信息：" + err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()
	responseLikeVideo, err = global.InteractServiceClient.GetLikeVideoUserId(ctx, &pb.DouyinLikeVideoRequest{
		VideoId: videoId,
	})
	zap.L().Info("通过GRPC接收到的响应：" + responseLikeVideo.String())
	return
}

// 获取视频评论信息
func douyinCommentList(video_id, token string) (responseCommentList *pb.DouyinCommentListResponse, err error) {
	// 将接收的客户端请求参数绑定到结构体上
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		zap.L().Error("用户id转换为整数失败！错误信息：" + err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.GlobalConfig.GRPC.GRPCTimeOut.CommonSecond))
	defer cancel()
	responseCommentList, err = global.InteractServiceClient.GetCommentList(ctx, &pb.DouyinCommentListRequest{
		VideoId: videoId,
		Token:   token,
	})
	zap.L().Info("通过GRPC接收到的响应：" + responseCommentList.String())
	return
}
