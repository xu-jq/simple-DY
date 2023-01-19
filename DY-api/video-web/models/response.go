/*
 * @Date: 2023-01-19 14:13:13
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-19 14:33:18
 * @FilePath: /simple-DY/DY-api/video-web/models/response.go
 * @Description: 后端发送响应的结构体
 */

package models

// 响应的状态码及描述
type ResponseCodeAndMessage struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// 1.1 视频流接口 响应参数​
type FeedResponse struct {
	Res       ResponseCodeAndMessage
	NextTime  int64   `json:"next_time"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []Video `json:"video_list"` // 视频列表
}

// 1.2.1 视频发布列表 响应参数
type PublishListResponse struct {
	Res       ResponseCodeAndMessage
	VideoList []Video `json:"video_list"` // 视频列表
}

// 1.2.2 投稿接口 响应参数
type PublishActionResponse struct {
	Res ResponseCodeAndMessage
}

// 1.3.1 用户信息 响应参数
type UserResponse struct {
	Res  ResponseCodeAndMessage
	User User `json:"user"` // 用户信息​
}

// 1.3.2 用户注册 1.3.3 用户登录  响应参数
type UserRegisterLoginResponse struct {
	Res    ResponseCodeAndMessage
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}
