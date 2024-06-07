package service

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"regexp"
	dal "wu-user-backend/biz/dal/user"
	"wu-user-backend/biz/model/user"
	"wu-user-backend/biz/pack"
	"wu-user-backend/utils/core"
	"wu-user-backend/utils/crypt"
	"wu-user-backend/utils/pkg/errno"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

func IsAccountValidate(userAccount string) bool {
	regx := "[^a-zA-Z0-9]+"
	matched, err := regexp.MatchString(regx, userAccount)
	if err != nil {
		return false
	}
	return matched
}

func (s *UserService) Register(req *user.RegisterReq) (userId int64, err error) {
	//1. 用户在前端输入账户，密码以及校验码
	userAccount := req.UserAccount
	userPassword := req.UserPassword
	checkPassword := req.CheckPassword
	//2. 校验用户的账户，密码，校验码是否符合要求
	//a.非空
	if userAccount == "" || userPassword == "" || checkPassword == "" {
		return 0, errno.DataIsEmptyErr
	}
	//b. 账户长度不少于4位
	if len(userAccount) < 4 {
		return 0, errno.LenOfAccountErr
	}
	//c. 密码长度不少于8位
	if len(userPassword) < 8 {
		return 0, errno.LenOfPasswordErr
	}
	//d. 账户不能存在
	isExist := dal.IsAccountExist(userAccount)
	if isExist {
		return 0, errno.AccountIsExistErr
	}
	//e. 账户不包含特殊字符
	isMatched := IsAccountValidate(userAccount)
	if isMatched {
		return 0, errno.AccountHaveSCErr
	}
	//f. 密码和校验码相同
	if !bytes.Equal([]byte(userPassword), []byte(checkPassword)) {
		return 0, errno.PasswordNotEqualErr
	}
	//3. 对密码进行加密
	hashPassword := crypt.HashPassword(userPassword)
	//4. 向数据库插入用户数据
	id := dal.CreateUser(userAccount, hashPassword)
	if id <= 0 {
		return 0, errno.DataIsFailedErr
	}
	return id, nil
}

func (s *UserService) Login(req *user.LoginReq) (user *user.User, err error) {
	//1. 用户在前端输入账户密码
	userAccount := req.UserAccount
	userPassword := req.UserPassword
	//2. 校验用户的账户，密码，校验码是否符合要求
	//a. 非空
	if userAccount == "" || userPassword == "" {
		return nil, errno.DataIsEmptyErr
	}
	//b. 账户长度不少于4位
	if len(userAccount) < 4 {
		return nil, errno.LenOfAccountErr
	}
	//c. 密码长度不少于8位
	if len(userPassword) < 8 {
		return nil, errno.LenOfPasswordErr
	}
	//d. 账户必须存在
	isExist := dal.IsAccountExist(userAccount)
	if !isExist {
		return nil, errno.AccountIsNotExistErr
	}
	//e. 账户不包含特殊字符
	isMatched := IsAccountValidate(userAccount)
	if isMatched {
		return nil, errno.AccountHaveSCErr
	}
	//3. 对密码进行校验
	userModel := dal.GetUserByAccount(userAccount)
	if userModel == nil {
		return nil, errno.AccountIsNotExistErr
	}
	if crypt.ValidatePassword(userPassword, userModel.UserPassword) != nil {
		return nil, errno.PasswordErr
	}
	//4. 返回脱敏的用户信息
	u := pack.User(userModel)
	//5. 记录用户登录态
	e := core.SetUserLoginState(s.c, u)
	if e != nil {
		return nil, e
	}
	return u, nil
}

func (s *UserService) IsAdmin() bool {
	u, e := core.GetUserLoginState(s.c)
	if e != nil {
		return false
	}
	if !bytes.Equal([]byte(u.UserRole), []byte("admin")) {
		return false
	}
	return true
}

func (s *UserService) Search(req *user.SearchReq) (user *user.User, err error) {
	if !s.IsAdmin() {
		return nil, errno.ForbiddenErr
	}
	userAccount := req.UserAccount
	userModel := dal.GetUserByAccount(userAccount)
	if userModel == nil {
		return nil, errno.AccountIsNotExistErr
	}
	u := pack.User(userModel)
	return u, nil
}

func (s *UserService) Delete(req *user.DeleteReq) error {
	if !s.IsAdmin() {
		return errno.ForbiddenErr
	}
	userId := req.UserId
	isDelete := dal.DeleteUserById(userId)
	if !isDelete {
		return errno.DeleteDateFailedErr
	}
	return nil
}

func (s *UserService) GetCurrentUser(req *user.GetCurrentUserReq) (user *user.User, err error) {
	u, e := core.GetUserLoginState(s.c)
	if e != nil {
		return nil, e
	}
	return u, nil
}

func (s *UserService) Logout(req *user.LogoutReq) error {
	e := core.RemoveUserLoginState(s.c)
	if e != nil {
		return e
	}
	return nil
}
