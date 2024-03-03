package auth

import (
	"net/http"

	"IkunHouse/cmd/api/internal/logic/auth"
	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"
	"IkunHouse/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewUserRegisterLogic(r.Context(), svcCtx)
		resp, err := l.UserRegister(&req)
		result.HttpResult(r, w, resp, err)
	}
}
