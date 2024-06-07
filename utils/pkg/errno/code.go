package errno

const (
	SuccessCode = 0
	SuccessMsg  = "ok"

	// Error
	ServiceErrCode             = 10001 //未知微服务错误
	ParamErrCode               = 10002 // 请求参数错误
	AuthorizationFailedErrCode = 10003 //鉴权失败
	NotLoginErrCode            = 10004 //未登录
	NotFoundErrCode            = 10005 //请求数据不存在
	SystemErrCode              = 10006 //系统内部异常
	OperationErrCode           = 10007 //操作失败
	ForbiddenErrCode           = 10008 //禁止访问
)
