package core

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"wu-user-backend/biz/model/user"
	"wu-user-backend/utils/pkg/constants"
	"wu-user-backend/utils/pkg/errno"
)

func SetUserLoginState(c *app.RequestContext, data interface{}) error {
	session := sessions.Default(c)
	sessionTTl := 86400
	session.Set(constants.USER_LOGIN_STATE, data)
	session.Options(sessions.Options{MaxAge: sessionTTl})
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

func RemoveUserLoginState(c *app.RequestContext) error {
	session := sessions.Default(c)
	session.Delete(constants.USER_LOGIN_STATE)
	session.Options(sessions.Options{MaxAge: -1})
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

func GetUserLoginState(c *app.RequestContext) (*user.User, error) {
	session := sessions.Default(c)
	model := session.Get(constants.USER_LOGIN_STATE)
	if model == nil {
		return nil, errno.NotLoginErr
	}
	return model.(*user.User), nil
}
