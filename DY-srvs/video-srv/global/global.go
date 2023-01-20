/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-20 19:20:52
 * @FilePath: /simple-DY/DY-srvs/video-srv/global/global.go
 * @Description: 全局变量
 */
package global

import (
	"simple-DY/DY-srvs/video-srv/config"
	"sync"

	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	GlobalConfig config.Config
	Wg           sync.WaitGroup
)

// type Users struct {
// 	Id       int
// 	Name     string
// 	Password string
// }

// func main() {

// 	// u := Users{
// 	// 	Id:       2,
// 	// 	Name:     "dfsfds",
// 	// 	Password: "fdsfd",
// 	// }
// 	// DB.Create(u)

// 	var t Users
// 	DB.Where("id = ?", 1).First(&t)
// 	fmt.Println("result: ", t)
// 	// fmt.Println(result)
// }
