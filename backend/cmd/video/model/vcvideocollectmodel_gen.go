// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	vcVideoCollectFieldNames          = builder.RawFieldNames(&VcVideoCollect{})
	vcVideoCollectRows                = strings.Join(vcVideoCollectFieldNames, ",")
	vcVideoCollectRowsExpectAutoSet   = strings.Join(stringx.Remove(vcVideoCollectFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	vcVideoCollectRowsWithPlaceHolder = strings.Join(stringx.Remove(vcVideoCollectFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	vcVideoCollectModel interface {
		Insert(ctx context.Context, data *VcVideoCollect) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*VcVideoCollect, error)
		Update(ctx context.Context, data *VcVideoCollect) error
		Delete(ctx context.Context, id int64) error
	}

	defaultVcVideoCollectModel struct {
		conn  sqlx.SqlConn
		table string
	}

	VcVideoCollect struct {
		Id     int64 `db:"id"`      // id
		Uid    int64 `db:"uid"`     // 用户id
		PostId int64 `db:"post_id"` // 视频id
	}
)

func newVcVideoCollectModel(conn sqlx.SqlConn) *defaultVcVideoCollectModel {
	return &defaultVcVideoCollectModel{
		conn:  conn,
		table: "`vc_video_collect`",
	}
}

func (m *defaultVcVideoCollectModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultVcVideoCollectModel) FindOne(ctx context.Context, id int64) (*VcVideoCollect, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", vcVideoCollectRows, m.table)
	var resp VcVideoCollect
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVcVideoCollectModel) Insert(ctx context.Context, data *VcVideoCollect) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, vcVideoCollectRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.PostId)
	return ret, err
}

func (m *defaultVcVideoCollectModel) Update(ctx context.Context, data *VcVideoCollect) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, vcVideoCollectRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Uid, data.PostId, data.Id)
	return err
}

func (m *defaultVcVideoCollectModel) tableName() string {
	return m.table
}
