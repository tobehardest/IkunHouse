package follow

import (
	"net/http"

	"IkunHouse/cmd/api/internal/logic/follow"
	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"
	"IkunHouse/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := follow.NewFollowLogic(r.Context(), svcCtx)
		resp, err := l.Follow(&req)
		result.HttpResult(r, w, resp, err)
	}
}
