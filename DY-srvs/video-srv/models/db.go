/*
 * @Date: 2023-01-20 19:46:14
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 10:48:33
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

// Videos表
type Videos struct {
	Id          int64
	AuthorId    int64
	FileName    string
	PublishTime int64
	Title       string
}

// Follows表
type Follows struct {
	Id         int64
	UserId     int64
	FollowerId int64
}
