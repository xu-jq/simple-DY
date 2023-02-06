/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 15:18:43
 * @FilePath: /simple-DY/DY-srvs/social-srv/model/base.go
 * @Description:
 */
package model

import (
	"database/sql/driver"
	"encoding/json"
	"simple-DY/DY-srvs/video-srv/models"
	"time"

	"gorm.io/gorm"
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseModel struct {
	ID        int32          `gorm:"primarykey;type:bigint" json:"id"` // 为什么使用int32， bigint
	CreatedAt time.Time      `gorm:"column:add_time" json:"-"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	IsDeleted bool           `json:"-"`
}

type Follows struct {
	ID         int64 `gorm:"primarykey;type:bigint" json:"id"`
	UserID     int64 `gorm:"column:user_id" json:"UserID"`
	FollowerID int64 `gorm:"column:follower_id" json:"FollowerID"`
}

type Messages struct {
	ID       int64     `gorm:"primarykey;type:bigint" json:"id"`
	UserID   int64     `gorm:"column:user_id" json:"user_id"`
	ToUserID int64     `gorm:"column:to_user_id" json:"to_user_id"`
	SentTime time.Time `gorm:"column:sent_time" json:"sent_time"`
	Content  string    `gorm:"column:content" json:"content"`
}

type FollowsAndUser struct {
	models.Users
	Follows
}
