package user

import "github.com/gogf/gf/v2/frame/g"

type RegisterCaptchaReq struct {
	g.Meta `path:"/registercaptcha" method:"get" tags:"获取注册验证码"`
	Email  string `json:"email" v:"required#邮箱为空"`
}
type RegisterCaptchaRes struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
