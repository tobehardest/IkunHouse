package oss

import (
	"context"
	"errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"os"
	"os/exec"
)

type KuboConfig struct {
	AccessKey  string
	SecretKey  string
	Bucket     string
	DomainName string // 域名
}

// TODO 获得封面url
func UploadData(kubo_cfg *KuboConfig, file *multipart.FileHeader, key string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	putPolicy := storage.PutPolicy{
		Scope: kubo_cfg.Bucket,
	}
	mac := qbox.NewMac(kubo_cfg.AccessKey, kubo_cfg.SecretKey)
	// 获取上传凭证
	upToken := putPolicy.UploadToken(mac)
	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 华南区
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	bucketManager := storage.NewBucketManager(mac, &cfg)
	fileInfo, err := bucketManager.Stat(kubo_cfg.Bucket, key)
	if err == nil && fileInfo.Fsize != 0 {
		// 文件已在云端存在
		return "", errors.New("云端已存在")
	}

	formUploader := storage.NewFormUploader(&cfg)
	// 上传后返回的结果
	ret := storage.PutRet{}
	// 额外参数
	putExtra := storage.PutExtra{}
	// 设置自定义key，可以指定上传目录以及文件名和后缀
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)

	// 上传后的文件访问路径
	url := kubo_cfg.DomainName + "/" + ret.Key
	return url, err
}

func GenerateDefaultCoverUrl(cfg *KuboConfig, media multipart.File) (string, error) {
	// 使用ffmpeg生成
	// CreateTemp 会保证生成唯一的文件夹
	thumbnaiFile, err := os.CreateTemp("tmp", "cover_*.jpg")
	defer thumbnaiFile.Close()
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-vframes", "1", thumbnaiFile.Name())
	cmd.Stdin = media
	if err := cmd.Run(); err != nil {
		return "", err
	}
	file_info, err := os.Stat(thumbnaiFile.Name())
	if err != nil {
		return "", err
	}
	// 创建虚拟文件头
	fileHeader := &multipart.FileHeader{
		Filename: thumbnaiFile.Name(),
		Size:     file_info.Size(),
	}

	// 上传对象存储
	url, err := UploadData(cfg, fileHeader, "")
	return url, err
}

// 如果为私人空间，需要创建私有链接
//func GetDownloadUrl(cfg *KuboConfig,key string) string {
//	var url string
//	mac := qbox.NewMac(cfg.AccessKey,cfg.SecretKey)
//	domain := cfg.DomainName
//	// 一小时有效
//	deadline := time.Now().Add(time.Second*3600).Unix()
//	privateAccessURL := storage.MakePrivateURL(mac,domain,key,deadline)
//}
