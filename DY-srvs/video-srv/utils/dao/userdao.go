/*
 * @Date: 2023-01-25 21:11:11
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 21:52:48
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/dao/userdao.go
 * @Description: 对users数据表的操作
 */

package dao

import (
	"simple-DY/DY-srvs/video-srv/global"
	"simple-DY/DY-srvs/video-srv/models"
)

/**
 * @description: 通过name获取Users表的信息,由于name一定是唯一的，因此只有查找到和没有查找到两种情况，不会出现查询出多个的情况
 * @param {string} name
 * @return {models.Users} user
 */
func GetUserByName(name string) models.Users {
	// 数据库查询和更新的模板
	user := models.Users{}

	// 根据姓名查找数据库中的用户信息
	global.DB.Where("name = ?", name).Find(&user)

	return user
}

/**
 * @description: 通过指定用户名和密码，插入到Users表中，Id在表中为自增，返回id
 * @param {string} name
 * @param {string} password
 * @return {int64} id
 */
func InsertUser(name, password string) int64 {
	// 构建插入的结构体
	user := models.Users{
		Name:     name,
		Password: password,
	}
	// 插入到Users表中
	global.DB.Create(&user)

	return user.Id
}
