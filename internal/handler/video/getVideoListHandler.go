package video

import (
	"net/http"

	"Qiniu/internal/logic/video"
	"Qiniu/internal/svc"
	"Qiniu/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetVideoListRequst
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewGetVideoListLogic(r.Context(), svcCtx)
		resp, err := l.GetVideoList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
