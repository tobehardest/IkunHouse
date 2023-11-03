package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Url           string `gorm:"column:url;size:256"`             // 视频路径
	CoverUrl      string `gorm:"column:url;size:256"`             // 封面路径
	Title         string `gorm:"column:title"`                    // 视频标题
	TypeId        uint   `gorm:"column:type_id;default:-1"`       // 所属的类型id,默认为-1，即为未选择类别
	AuthorId      string `gorm:"column:author_id"`                // 作者id
	FavoriteCount int64  `gorm:"column:favorite_count;default:0"` // 视频点赞总数
	CommentCount  int64  `gorm:"column:comment_count;default:0"`  // 视频评论总数
}

func (v *Video) TableName() string {
	return "t_video"
}
