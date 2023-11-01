package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"video_clip/cmd/api/internal/logic/comment"
	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"
)

func CommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewCommentLogic(r.Context(), svcCtx)
		resp, err := l.Comment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
