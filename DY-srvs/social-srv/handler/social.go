/**
* @Author Wang Hui
* @Description
* @Date
**/
package handler

import (
	"context"
	"errors"
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
	result := global.DB.Raw("SELECT users.name,users.id,follows.user_id,follows.follower_id "+
		"FROM follows LEFT JOIN users ON follows.user_id = users.id where follows.follower_id = ?", req.UserId).
		Find(&follows)
	if result.Error != nil {
		zap.S().Error("GetFollowList出错：", result.Error)
		return nil, result.Error
	}
	zap.S().Info(result)
	resp := &proto.GetFollowListResponse{}
	for _, v := range follows {
		resp.UserList = append(resp.GetUserList(), &proto.User{
			Id:       v.Id,
			Name:     v.Name,
			IsFollow: true,
		})
	}
	return resp, nil
}

// GetFollowerList 用户粉丝列表
func (s *SocialServer) GetFollowerList(c context.Context, req *proto.FollowerListRequest) (*proto.FollowerListResponse, error) {
	zap.S().Info("GetFollowList Running")
	var follows []model.FollowsAndUser
	result := global.DB.Debug().Raw("SELECT users.name,users.id,follows.user_id,follows.follower_id "+
		"FROM follows LEFT JOIN users ON follows.follower_id = users.id where user_id = ?", req.UserId).
		Find(&follows)
	var follow []model.FollowsAndUser
	getFollow := global.DB.Debug().Raw("select follows.user_id FROM follows where follower_id = ?", req.UserId).
		Find(&follow)
	if result.Error != nil {
		zap.S().Error("GetFollowList出错：", result.Error)
		return nil, result.Error
	}
	if getFollow.Error != nil {
		zap.S().Error("GetFollowList出错：", getFollow.Error)
		return nil, result.Error
	}
	followMap := make(map[int64]struct{}, 0)
	for _, v := range follow {
		followMap[v.UserID] = struct{}{}
	}
	zap.S().Info(result)
	resp := &proto.FollowerListResponse{}
	for _, v := range follows {
		isFollow := false
		if _, ok := followMap[v.Id]; ok {
			isFollow = true
		}
		resp.UserList = append(resp.UserList, &proto.User{
			Id:       v.Id,
			Name:     v.Name,
			IsFollow: isFollow,
		})
	}
	return resp, nil
}

// GetFriendList 用户好友列表
func (s *SocialServer) GetFriendList(c context.Context, req *proto.GetFriendListRequest) (*proto.GetFriendListResponse, error) {
	zap.S().Info("GetFollowList Running")
	var follows []model.FollowsAndUser
	result := global.DB.Raw("SELECT users.name,users.id,follows.user_id,follows.follower_id "+
		"FROM follows LEFT JOIN users ON follows.follower_id = users.id where user_id = ?", req.UserId).
		Find(&follows)
	if result.Error != nil {
		zap.S().Error("GetFollowList出错：", result.Error)
		return nil, result.Error
	}
	zap.S().Info(result)
	resp := &proto.GetFriendListResponse{}
	for _, v := range follows {
		msg := &model.Messages{}
		global.DB.Raw("SELECT * FROM `messages` WHERE (user_id =? and to_user_id = ?) or (user_id =? and to_user_id = ?) "+
			"ORDER BY sent_time desc LIMIT 0,1", req.UserId, v.FollowerID, v.FollowerID, req.UserId).First(&msg)
		resMsg := ""
		if msg.Content != "" {
			resMsg = msg.Content
		}
		msgType := 1
		if msg.ToUserID == req.UserId {
			msgType = 0
		}
		resp.UserList = append(resp.UserList, &proto.FriendUser{
			Message: resMsg,
			MsgType: int64(msgType),
			Id:      v.Id,
			Name:    v.Name,
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
		result = global.DB.Where(&model.Follows{
			UserID:     req.ToUserId,
			FollowerID: req.UserId,
		}).First(&follows)
		if result.RowsAffected > 0 {
			return &proto.RelationActionResponse{}, errors.New("已关注")
		}
		// 插入记录
		result = global.DB.Create(&model.Follows{
			UserID:     req.ToUserId,
			FollowerID: req.UserId,
		})
	}
	// 取消关注操作
	if req.ActionType == 2 {
		result = global.DB.Where("user_id = ? and follower_id = ?", req.ToUserId, req.UserId).
			Delete(&model.Follows{})
	}
	if result != nil && result.Error != nil {
		zap.S().Error("RelationAction出错：", result.Error)
	}
	return &proto.RelationActionResponse{}, nil
}

func (s *SocialServer) MsgChat(c context.Context, req *proto.MsgChatRequest) (*proto.MsgChatResponse, error) {
	zap.S().Info("MsgChat Running")
	zap.S().Info("MsgChat 参数：", req)
	var mesList []model.Messages
	result := global.DB.Debug().Where("(user_id =? and to_user_id = ?) "+
		"or (user_id =? and to_user_id = ?)", req.UserId, req.ToUserId, req.ToUserId, req.UserId).Order("sent_time").
		Find(&mesList)
	if result.Error != nil {
		zap.S().Error("MsgChat出错：", result.Error)
		return nil, result.Error
	}
	var resp []*proto.Msg
	count := 0
	for _, v := range mesList {
		count++
		resp = append(resp, &proto.Msg{
			Id:         v.ID,
			ToUserId:   v.ToUserID,
			FromUserId: v.UserID,
			Content:    v.Content,
			CreateTime: int64(count),
		})
	}
	return &proto.MsgChatResponse{MessageList: resp}, nil
}

func (s *SocialServer) MsgAction(c context.Context, req *proto.MsgActionRequest) (*proto.MsgActionResponse, error) {
	if req.ActionType == 1 {
		result := global.DB.Create(&model.Messages{
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
