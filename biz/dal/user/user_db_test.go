package dal

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
	"wu-user-backend/biz/dal"
)

func TestIsAccountExist(t *testing.T) {
	dal.Init()
	userAccount := "test1"
	b := IsAccountExist(userAccount)
	spew.Dump(b)
}

func TestCreateUser(t *testing.T) {
	dal.Init()
	userAccount := "test1"
	userPassword := "12345678"
	id := CreateUser(userAccount, userPassword)
	spew.Dump(id)
}

func TestGetUserByAccount(t *testing.T) {
	dal.Init()
	userAccount := "test1"
	u := GetUserByAccount(userAccount)
	spew.Dump(u)
}

func TestDeleteUserById(t *testing.T) {
	dal.Init()
	id := int64(4)
	b := DeleteUserById(id)
	spew.Dump(b)
}
