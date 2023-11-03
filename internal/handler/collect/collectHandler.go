package collect

import (
	"net/http"

	"Qiniu/internal/logic/collect"
	"Qiniu/internal/svc"
	"Qiniu/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CollectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := collect.NewCollectLogic(r.Context(), svcCtx)
		resp, err := l.Collect(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
