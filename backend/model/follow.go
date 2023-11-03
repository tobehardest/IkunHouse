package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	UserId     string `gorm:"column:user_id;"`    // 关注者
	FollowerId string `gorm:"column:follower_id"` // 被关注者id
}
