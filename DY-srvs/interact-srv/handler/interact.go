package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"simple-DY/DY-srvs/interact-srv/global"
	"simple-DY/DY-srvs/interact-srv/proto"
	myJwt "simple-DY/DY-srvs/interact-srv/utils/jwt"
	"simple-DY/DY-srvs/interact-srv/utils/key"
)

type InteractServer struct {
	proto.UnimplementedInteractServiceServer
}

func (*InteractServer) FavoriteAction(ctx context.Context, req *proto.DouyinFavoriteActionRequest) (*proto.DouyinFavoriteActionResponse, error) {
	// 解析token得到user.id
	claims, _ := myJwt.ParseToken(req.Token)
	userId := claims.Id
	videoId := req.VideoId
	actionType := req.ActionType

	myKey := key.KeyUserFavorite(userId)             // 点赞key
	isMember := global.RDB.SIsMember(myKey, videoId) // 判断是否点赞
	if actionType == 1 && isMember.Val() {           // 点赞操作
		global.RDB.SAdd(myKey, videoId)
		//ToDO 将点赞信息保存到数据库

		return &proto.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		}, nil
	} else if actionType == 0 && !isMember.Val() { //取消点赞操作
		global.RDB.SRem(myKey, videoId)
		//ToDO 将取消点赞信息保存到数据库

		return &proto.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "取消点赞成功",
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "参数无效")
}

func (*InteractServer) GetFavoriteList(ctx context.Context, req *proto.DouyinFavoriteListRequest) (*proto.DouyinFavoriteListResponse, error) {
	userId := req.UserId
	myKey := key.KeyUserFavorite(userId)
	videoIds := global.RDB.SMembers(myKey)
	for _, videoId := range videoIds.Val() {
		//TODO 根据video_id查video及user将其封装进resp

	}
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}

func (*InteractServer) CommentAction(ctx context.Context, req *proto.DouyinCommentActionRequest) (*proto.DouyinCommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}

func (*InteractServer) GetCommentList(ctx context.Context, req *proto.DouyinCommentListRequest) (*proto.DouyinCommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}

func (*InteractServer) GetLikeVideoUserId(ctx context.Context, req *proto.DouyinLikeVideoRequest) (*proto.DouyinLikeVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
