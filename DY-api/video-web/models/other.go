/*
 * @Date: 2023-01-19 14:13:27
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-19 14:29:57
 * @FilePath: /simple-DY/DY-api/video-web/models/other.go
 * @Description: 提前写的，但是不是本人负责部分的结构体
 */

package models

// // 视频评论
// type Comment struct {
// 	Id         int64  `json:"id"`          // 视频评论id​
// 	User       User   `json:"author"`      // 评论用户信息​
// 	Content    string `json:"content"`     // 评论内容
// 	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
// }

// // 消息
// type Message struct {
// 	Id         int64  `json:"id"`          // 消息id​
// 	Content    string `json:"content"`     // 消息内容
// 	CreateTime string `json:"create_time"` // 消息创建时间
// }

// // 2.1.1 喜欢列表 请求参数

// type douyin_favorite_list_request struct {
// 	UserId int64  `json:"user_id"` // 用户id
// 	Token  string `json:"token"`   // 用户鉴权token
// }

// // 2.1.1 喜欢列表 响应参数

// type douyin_favorite_list_response struct {
// 	Res       ResponseCodeAndMessage
// 	VideoList []Video `json:"video_list"` // 用户点赞视频列表
// }

// // 2.1.2 赞操作 请求参数

// type douyin_favorite_action_request struct {
// 	Token      string `json:"token"`       // 用户鉴权token
// 	VideoId    int64  `json:"video_id"`    // 视频id
// 	ActionType int32  `json:"action_type"` // 1-点赞，2-取消点赞
// }

// // 2.1.2 赞操作 响应参数

// type douyin_favorite_action_response struct {
// 	Res ResponseCodeAndMessage
// }

// // 2.2.1 视频评论列表 请求参数

// type douyin_comment_list_request struct {
// 	Token   string `json:"token"`    // 用户鉴权token
// 	VideoId int64  `json:"video_id"` // 视频id
// }

// // 2.2.1 视频评论列表 响应参数

// type douyin_comment_list_response struct {
// 	Res         ResponseCodeAndMessage
// 	CommentList []Comment `json:"comment_list"` // 评论列表
// }

// // 2.2.2 评论操作 请求参数

// type douyin_comment_action_request struct {
// 	Token       string `json:"token"`        // 用户鉴权token
// 	VideoId     int64  `json:"video_id"`     // 视频id
// 	ActionType  int32  `json:"action_type"`  // 1-发布评论，2-删除评论
// 	CommentText string `json:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
// 	CommentId   int64  `json:"comment_id"`   // 要删除的评论id，在action_type=2的时候使用
// }

// // 2.2.2 评论操作 响应参数

// type douyin_comment_action_response struct {
// 	Res     ResponseCodeAndMessage
// 	Comment Comment `json:"comment"` // 评论成功返回评论内容，不需要重新拉取整个列表
// }

// // 3.1.1 关注操作 请求参数

// type douyin_relation_action_request struct {
// 	Token      string `json:"token"`       // 用户鉴权token
// 	ToUserId   string `json:"to_user_id"`  // 对方用户id
// 	ActionType int32  `json:"action_type"` // 1-关注，2-取消关注
// }

// // 3.1.1 关注操作 响应参数

// type douyin_relation_action_response struct {
// 	Res ResponseCodeAndMessage
// }

// // 3.1.2 用户关注列表 请求参数

// type douyin_relation_follow_list_request struct {
// 	UserId int64  `json:"user_id"` // 用户id
// 	Token  string `json:"token"`   // 用户鉴权token
// }

// // 3.1.2 用户关注列表 响应参数

// type douyin_relation_follow_list_response struct {
// 	Res      ResponseCodeAndMessage
// 	UserList []User `json:"user_list"` // 用户信息列表
// }

// // 3.1.3 用户粉丝列表 请求参数

// type douyin_relation_follower_list_request struct {
// 	UserId int64  `json:"user_id"` // 用户id
// 	Token  string `json:"token"`   // 用户鉴权token
// }

// // 3.1.3 用户粉丝列表 响应参数

// type douyin_relation_follower_list_response struct {
// 	Res      ResponseCodeAndMessage
// 	UserList []User `json:"user_list"` // 用户列表
// }

// // 3.1.4 好友列表 请求参数

// type douyin_relation_friend_list_request struct {
// 	UserId int64  `json:"user_id"` // 用户id
// 	Token  string `json:"token"`   // 用户鉴权token
// }

// // 3.1.4 好友列表 响应参数

// type douyin_relation_friend_list_response struct {
// 	Res      ResponseCodeAndMessage
// 	UserList []User `json:"user_list"` // 用户列表
// }

// // 3.2.1 聊天记录 请求参数

// type douyin_message_chat_request struct {
// 	Token    string `json:"token"`      // 用户鉴权token
// 	ToUserId string `json:"to_user_id"` // 对方用户id
// }

// // 3.2.1 聊天记录 响应参数

// type douyin_message_chat_response struct {
// 	Res         ResponseCodeAndMessage
// 	MessageList []Message `json:"message_list"` // 消息列表
// }

// // 3.2.2 发送消息 请求参数

// type douyin_message_action_request struct {
// 	Token      string `json:"token"`       // 用户鉴权token
// 	ToUserId   string `json:"to_user_id"`  // 对方用户id
// 	ActionType int32  `json:"action_type"` // 1-发送消息
// 	Content    string `json:"content"`     // 消息内容
// }

// // 3.2.2 发送消息 响应参数

// type douyin_message_action_response struct {
// 	Res ResponseCodeAndMessage
// }
