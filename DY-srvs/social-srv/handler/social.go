/**
* @Author Wang Hui
* @Description
* @Date
**/
package handler

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"simple-DY/DY-srvs/social-srv/global"
	"simple-DY/DY-srvs/social-srv/model"
	"simple-DY/DY-srvs/social-srv/proto"
	"time"
)

type SocialServer struct {
	*proto.UnimplementedSocialServiceServer
}

// GetFollowList 获取用户关注的所有用户列表。
func (s *SocialServer) GetFollowList(c context.Context, req *proto.GetFollowListRequest) (*proto.GetFollowListResponse, error) {
	zap.S().Info("GetFollowList Running")
	var follows []model.FollowsAndUser
	result := global.DB.Raw("SELECT users.name,follows.user_id,follows.follower_id "+
		"FROM users LEFT JOIN follows ON follows.user_id = users.id where users.id = ?", req.UserId).
		Find(&follows)
	if result.Error != nil {
		zap.S().Error("GetFollowList出错：", result.Error)
		return nil, result.Error
	}
	zap.S().Info(result)
	resp := &proto.GetFollowListResponse{}
	for _, v := range follows {
		resp.UserList = append(resp.GetUserList(), &proto.User{
			Id:            v.ID,
			Name:          v.Name,
			FollowCount:   int64(len(follows)),
			FollowerCount: 10,
			IsFollow:      true,
		})
	}
	return resp, nil
}

// GetFollowerList 用户粉丝列表
func (s *SocialServer) GetFollowerList(c context.Context, req *proto.FollowerListRequest) (*proto.FollowerListResponse, error) {
	zap.S().Info("GetFollowList Running")
	var follows []model.FollowsAndUser
	result := global.DB.Raw("SELECT users.name,follows.user_id,follows.follower_id "+
		"FROM follows LEFT JOIN users ON follows.follower_id = users.id where user_id = ?", req.UserId).
		Find(&follows)
	if result.Error != nil {
		zap.S().Error("GetFollowList出错：", result.Error)
		return nil, result.Error
	}
	zap.S().Info(result)
	resp := &proto.FollowerListResponse{}
	for _, v := range follows {
		resp.UserList = append(resp.UserList, &proto.User{
			Id:   v.ID,
			Name: v.Name,
		})
	}
	return resp, nil
}

// GetFriendList 用户好友列表
func (s *SocialServer) GetFriendList(c context.Context, req *proto.GetFriendListRequest) (*proto.GetFriendListResponse, error) {
	zap.S().Info("GetFollowList Running")
	var follows []model.FollowsAndUser
	result := global.DB.Raw("SELECT users.name,follows.user_id,follows.follower_id "+
		"FROM follows LEFT JOIN users ON follows.follower_id = users.id where user_id = ?", req.UserId).
		Find(&follows)
	if result.Error != nil {
		zap.S().Error("GetFollowList出错：", result.Error)
		return nil, result.Error
	}
	zap.S().Info(result)
	resp := &proto.GetFriendListResponse{}
	for _, v := range follows {
		resp.UserList = append(resp.UserList, &proto.User{
			Id:   v.ID,
			Name: v.Name,
		})
	}
	return resp, nil
}

// RelationAction 取关和关注
func (s *SocialServer) RelationAction(c context.Context, req *proto.RelationActionRequest) (*proto.RelationActionResponse, error) {
	zap.S().Info("RelationAction Running")
	var result *gorm.DB
	// 关注操作
	if req.ActionType == 1 {
		// 是不是以及存在此条记录
		follows := &model.Follows{}
		global.DB.Where(model.Follows{
			UserID:     req.ToUserId,
			FollowerID: req.UserId,
		}).First(&follows)
		if follows != nil {
			return nil, nil
		}
		// 插入记录
		result = global.DB.Save(&model.Follows{
			UserID:     req.ToUserId,
			FollowerID: req.UserId,
		})
	}
	// 取消关注操作
	if req.ActionType == 2 {
		result = global.DB.Delete(&model.Follows{
			UserID:     req.ToUserId,
			FollowerID: req.UserId,
		})
	}
	if result.Error != nil {
		zap.S().Error("RelationAction出错：", result.Error)
	}
	return &proto.RelationActionResponse{}, nil
}

func (s *SocialServer) MsgChat(c context.Context, req *proto.MsgChatRequest) (*proto.MsgChatResponse, error) {
	var mesList []model.Message
	result := global.DB.Where(&model.Message{
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
	}).Find(&mesList)
	if result.Error != nil {
		zap.S().Error("MsgChat出错：", result.Error)
		return nil, result.Error
	}
	var resp []*proto.Msg
	for _, v := range mesList {
		resp = append(resp, &proto.Msg{
			Id:         v.ID,
			Content:    v.Content,
			CreateTime: v.SentTime.String(),
		})
	}
	return &proto.MsgChatResponse{MessageList: resp}, nil
}

func (s *SocialServer) MsgAction(c context.Context, req *proto.MsgActionRequest) (*proto.MsgActionResponse, error) {
	if req.ActionType == 1 {
		result := global.DB.Save(&model.Message{
			UserID:   req.UserId,
			ToUserID: req.ToUserId,
			SentTime: time.Now(),
			Content:  req.Content,
		})
		if result.Error != nil {
			zap.S().Error("MsgAction：", result.Error)
			return &proto.MsgActionResponse{}, nil
		}
	}
	return &proto.MsgActionResponse{}, nil
}
