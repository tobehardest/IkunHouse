package model

import "gorm.io/gorm"

// 点赞表 - 记录用户和视频的点赞关系
type Favorite struct {
	gorm.Model
	VideoId uint   `gorm:"video_id"` // 点赞的视频id
	UserId  string `gorm:"user_id"`  // 点赞人的id
}

func (f *Favorite) TableName() string {
	return "t_favorite"
}
