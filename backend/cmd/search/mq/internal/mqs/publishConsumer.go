package mqs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"video_clip/cmd/search/mq/internal/svc"
	"video_clip/cmd/search/mq/internal/types"
)

type ShortVideoPublishConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *ShortVideoPublishConsumer {
	return &ShortVideoPublishConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortVideoPublishConsumer) Consume(key, val string) error {
	var msg *types.TestMsg
	err := json.Unmarshal([]byte(val), &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}
	return nil
}

func (l *ShortVideoPublishConsumer) UpdateShortVideoIndex(ctx context.Context, data []*types.TestMsg) error {

	return l.BatchAddToIndex(ctx, data)
}

func (l *ShortVideoPublishConsumer) BatchAddToIndex(ctx context.Context, data []*types.TestMsg) error {
	if len(data) == 0 {
		return nil
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: l.svcCtx.Es.Client,
		Index:  types.ShortVideoIndex,
	})
	if err != nil {
		return err
	}
	for _, d := range data {
		v, err := json.Marshal(d)
		if err != nil {
			return err
		}

		payload := fmt.Sprintf(`{"doc":%s,"doc_as_upsert":true}`, string(v))
		err = bi.Add(ctx, esutil.BulkIndexerItem{
			Action:     "update",
			DocumentID: fmt.Sprintf("%d", d.Id),
			Body:       strings.NewReader(payload),
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem) {
			},
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *ShortVideoPublishConsumer) DeleteFromIndex(ctx context.Context, data []*types.TestMsg) error {

	return nil
}
