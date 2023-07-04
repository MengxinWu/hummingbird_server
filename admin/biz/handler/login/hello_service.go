// Code generated by hertz generator.

package login

import (
	"admin/biz/dal"
	login "admin/biz/model/login"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Login .
// @router /login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var (
		req     login.LoginReq
		account *login.Account
		err     error
	)
	resp := new(login.LoginResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// 获取用户
	if account, err = dal.GetAccount(req.Name); err != nil {

	}
	// 验证账号
	if account == nil {

	}
	// 验证密码
	if req.Passwd != account.Passwd {
		resp.Code =
			c.JSON(consts.StatusOK)
	}
	resp.ID = account.Id
	c.JSON(consts.StatusOK, resp)
}