package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ VideoTagModel = (*customVideoTagModel)(nil)

type (
	// VideoTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoTagModel.
	VideoTagModel interface {
		videoTagModel
	}

	customVideoTagModel struct {
		*defaultVideoTagModel
	}
)

// NewVideoTagModel returns a model for the database table.
func NewVideoTagModel(conn sqlx.SqlConn) VideoTagModel {
	return &customVideoTagModel{
		defaultVideoTagModel: newVideoTagModel(conn),
	}
}
