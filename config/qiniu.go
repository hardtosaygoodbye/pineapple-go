package config

var Qiniu qiniuStruct

type qiniuStruct struct {
	AccessKey string `ini:"access_key"`
	SecretKey string `ini:"secret_key"`
	Expire    int    `ini:"expire"`
}
