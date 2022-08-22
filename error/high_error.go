package error

type HighError struct {
	Code  int
	Error error
}

/**
 * code定义
 * define rule = HighError[XXX]Code
 */
const (
	HighErrorSuccessCode      = 0
	HighErrorSqlErrorCode     = 1
	HighErrorParamErrorCode   = 2
	HighErrorFileErrorCode    = 3
	HighErrorServiceErrorCode = 4 // 服务器内部错误
	HighErrorNotAuthCode      = 5 // 没有登录
	HighErrorAuthFailedCode   = 6 // 权限验证错误
)
