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
	postFieldNames          = builder.RawFieldNames(&Post{})
	postRows                = strings.Join(postFieldNames, ",")
	postRowsExpectAutoSet   = strings.Join(stringx.Remove(postFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	postRowsWithPlaceHolder = strings.Join(stringx.Remove(postFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	postModel interface {
		Insert(ctx context.Context, data *Post) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Post, error)
		Update(ctx context.Context, data *Post) error
		Delete(ctx context.Context, id int64) error
		// 收藏/评论/点赞数量增加
		UpdateCommentNum(ctx context.Context, id int64, count int64) error
		UpdateLikeNum(ctx context.Context, id int64, count int64) error
		UpdateCollectNum(ctx context.Context, id int64, count int64) error
		// 批量获得视频
	}

	defaultPostModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Post struct {
		Id            int64          `db:"id"`             // 主键ID
		Uid           int64          `db:"uid"`            // 作者id
		Title         string         `db:"title"`          // 标题
		Content       sql.NullString `db:"content"`        // 内容
		Media         sql.NullString `db:"media"`          // 文件
		CoverUrl      sql.NullString `db:"cover_url"`      // 封面
		VideoSha256   sql.NullString `db:"video_sha256"`   // 视频哈希值
		Type          int64          `db:"type"`           // 帖子类型：1图文2视频3文章
		Address       sql.NullString `db:"address"`        // 地址名称
		Longitude     float64        `db:"longitude"`      // 经度
		Latitude      float64        `db:"latitude"`       // 纬度
		CommentNum    int64          `db:"comment_num"`    // 评论数
		LikeNum       int64          `db:"like_num"`       // 点赞数
		CollectNum    int64          `db:"collect_num"`    // 收藏数
		ShareNum      int64          `db:"share_num"`      // 分享数
		VideoDuration string         `db:"video_duration"` // 视频时长
		PublishTime   time.Time      `db:"publish_time"`   // 发布时间
		CreateTime    time.Time      `db:"create_time"`    // 创建时间
		UpdateTime    time.Time      `db:"update_time"`    // 最后修改时间
		Status        int64          `db:"status"`         // 状态 0:待审核 1:审核不通过 2:可见 3:私密 4:用户删除
	}
)

func newPostModel(conn sqlx.SqlConn) *defaultPostModel {
	return &defaultPostModel{
		conn:  conn,
		table: "`post`",
	}
}

func (m *defaultPostModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPostModel) FindOne(ctx context.Context, id int64) (*Post, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", postRows, m.table)
	var resp Post
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

func (m *defaultPostModel) Insert(ctx context.Context, data *Post) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, postRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.Title, data.Content, data.Media, data.CoverUrl, data.VideoSha256, data.Type, data.Address, data.Longitude, data.Latitude, data.CommentNum, data.LikeNum, data.CollectNum, data.ShareNum, data.VideoDuration, data.PublishTime, data.Status)
	return ret, err
}

func (m *defaultPostModel) Update(ctx context.Context, data *Post) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, postRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Uid, data.Title, data.Content, data.Media, data.CoverUrl, data.VideoSha256, data.Type, data.Address, data.Longitude, data.Latitude, data.CommentNum, data.LikeNum, data.CollectNum, data.ShareNum, data.VideoDuration, data.PublishTime, data.Status, data.Id)
	return err
}

func (m *defaultPostModel) tableName() string {
	return m.table
}

func (m *defaultPostModel) UpdateCommentNum(ctx context.Context, id int64, count int64) error {
	query := fmt.Sprintf("update %s set comment_num = comment_num + ? where id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, count, id)
	return err
}

func (m *defaultPostModel) UpdateLikeNum(ctx context.Context, id int64, count int64) error {
	query := fmt.Sprintf("update %s set like_num = like_num + ? where id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, count, id)
	return err
}

func (m *defaultPostModel) UpdateCollectNum(ctx context.Context, id int64, count int64) error {
	query := fmt.Sprintf("update %s set collect_num = collect_num + ? where id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, count, id)
	return err
}
