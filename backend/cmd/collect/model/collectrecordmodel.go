package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CollectRecordModel = (*customCollectRecordModel)(nil)

type (
	// CollectRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollectRecordModel.
	CollectRecordModel interface {
		collectRecordModel
		withSession(session sqlx.Session) CollectRecordModel
	}

	customCollectRecordModel struct {
		*defaultCollectRecordModel
	}
)

// NewCollectRecordModel returns a model for the database table.
func NewCollectRecordModel(conn sqlx.SqlConn) CollectRecordModel {
	return &customCollectRecordModel{
		defaultCollectRecordModel: newCollectRecordModel(conn),
	}
}

func (m *customCollectRecordModel) withSession(session sqlx.Session) CollectRecordModel {
	return NewCollectRecordModel(sqlx.NewSqlConnFromSession(session))
}
