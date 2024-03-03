package video

import (
	"IkunHouse/cmd/video/rpc/videoclient"
	"net/http"

	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 24

type UplocadVideoLogic struct {
	logx.Logger
	r      *http.Request
	svcCtx *svc.ServiceContext
}

func NewUplocadVideoLogic(r *http.Request, svcCtx *svc.ServiceContext) *UplocadVideoLogic {
	return &UplocadVideoLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		svcCtx: svcCtx,
	}
}

func (l *UplocadVideoLogic) UplocadVideo(req *types.UploadVideoReq) (resp *types.UploadVideoRes, err error) {
	// todo: add your logic here and delete this line

	//userId := ctxdata.GetUidFromCtx(l.r.Context())
	l.r.ParseMultipartForm(maxFileSize)
	file, header, err := l.r.FormFile("avatar")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	logx.Infof("upload file: %+v, file size: %d, MIME header: %+v",
		header.Filename, header.Size, header.Header)

	//fileName := header.Filename
	avatarBytes := make([]byte, header.Size)
	n, err := file.Read(avatarBytes)
	if err != nil {
		// todo
	}
	if int64(n) < header.Size {
		// todo
	}

	l.svcCtx.VideoClient.UploadVideo(l.r.Context(), &videoclient.UploadVideoReq{
		Avatar: avatarBytes,
	})
	return
}
