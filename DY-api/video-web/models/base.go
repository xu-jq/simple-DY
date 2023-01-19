/*
 * @Date: 2023-01-19 14:28:00
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-19 14:28:14
 * @FilePath: /simple-DY/DY-api/video-web/models/base.go
 * @Description: 信息类结构体
 */

package models

// 视频信息
type Video struct {
	Id            int64  `json:"id"`             // 视频唯一标识
	Author        User   `json:"author"`         // 视频作者信息​
	PlayUrl       string `json:"play_url"`       // 视频播放地址​
	CoverUrl      string `json:"cover_url"`      // 视频封面地址​
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string `json:"title"`          // 视频标题
}

// 用户信息
type User struct {
	Id            int64  `json:"id"`             // 用户id​
	Name          string `json:"name"`           // 用户名称​
	FollowCount   int64  `json:"follow_count"`   // 关注总数​
	FollowerCount int64  `json:"follower_count"` // 粉丝总数​
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注​
}
