package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"video_clip/cmd/api/internal/logic/auth"
	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"
	"video_clip/pkg/result"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		result.HttpResult(r, w, resp, err)
	}
}
