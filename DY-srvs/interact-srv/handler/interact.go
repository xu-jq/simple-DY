package handler

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"simple-DY/DY-srvs/interact-srv/global"
	"simple-DY/DY-srvs/interact-srv/model"
	"simple-DY/DY-srvs/interact-srv/proto"
	"simple-DY/DY-srvs/interact-srv/utils/jwt"
	"simple-DY/DY-srvs/interact-srv/utils/key"
	"strconv"
	"strings"
	"time"
)

type InteractServer struct {
	proto.UnimplementedInteractServiceServer
}

func (*InteractServer) FavoriteAction(ctx context.Context, req *proto.DouyinFavoriteActionRequest) (*proto.DouyinFavoriteActionResponse, error) {
	// 解析token得到user.id
	tokenId, err := jwt.ParseToken(strings.Fields(req.Token)[1])
	if err != nil {
		zap.S().Error("token:", err.Error())
		return nil, err
	}
	userId, _ := strconv.ParseInt(tokenId.Id, 10, 64)
	videoId := req.VideoId
	actionType := req.ActionType
	myKey := key.KeyUserFavorite(userId)             // 点赞key
	isMember := global.RDB.SIsMember(myKey, videoId) // 判断是否点赞
	if actionType == 1 && !isMember.Val() {          // 点赞操作
		global.RDB.SAdd(myKey, videoId)
		//将点赞信息保存到数据库
		like := model.Likes{
			UserId:  userId,
			VideoId: videoId,
		}
		if err := global.DB.Create(&like).Error; err != nil {
			zap.S().Error("create like err:", err.Error())
			return nil, err
		}
		return &proto.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		}, nil
	} else if actionType == 2 && isMember.Val() { //取消点赞操作
		global.RDB.SRem(myKey, videoId)
		//将取消点赞信息保存到数据库
		global.DB.Where("user_id = ? and video_id = ?", userId, videoId).Delete(&model.Likes{})
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
	var videos []*proto.Video
	for _, videoId := range videoIds.Val() {
		vid, _ := strconv.Atoi(videoId)
		videoInfo, err := global.VideoSrvClient.VideoInfo(context.Background(), &proto.DouyinVideoInfoRequest{VideoId: int64(vid)})
		if err != nil {
			zap.S().Error("srv conn video.videoInfo err:", err.Error())
			return nil, err
		}
		author := videoInfo.VideoList.Author
		if author == nil {
			zap.S().Error("this author is nil")
			return nil, err
		}
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
	tokenId, err := jwt.ParseToken(strings.Fields(req.Token)[1])
	if err != nil {
		zap.S().Error("token:", err.Error())
		return nil, err
	}
	userId, _ := strconv.ParseInt(tokenId.Id, 10, 64)
	actionType := req.ActionType
	if actionType == 1 { //发表评论
		comment := model.Comments{
			UserId:      userId,
			VideoId:     req.VideoId,
			CommentText: req.CommentText,
			CreateTime:  time.Now(),
		}
		if err := global.DB.Create(&comment).Error; err != nil {
			zap.S().Error("create comment err:", err.Error())
			return nil, err
		}
		userInfo, err := global.VideoSrvClient.UserInfo(context.Background(), &proto.DouyinUserRequest{
			UserId: userId,
			Token:  req.Token,
		})
		if err != nil {
			zap.S().Error("conn videoSrv.userInfo err:", err.Error())
		}
		resp := proto.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "评论成功",
			Comment: &proto.Comment{
				Id: comment.Id,
				User: &proto.User{
					Id:            userInfo.User.Id,
					Name:          userInfo.User.Name,
					FollowCount:   userInfo.User.FollowCount,
					FollowerCount: userInfo.User.FollowerCount,
					IsFollow:      userInfo.User.IsFollow,
				},
				Content:    comment.CommentText,
				CreateDate: comment.CreateTime.Format("01-02"),
			},
		}
		return &resp, nil
	} else if actionType == 2 { //删除评论
		if res := global.DB.Where("id = ?", req.CommentId).Delete(&model.Comments{}); res.RowsAffected == 0 {
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
	var comments []*model.Comments
	res := global.DB.Where("video_id=?", req.VideoId).Find(&comments)
	zap.S().Info("sqlErr:", res.Error)
	if res.RowsAffected == 0 {
		resp := proto.DouyinCommentListResponse{
			StatusCode:  0,
			StatusMsg:   "暂无评论",
			CommentList: nil,
		}
		return &resp, nil
	}
	var commonLists []*proto.Comment
	for _, comment := range comments {
		userInfo, err := global.VideoSrvClient.UserInfo(ctx, &proto.DouyinUserRequest{
			UserId: comment.UserId,
			Token:  req.Token,
		})
		if err != nil {
			zap.S().Error("conn videoSrv.userInfo err:", err.Error())
		}
		if userInfo == nil {
			zap.S().Error("user is nil")
			return nil, err
		}
		commonList := &proto.Comment{
			Id: comment.VideoId,
			User: &proto.User{
				Id:            userInfo.User.Id,
				Name:          userInfo.User.Name,
				FollowCount:   userInfo.User.FollowCount,
				FollowerCount: userInfo.User.FollowerCount,
				IsFollow:      userInfo.User.IsFollow,
			},
			Content:    comment.CommentText,
			CreateDate: comment.CreateTime.Format("01-02"),
		}
		commonLists = append(commonLists, commonList)
	}
	resp := proto.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "获取评论成功",
		CommentList: commonLists,
	}
	return &resp, nil
}

func (*InteractServer) GetLikeVideoUserId(ctx context.Context, req *proto.DouyinLikeVideoRequest) (*proto.DouyinLikeVideoResponse, error) {
	var likes []*model.Likes
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
