package logic

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"

	localtypes "video_clip/cmd/search/rpc/internal/types"

	"video_clip/cmd/search/rpc/internal/svc"
	"video_clip/cmd/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchShortVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchShortVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchShortVideoLogic {
	return &SearchShortVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchShortVideoLogic) SearchShortVideo(in *search.SearchShortVideoReq) (*search.SearchShortVideoRes, error) {
	searchRet, err := l.svcCtx.Es.Search().
		Index(localtypes.ShortVideoIndex).
		Query(&types.Query{
			Match: map[string]types.MatchQuery{
				//"title":       {Query: in.Query},
				//"content":     {Query: in.Query},
				//"author_name": {Query: in.Query},
				"name": {Query: "李华"},
			},
		}).
		//From(int(in.Cursor)).
		//Size(int(in.PageSize)).
		Do(l.ctx)
	if err != nil {

	}

	svs := make([]localtypes.TestMsg, len(searchRet.Hits.Hits), len(searchRet.Hits.Hits))
	err = json.Unmarshal(searchRet.Hits.Hits[0].Source_, &svs)
	if err != nil {

	}
	return &search.SearchShortVideoRes{}, nil
}
