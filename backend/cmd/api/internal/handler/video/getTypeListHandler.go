package video

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"video_clip/cmd/api/internal/logic/video"
	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"
)

func GetTypeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTypeListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewGetTypeListLogic(r.Context(), svcCtx)
		resp, err := l.GetTypeList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
