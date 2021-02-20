package service

import (
	"pineapple-go/config"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

type qiniuService struct {
}

var QiniuService = qiniuService{}

func (qs qiniuService) GetToken() string {
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "swiftwhale"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = uint64(config.Qiniu.Expire) //示例2小时有效期

	mac := qbox.NewMac(config.Qiniu.AccessKey, config.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
