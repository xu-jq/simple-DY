/**
* @Author Wang Hui
* @Description
* @Date
**/
package handler

import (
	"context"
	"go.uber.org/zap"
	"simple-DY/DY-srvs/social-srv/global"
	"simple-DY/DY-srvs/social-srv/model"
	"simple-DY/DY-srvs/social-srv/proto"
)

type SocialServer struct {
	*proto.UnimplementedSocialServiceServer
}

func (s *SocialServer) GetFollowList(c context.Context, req *proto.GetFollowListRequest) (*proto.GetFollowListResponse, error) {
	zap.S().Info("GetFollowList Run")
	var follows []model.Follows
	result := global.DB.Where(&model.Follows{ID: req.UserId}).Find(&follows)
	if result.Error != nil {
		return nil, result.Error
	}
	zap.S().Info(result)
	resp := &proto.GetFollowListResponse{}
	for _, v := range follows {
		resp.UserList = append(resp.GetUserList(), &proto.User{Id: v.FollowerID})
	}
	return resp, nil
}
func (s *SocialServer) GetFollowerList(c context.Context, req *proto.FollowerListRequest) (*proto.FollowerListResponse, error) {
	return nil, nil
}
func (s *SocialServer) GetFriendList(c context.Context, req *proto.GetFriendListRequest) (*proto.GetFriendListResponse, error) {
	return nil, nil
}
func (s *SocialServer) RelationAction(c context.Context, req *proto.RelationActionRequest) (*proto.RelationActionResponse, error) {
	return nil, nil
}
func (s *SocialServer) MsgChat(c context.Context, req *proto.MsgChatRequest) (*proto.MsgChatResponse, error) {
	return nil, nil
}
func (s *SocialServer) MsgAction(c context.Context, req *proto.MsgActionRequest) (*proto.MsgActionResponse, error) {
	return nil, nil
}
