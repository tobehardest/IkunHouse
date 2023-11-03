package video

import (
	"crypto/sha256"
	"io"
	"net/http"
	"video_clip/pkg/oss"

	"github.com/zeromicro/go-zero/rest/httpx"
	"video_clip/cmd/api/internal/logic/video"
	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"
)

func UploadVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// TODO 如果视频已经存在，则不再上传，而是上传封面

		// 获得视频
		media, file_header, err := r.FormFile("video")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		defer media.Close()
		kubo_cfg := &oss.KuboConfig{
			AccessKey:  svcCtx.Config.Oss.AccessKey,
			SecretKey:  svcCtx.Config.Oss.SecretKey,
			Bucket:     svcCtx.Config.Oss.Bucket,
			DomainName: svcCtx.Config.Oss.Domain,
		}
		cover, coverHeader, err := r.FormFile("cover")
		if err != nil {
			if err == http.ErrMissingFile {
				// 文件不存在，则设置默认封面
				req.CoverUrl, err = oss.GenerateDefaultCoverUrl(kubo_cfg, media)
				if err != nil {
					httpx.ErrorCtx(r.Context(), w, err)
					return
				}
			} else {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
		} else {
			// 上传封面
			hasher := sha256.New()
			_, err = io.Copy(hasher, cover)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
			hash := string(hasher.Sum(nil))
			req.CoverUrl, err = oss.UploadData(kubo_cfg, coverHeader, hash)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
		}

		// 计算视频的哈希值
		hasher := sha256.New()
		_, err = io.Copy(hasher, media)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 计算最后的ShA-256
		hash := string(hasher.Sum(nil))

		// 上传视频
		req.Media, err = oss.UploadData(kubo_cfg, file_header, hash)
		req.Sha256 = hash

		// TODO 解析上传者id

		l := video.NewUploadVideoLogic(r.Context(), svcCtx)
		resp, err := l.UploadVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
