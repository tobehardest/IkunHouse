package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FdCommentModel = (*customFdCommentModel)(nil)

type (
	// FdCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFdCommentModel.
	FdCommentModel interface {
		fdCommentModel
	}

	customFdCommentModel struct {
		*defaultFdCommentModel
	}
)

// NewFdCommentModel returns a model for the database table.
func NewFdCommentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FdCommentModel {
	return &customFdCommentModel{
		defaultFdCommentModel: newFdCommentModel(conn, c, opts...),
	}
}
