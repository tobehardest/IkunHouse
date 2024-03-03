package collect

import (
	"net/http"

	"IkunHouse/cmd/api/internal/logic/collect"
	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CollectListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollectListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := collect.NewCollectListLogic(r.Context(), svcCtx)
		resp, err := l.CollectList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
