package constant

type ResponseCode int

const (
	SUCCESS  ResponseCode = 0   //成功
	CODE_404 ResponseCode = 404 // 页面未找到
	CODE_500 ResponseCode = 500 // 服务器错误
	CODE_403 ResponseCode = 403 // 访问被拒绝
	CODE_401 ResponseCode = 401 // 认证失败
	CODE_400 ResponseCode = 400 // 未知错误

	USER_LOGIN_FAILED    ResponseCode = 1001 //登录失败
	USER_NOT_EXISTS      ResponseCode = 1002 //用户不存在
	USER_JWT_ERROR       ResponseCode = 1003 //登录生成jwt token失败
	USER_VERIFY_FAILD    ResponseCode = 1004 //jwt认证失败
	USER_JWT_PARSE_FAILD ResponseCode = 1005 //jwt解析失败

	REDIS_ERROR              ResponseCode = 5000 //redis 错误
	REDIS_KEY_NOT_EXISTS_ERR ResponseCode = 5001 //redis key不存在
)

var codeTextMap = map[ResponseCode]string{
	SUCCESS:              "成功",
	CODE_403:             "认证失败",
	CODE_404:             "页面不存在",
	CODE_500:             "服务器内部错误",
	USER_JWT_PARSE_FAILD: "登录失败",
}

func GetCodeText(code ResponseCode) string {
	if value, ok := codeTextMap[code]; ok {
		return value
	}
	return "Unkown code text"
}
