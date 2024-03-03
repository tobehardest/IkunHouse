package video

import (
	"net/http"

	"IkunHouse/cmd/api/internal/logic/video"
	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UplocadVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewUplocadVideoLogic(r, svcCtx)
		resp, err := l.UplocadVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
