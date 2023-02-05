package model

import "time"

type Comment struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	VideoId     int64     `json:"video_id"`
	CommentText string    `json:"comment_text"`
	CreateTime  time.Time `json:"create_time"`
}
