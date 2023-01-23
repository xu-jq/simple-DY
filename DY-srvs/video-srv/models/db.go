/*
 * @Date: 2023-01-20 19:46:14
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-23 12:07:29
 * @FilePath: /simple-DY/DY-srvs/video-srv/models/db.go
 * @Description: 数据库结构体
 */
package models

// Users表
type Users struct {
	Id       int64
	Name     string
	Password string
}

type Videos struct {
	Id          int64
	AuthorId    int64
	FileName    string
	VideoSuffix string
	PublishTime int64
	Title       string
}

type VideosAndAuthor struct {
	Id          int64
	AuthorId    int64
	FileName    string
	VideoSuffix string
	PublishTime *LocalTime
	Title       string
	Name        string
	Password    string
}
