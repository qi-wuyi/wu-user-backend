// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "wu-user-backend/biz/model/user"
	"wu-user-backend/biz/pack"
	service "wu-user-backend/biz/service/user"
	"wu-user-backend/utils/pkg/errno"
)

// Register .
// @router /api/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.RegisterResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}
	id, err := service.NewUserService(ctx, c).Register(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.RegisterResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}
	c.JSON(consts.StatusOK, user.RegisterResp{
		BaseResp: &user.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.SuccessMsg,
		},
		UserId: id,
	})
}

// Login .
// @router /api/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.LoginResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}
	u, err := service.NewUserService(ctx, c).Login(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.LoginResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}
	c.JSON(consts.StatusOK, user.LoginResp{
		BaseResp: &user.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.SuccessMsg,
		},
		User: u,
	})
}

// Search .
// @router /api/user/search [GET]
func Search(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.SearchReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.SearchResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}
	u, err := service.NewUserService(ctx, c).Search(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.SearchResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}

	c.JSON(consts.StatusOK, user.SearchResp{
		BaseResp: &user.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.SuccessMsg,
		},
		User: u,
	})
}

// Delete .
// @router /api/user/delete [POST]
func Delete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DeleteResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}
	e := service.NewUserService(ctx, c).Delete(&req)
	if e != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DeleteResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}

	c.JSON(consts.StatusOK, user.DeleteResp{
		BaseResp: &user.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.SuccessMsg,
		},
	})
}

// GetCurrentUser .
// @router /api/user/current [GET]
func GetCurrentUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetCurrentUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.GetCurrentUserResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}
	u, err := service.NewUserService(ctx, c).GetCurrentUser(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.GetCurrentUserResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}
	c.JSON(consts.StatusOK, user.GetCurrentUserResp{
		BaseResp: &user.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.SuccessMsg,
		},
		User: u,
	})
}

// Logout .
// @router /api/user/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LogoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.LogoutResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}
	e := service.NewUserService(ctx, c).Logout(&req)
	if e != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.LogoutResp{
			BaseResp: &user.BaseResp{
				Code: resp.Code,
				Msg:  resp.Msg,
			},
		})
		return
	}
	c.JSON(consts.StatusOK, user.LogoutResp{
		BaseResp: &user.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.SuccessMsg,
		},
	})
}
