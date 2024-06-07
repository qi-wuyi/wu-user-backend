package errno

import (
	"fmt"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code = %d, err_msg = %s", e.ErrCode, e.ErrMsg)
}

func (e ErrNo) WithMsg(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{ErrCode: code, ErrMsg: msg}
}
