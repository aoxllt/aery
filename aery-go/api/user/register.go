package user

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta     `path:"/register" methon:"post" tags:"注册" summary:"用户注册"`
	Username   string `json:"username " v:"required#用户名不能为空"`
	Password   string `json:"password" v:"required#密码不能为空"`
	Repassword string `json:"repassword" v:"required#确认为空"`
	Email      string `json:"email" v:"required#邮箱为空"`
	Captcha    string `json:"captcha" v:"required#验证码为空"`
}

type RegisterRes struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
