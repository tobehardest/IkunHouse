package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"video_clip/cmd/api/internal/logic/comment"
	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"
)

func CommentDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewCommentDetailLogic(r.Context(), svcCtx)
		resp, err := l.CommentDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
