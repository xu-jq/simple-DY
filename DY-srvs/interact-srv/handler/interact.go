package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"simple-DY/DY-srvs/interact-srv/global"
	"simple-DY/DY-srvs/interact-srv/model"
	"simple-DY/DY-srvs/interact-srv/proto"
	myJwt "simple-DY/DY-srvs/interact-srv/utils/jwt"
	"simple-DY/DY-srvs/interact-srv/utils/key"
	"strconv"
	"time"
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
	uId, _ := strconv.Atoi(userId)
	myKey := key.KeyUserFavorite(userId)             // 点赞key
	isMember := global.RDB.SIsMember(myKey, videoId) // 判断是否点赞
	if actionType == 1 && isMember.Val() {           // 点赞操作
		global.RDB.SAdd(myKey, videoId)
		//将点赞信息保存到数据库
		like := model.Like{
			UserId:  int64(uId),
			VideoId: videoId,
		}
		if err := global.DB.Create(&like).Error; err != nil {
			return nil, err
		}
		return &proto.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		}, nil
	} else if actionType == 0 && !isMember.Val() { //取消点赞操作
		global.RDB.SRem(myKey, videoId)
		//将取消点赞信息保存到数据库
		if res := global.DB.Where("user_id = ?, video_id = ?", uId, videoId).Delete(&model.Like{}); res.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "未点赞")
		}
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
	if videoIds.Val() == nil {
		return &proto.DouyinFavoriteListResponse{
			StatusCode: 0,
			StatusMsg:  "列表为空",
			VideoList:  nil,
		}, nil
	}
	// ToDo 调用服务，根据videoIds封装resp
	var videos []*proto.Video
	for _, videoId := range videoIds.Val() {
		vid, _ := strconv.Atoi(videoId)
		videoInfo, _ := global.VideoSrvClient.VideoInfo(context.Background(), &proto.DouyinVideoInfoRequest{VideoId: int64(vid)})
		author := videoInfo.VideoList.Author
		video := &proto.Video{
			Id: int64(vid),
			Author: &proto.User{
				Id:            author.Id,
				Name:          author.Name,
				FollowCount:   author.FollowerCount,
				FollowerCount: author.FollowerCount,
				IsFollow:      author.IsFollow,
			},
			PlayUrl:       videoInfo.VideoList.PlayUrl,
			CoverUrl:      videoInfo.VideoList.CoverUrl,
			FavoriteCount: videoInfo.VideoList.FavoriteCount,
			CommentCount:  videoInfo.VideoList.CommentCount,
			IsFavorite:    videoInfo.VideoList.IsFavorite,
			Title:         videoInfo.VideoList.Title,
		}
		videos = append(videos, video)
	}

	resp := proto.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "成功",
		VideoList:  videos,
	}
	return &resp, nil
}

func (*InteractServer) CommentAction(ctx context.Context, req *proto.DouyinCommentActionRequest) (*proto.DouyinCommentActionResponse, error) {
	claims, _ := myJwt.ParseToken(req.Token)
	userId := claims.Id
	actionType := req.ActionType
	uId, _ := strconv.Atoi(userId)
	if actionType == 1 { //发表评论
		comment := model.Comment{
			UserId:      int64(uId),
			VideoId:     req.VideoId,
			CommentText: req.CommentText,
			CreateTime:  time.Time{},
		}
		if err := global.DB.Create(&comment).Error; err != nil {
			return nil, err
		}
		resp := proto.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "评论成功",
			//ToDo 调用服务根据comment_id,封装comment结构体
			Comment: nil,
		}
		return &resp, nil
	} else if actionType == 2 { //删除评论
		if res := global.DB.Where("id = ?", req.CommentId).Delete(&model.Comment{}); res.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "未找到该评论")
		}
		resp := proto.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "删除评论成功",
		}
		return &resp, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "参数无效")
}

func (*InteractServer) GetCommentList(ctx context.Context, req *proto.DouyinCommentListRequest) (*proto.DouyinCommentListResponse, error) {
	var comments []model.Comment
	if res := global.DB.Where("video_id=?", req.VideoId).Find(&comments); res.RowsAffected == 0 {
		resp := proto.DouyinCommentListResponse{
			StatusCode:  0,
			StatusMsg:   "获取评论成功",
			CommentList: nil,
		}
		return &resp, nil
	}
	// ToDo 根据comments调用服务封装resp

	resp := proto.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "获取评论成功",
		CommentList: nil,
	}
	return &resp, nil
}

func (*InteractServer) GetLikeVideoUserId(ctx context.Context, req *proto.DouyinLikeVideoRequest) (*proto.DouyinLikeVideoResponse, error) {
	var likes []model.Like
	if res := global.DB.Where("video_id=?", req.VideoId).Find(&likes); res.RowsAffected == 0 {
		return &proto.DouyinLikeVideoResponse{
			VideoId: req.VideoId,
			UserId:  nil,
		}, nil
	}
	var uIds []int64
	for _, like := range likes {
		uIds = append(uIds, like.UserId)
	}
	return &proto.DouyinLikeVideoResponse{
		VideoId: req.VideoId,
		UserId:  uIds,
	}, nil
}
