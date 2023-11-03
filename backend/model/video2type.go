package model

import "gorm.io/gorm"

// 视频以及类别映射
type Video_type struct {
	gorm.Model
	TypeId  uint `gorm:"column:type_id"`
	VideoId uint `gorm:"column:video_id"`
}
