package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            string `gorm:"primaryKey;size:32"`
	Username      string `gorm:"column:username;size:16;uniqueIndex"`
	Password      string `gorm:"column:password;size:16"`
	Avatar        string `gorm:"column:avatar;size:256"`          // 头像
	Introduce     string `gorm:"column:introduce;size:256"`       // 个人简介
	FollowCount   int64  `gorm:"column:follow_count;default:0"`   // 关注数
	FollowerCount int64  `gorm:"column:follower_count;default:0"` // 粉丝数
	WorkCount     int64  `gorm:"column:work_count;default:0"`     // 作品数
	CollectCount  int64  `gorm:"column:collect_count;default:0"`  // 收藏数
	CreatedAt     time.Time
	UpdateAt      time.Time
	DeleteAt      gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
	return "t_user"
}
