package users

import (
	"aery-go/internal/dao"
	"aery-go/internal/model/entity"
	"aery-go/internal/service"
	"aery-go/internal/service/userService"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
)

func init() {
	userService.RegisterForgotPasswordService(&forgotpassword{})
}

type forgotpassword struct{}

func (f forgotpassword) Check(ctx context.Context, username string, email string, captcha string) (bool, string) {
	//TODO implement me
	user := entity.Users{}
	err := dao.Users.Ctx(ctx).Where("user_name=?", username).Where("user_email=?", email).Scan(&user)
	if err != nil {
		return false, "错误"
	}
	if strconv.Itoa(user.UserId) == "" {
		return false, "用户名和邮箱不匹配"
	}
	code, err := g.Redis().Do(ctx, "HGET", "RegisterCaptcha:"+email, "ans")
	if err != nil {
		return false, "错误"
	}
	if code.String() == captcha {
		service.CookieServiceforController().SetCookie(ctx, "cpwd")
		return true, "ok"
	}
	return false, "验证码错误"
}
