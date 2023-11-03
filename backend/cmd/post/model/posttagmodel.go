package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PostTagModel = (*customPostTagModel)(nil)

type (
	// PostTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostTagModel.
	PostTagModel interface {
		postTagModel
	}

	customPostTagModel struct {
		*defaultPostTagModel
	}
)

// NewPostTagModel returns a model for the database table.
func NewPostTagModel(conn sqlx.SqlConn) PostTagModel {
	return &customPostTagModel{
		defaultPostTagModel: newPostTagModel(conn),
	}
}
