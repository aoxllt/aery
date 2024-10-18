package userController

import (
	"aery-go/api/user"
	"aery-go/internal/service/userService"
	"context"
)

type RegisterController struct{}

var Register RegisterController

func (c *RegisterController) GetRegisterCaptcha(ctx context.Context, req *user.RegisterCaptchaReq) (res *user.RegisterCaptchaRes, err error) {
	res = &user.RegisterCaptchaRes{}
	res.Status, res.Message, err = userService.RegisterforController().GetRegisterCaptcha(ctx, req.Email)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return res, err
	}
	return
}

func (c *RegisterController) UserRegister(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error) {
	res = &user.RegisterRes{}
	res.Status, res.Message, err = userService.RegisterforController().UserRegister(ctx, req.Username, req.Password, req.Repassword, req.Email, req.Captcha)
	if err != nil {
		res.Status = false
		res.Message = "注册失败，请稍后重试"
		return res, nil
	}
	return res, nil
}
