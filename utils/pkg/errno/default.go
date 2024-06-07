package errno

var (
	Success = NewErrNo(SuccessCode, SuccessMsg)

	//未知微服务错误
	ServiceErr = NewErrNo(ServiceErrCode, "service is unable to start successfully")

	//鉴权失败
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "authorization failed")

	//用户未登录
	NotLoginErr = NewErrNo(NotLoginErrCode, "user not login")

	//请求数据不存在
	NotFoundErr = NewErrNo(NotFoundErrCode, "data is not found")

	//系统报错
	SystemErr = NewErrNo(SystemErrCode, "system error")

	//操作失败
	OperationErr = NewErrNo(OperationErrCode, "operation failed")

	//无访问权限
	ForbiddenErr = NewErrNo(ForbiddenErrCode, "no auth")
)

var (
	//未输入数据
	DataIsEmptyErr = NewErrNo(ParamErrCode, "data is not input")

	//账户输入长度不足
	LenOfAccountErr = NewErrNo(ParamErrCode, "len of account is less than 4")

	//密码输入长度不足
	LenOfPasswordErr = NewErrNo(ParamErrCode, "len of password is less than 8")

	//账户已存在
	AccountIsExistErr = NewErrNo(ParamErrCode, "account is exist,please rename your account")

	//账户不存在
	AccountIsNotExistErr = NewErrNo(ParamErrCode, "account is not exist,please register your account")

	//账户包含特殊字符
	AccountHaveSCErr = NewErrNo(ParamErrCode, "account contains special characters,please rename your account")

	//密码和校验码不同
	PasswordNotEqualErr = NewErrNo(ParamErrCode, "your checkPassword is not equal with your password")

	//密码错误
	PasswordErr = NewErrNo(ParamErrCode, "password entered incorrectly ")

	//插入数据失败
	DataIsFailedErr = NewErrNo(ParamErrCode, "data is not insert")

	//删除数据失败
	DeleteDateFailedErr = NewErrNo(ParamErrCode, "delete date is failed")

	//token签发错误
	TokenIsFailedErr = NewErrNo(ParamErrCode, "token is failed to signed")
)
