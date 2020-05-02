package services

import (
	"bytes"
	"context"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/wingsico/movie_server/constants"
	"io"
	"mime/multipart"
)

func UploadFile(file multipart.File, name string)(url string, err error) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}
	// 生成最后的文件字节流
	fileStream := buf.Bytes()

	putPolicy := storage.PutPolicy{
		Scope: constants.Bucket,
	}
	mac := qbox.NewMac(constants.AccessKey, constants.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": name,
		},
	}
	dataLen := int64(len(fileStream))
	err = formUploader.Put(context.Background(), &ret, upToken, name, bytes.NewReader(fileStream), dataLen, &putExtra)
	if err != nil {
		return "", err
	}

	url = constants.Domain + ret.Key
	return url, nil
}