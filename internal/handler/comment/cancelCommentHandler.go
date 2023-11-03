package comment

import (
	"net/http"

	"Qiniu/internal/logic/comment"
	"Qiniu/internal/svc"
	"Qiniu/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CancelCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CancelCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewCancelCommentLogic(r.Context(), svcCtx)
		resp, err := l.CancelComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
