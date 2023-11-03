package follow

import (
	"net/http"

	"Qiniu/internal/logic/follow"
	"Qiniu/internal/svc"
	"Qiniu/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UnFollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UnFollowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := follow.NewUnFollowLogic(r.Context(), svcCtx)
		resp, err := l.UnFollow(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
