package model

import "gorm.io/gorm"

/*
	收藏表
*/

type Collect struct {
	gorm.Model
	VideoId uint   `gorm:"column:video_id"` // 收藏的视频id
	UserId  string `gorm:"column:user_id"`  // 用户id
}
