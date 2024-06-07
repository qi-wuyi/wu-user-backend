package pack

import (
	"errors"
	"wu-user-backend/biz/model/user"
	"wu-user-backend/utils/pkg/errno"
)

func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{
		Code: err.ErrCode,
		Msg:  err.ErrMsg,
	}
}

func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMsg(err.Error())
	return baseResp(s)
}
