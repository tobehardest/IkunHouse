package comment

import (
	"net/http"

	"IkunHouse/cmd/api/internal/logic/comment"
	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
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
