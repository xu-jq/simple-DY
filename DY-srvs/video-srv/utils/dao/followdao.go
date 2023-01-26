/*
 * @Date: 2023-01-26 10:46:28
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 10:53:00
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/dao/followdao.go
 * @Description: follows表的操作
 */
package dao

import (
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/models"
	"strconv"
)

/**
 * @description: 查询follows表中user_id与id匹配的字段数量，即用户的关注总数
 * @param {int64} id
 * @return {int64} count
 */
func CountFollow(id int64) int64 {
	var count int64
	global.DB.Model(&models.Follows{}).Where("user_id = ?", strconv.FormatInt(id, 10)).Count(&count)
	return count
}

/**
 * @description: 查询follows表中follower_id与id匹配的字段数量，即用户的粉丝总数
 * @param {int64} id
 * @return {int64} count
 */
func CountFollower(id int64) int64 {
	var count int64
	global.DB.Model(&models.Follows{}).Where("follower_id = ?", strconv.FormatInt(id, 10)).Count(&count)
	return count
}
