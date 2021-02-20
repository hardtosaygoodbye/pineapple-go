package config

var Wechat wechatStruct

type wechatStruct struct {
	AppID     string `ini:"app_id"`
	AppSecret string `ini:"app_secret"`
}
