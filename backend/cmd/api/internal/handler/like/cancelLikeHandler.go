package like

import (
	"net/http"

	"IkunHouse/cmd/api/internal/logic/like"
	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CancelLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CancelLikeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := like.NewCancelLikeLogic(r.Context(), svcCtx)
		resp, err := l.CancelLike(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
