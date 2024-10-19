package userController

import (
	"aery-go/api/user"
	"aery-go/internal/service"
	"aery-go/internal/service/userService"
	"context"
)

type ChangepasswordController struct{}

var ChangePasswordController ChangepasswordController

func (c *ChangepasswordController) Post(ctx context.Context, req *user.ChangePasswordReq) (res *user.ChangePasswordRes, err error) {
	res = &user.ChangePasswordRes{}
	status := service.InterceptorServiceforController().Changepassword(ctx)
	if !status {
		res.Status = false
		res.Message = "非法访问"
		return
	}
	if req.Password == req.Repassword {
		res.Status, res.Message, err = userService.ChangePasswordServiceforController().UpadatePassword(ctx, req.Password, req.Username)
		if err != nil || res.Status == false {
			res.Status = false
			res.Message = "错误"
			return
		}
		return res, nil
	}
	res.Status = false
	res.Message = "密码不一致"
	return
}
