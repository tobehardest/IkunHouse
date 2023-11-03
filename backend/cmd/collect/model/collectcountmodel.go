package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CollectCountModel = (*customCollectCountModel)(nil)

type (
	// CollectCountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollectCountModel.
	CollectCountModel interface {
		collectCountModel
		withSession(session sqlx.Session) CollectCountModel
	}

	customCollectCountModel struct {
		*defaultCollectCountModel
	}
)

// NewCollectCountModel returns a model for the database table.
func NewCollectCountModel(conn sqlx.SqlConn) CollectCountModel {
	return &customCollectCountModel{
		defaultCollectCountModel: newCollectCountModel(conn),
	}
}

func (m *customCollectCountModel) withSession(session sqlx.Session) CollectCountModel {
	return NewCollectCountModel(sqlx.NewSqlConnFromSession(session))
}
