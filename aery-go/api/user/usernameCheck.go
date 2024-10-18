package user

import "github.com/gogf/gf/v2/frame/g"

type UsernameCheckReq struct {
	g.Meta   `path:"/checkusername" method:"get" tags:"检查用户名"`
	Username string `json:"username" v:"required#用户名为空"`
}

type UsernameCheckRes struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
