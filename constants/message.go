package constants

//goland:noinspection ALL
const (
	AuthSuccess           = "用户认证成功"
	CreateUserSuccess     = "创建用户成功"
	CreateUserFail        = "创建用户失败"
	UserPassValidatorFail = "密码校验失败！（至少8位，包含大小写、数字、特殊字符）"
	UsernameValidatorFail = "用户名校验失败！只能是字母或字母+数字（不能有下划线、符号等）"
	NoSuchUser            = "用户不存在！"
	PasswordFail          = "密码错误！"
	LoginSuccess          = "登录成功"
	LoginError            = "用户名或密码错误"
	TokenExpired          = "token已过期"
	TokenInvalid          = "无效的token"
	TokenMalformed        = "token格式错误"
	TokenSignatureInvalid = "token签名无效"
	TokenParasError       = "token解析失败"
	TokenValid            = "token验证通过"
)
