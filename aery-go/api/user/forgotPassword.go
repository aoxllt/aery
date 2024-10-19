package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ForgotPasswordReq struct {
	g.Meta   `path:"/forgotpassword" method:"post" tags:"忘记密码"`
	UserName string `json:"username" v:"required#用户名为空"`
	Email    string `json:"email"  v:"required#email为空"`
	Captcha  string `json:"captcha" v:"required#验证码为空"`
}
type ForgotPasswordRes struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
