package config

var Log logStruct

//日志配置结构
type logStruct struct {
	LogLevel  string `ini:"level"`
	Formatter string `ini:"formatter"`
}
