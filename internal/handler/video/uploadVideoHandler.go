package video

import (
	"net/http"

	"Qiniu/internal/logic/video"
	"Qiniu/internal/svc"
	"Qiniu/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewUploadVideoLogic(r.Context(), svcCtx)
		resp, err := l.UploadVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
