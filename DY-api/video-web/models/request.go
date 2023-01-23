/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-23 09:51:48
 * @FilePath: /simple-DY/DY-api/video-web/models/request.go
 * @Description: 后端接收请求的结构体
 */

package models

// 1.1 视频流接口 请求参数​
type FeedRequest struct {
	LatestTime int64  `json:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `json:"token"`       // 用户登录状态下设置
}

// 1.2.1 视频发布列表 请求参数
type PublishListRequest struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

// 1.2.2 投稿接口 请求参数

type PublishActionRequest struct {
	Token string `json:"token"` // 用户鉴权token
	Data  []byte `json:"data"`  // 视频数据
	Title string `json:"title"` // 视频标题
}

// 1.3.1 用户信息 请求参数

type UserRequest struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

// 1.3.2 用户注册 1.3.3 用户登录  请求参数

type UserRegisterLoginRequest struct {
	UserName string `json:"username"` // 注册用户名，最长32个字符
	Password string `json:"password"` // 密码，最长32个字符
}
