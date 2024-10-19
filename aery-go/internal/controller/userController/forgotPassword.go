package userController

import (
	"aery-go/api/user"
	"aery-go/internal/service/userService"
	"context"
)

type forgotPasswordController struct{}

var ForgotPassword = forgotPasswordController{}

func (c *forgotPasswordController) ForgotPasswordController(ctx context.Context, req *user.ForgotPasswordReq) (res *user.ForgotPasswordRes, err error) {
	res = &user.ForgotPasswordRes{}
	res.Status, res.Message = userService.ForgotPasswordforController().Check(ctx, req.UserName, req.Email, req.Captcha)
	return
}
