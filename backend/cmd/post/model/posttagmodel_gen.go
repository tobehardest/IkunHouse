// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	postTagFieldNames          = builder.RawFieldNames(&PostTag{})
	postTagRows                = strings.Join(postTagFieldNames, ",")
	postTagRowsExpectAutoSet   = strings.Join(stringx.Remove(postTagFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	postTagRowsWithPlaceHolder = strings.Join(stringx.Remove(postTagFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	postTagModel interface {
		Insert(ctx context.Context, data *PostTag) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PostTag, error)
		Update(ctx context.Context, data *PostTag) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPostTagModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PostTag struct {
		Id         int64     `db:"id"`          // 主键ID
		TagId      int64     `db:"tag_id"`      // 标签id
		VideoId    int64     `db:"video_id"`    // 视频id
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 最后修改时间
	}
)

func newPostTagModel(conn sqlx.SqlConn) *defaultPostTagModel {
	return &defaultPostTagModel{
		conn:  conn,
		table: "`post_tag`",
	}
}

func (m *defaultPostTagModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPostTagModel) FindOne(ctx context.Context, id int64) (*PostTag, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", postTagRows, m.table)
	var resp PostTag
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

func (m *defaultPostTagModel) Insert(ctx context.Context, data *PostTag) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, postTagRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.TagId, data.VideoId)
	return ret, err
}

func (m *defaultPostTagModel) Update(ctx context.Context, data *PostTag) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, postTagRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.TagId, data.VideoId, data.Id)
	return err
}

func (m *defaultPostTagModel) tableName() string {
	return m.table
}