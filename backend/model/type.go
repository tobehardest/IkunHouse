package model

import "gorm.io/gorm"

// 类别表 - 视频类别的字典
type Type struct {
	gorm.Model
	Name string `gorm:"column:name"` // 类别名称 - 例如体育/娱乐等等
}

func (t *Type) TableName() string {
	return "t_type"
}
