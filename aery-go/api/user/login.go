package user

import "github.com/gogf/gf/v2/frame/g"

/*
1.request参数
2.respone参数
*/

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"登录" summary:"用户登录"` // 请求的元数据
	Username string                                                 `json:"username" v:"required#用户名不能为空"` // 用户名
	Password string                                                 `json:"password" v:"required#密码不能为空"`  // 密码
	Captcha  string                                                 `json:"captcha" v:"required#验证码为空"`    //验证码
}

// LoginRes 定义登录响应的结构
type LoginRes struct {
	Status  bool   `json:"status" ` //状态
	Message string `json:"message"` // 提示信息
}
