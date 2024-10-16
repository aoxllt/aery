package userController

import (
	"aery-go/api/user"
	"aery-go/internal/service"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type RegisterController struct{}

var Register RegisterController

func (c *RegisterController) GetRegisterCaptcha(ctx context.Context, req *user.RegisterCaptchaReq) (res *user.RegisterCaptchaRes, err error) {
	res = &user.RegisterCaptchaRes{}
	r := ghttp.RequestFromCtx(ctx)
	uuid := Captcha.getCooike(r)
	res.Status, res.Message, err = service.RegisterforController().GetRegisterCaptcha(ctx, uuid, req.Email)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return res, err
	}
	return
}

func (c *RegisterController) UserRegister(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error) {
	res = &user.RegisterRes{}
	return
}
