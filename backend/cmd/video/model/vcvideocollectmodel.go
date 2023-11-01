package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ VcVideoCollectModel = (*customVcVideoCollectModel)(nil)

type (
	// VcVideoCollectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVcVideoCollectModel.
	VcVideoCollectModel interface {
		vcVideoCollectModel
	}

	customVcVideoCollectModel struct {
		*defaultVcVideoCollectModel
	}
)

// NewVcVideoCollectModel returns a model for the database table.
func NewVcVideoCollectModel(conn sqlx.SqlConn) VcVideoCollectModel {
	return &customVcVideoCollectModel{
		defaultVcVideoCollectModel: newVcVideoCollectModel(conn),
	}
}
