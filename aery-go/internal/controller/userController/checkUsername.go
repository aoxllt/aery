package userController

import (
	"aery-go/api/user"
	"aery-go/internal/service/userService"
	"context"
)

type CheckUsernameController struct{}

var CheckUsername = CheckUsernameController{}

func (c *CheckUsernameController) UsernameCheck(ctx context.Context, req *user.UsernameCheckReq) (res *user.UsernameCheckRes, err error) {
	res = &user.UsernameCheckRes{}
	status := userService.CheckUsernameforController().CheckUsername(ctx, req.Username)
	if status {
		res.Status = true
		res.Message = "exist"
		return res, nil
	}
	res.Status = false
	res.Message = "not_exist"
	return res, nil
}
