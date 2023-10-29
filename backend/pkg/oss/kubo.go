package oss

import (
	"bytes"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadData(key string, data []byte) {
	// 空间名称
	putPolicy := storage.PutPolicy{
		Scope: "bucket",
	}
	mac := qbox.NewMac("accesskey", "secretKey")
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 选择空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否开启https
	cfg.UseHTTPS = false
	// 是否使用CDN上传
	cfg.UseCdnDomains = false

	bucketManager := storage.NewBucketManager(mac, &cfg)
	fileInfo, err := bucketManager.Stat("bucket", key)
	if err == nil && fileInfo.Fsize != 0 {
		// 当文件在云端存在则不上传
		return
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.NewFormUploader(&cfg)
	putExtra := storage.PutExtra{}
	length := int64(len(data))
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), length, &putExtra)
	if err != nil {
		return
	}
}

//func GetDownloadUrl(key string) string {
//
//}
