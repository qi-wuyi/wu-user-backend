// Code generated by hertz generator.

package main

import (
	"encoding/gob"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
	"wu-user-backend/biz/dal"
	"wu-user-backend/biz/model/user"
	"wu-user-backend/utils/pkg/constants"
	"wu-user-backend/utils/pkg/log"
)

func init() {
	log.LogConfig()
	dal.Init()
	gob.Register(&user.User{})
}

func main() {
	h := server.Default(server.WithHostPorts(constants.EndPort))
	store := cookie.NewStore([]byte("secret"))
	h.Use(sessions.New("mysession", store))
	register(h)
	h.Spin()
}
