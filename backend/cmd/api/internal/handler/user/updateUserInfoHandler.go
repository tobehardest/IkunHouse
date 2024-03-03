package user

import (
	"net/http"

	"IkunHouse/cmd/api/internal/logic/user"
	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"
	"IkunHouse/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
