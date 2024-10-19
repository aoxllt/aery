package user

import "github.com/gogf/gf/v2/frame/g"

type ChangePasswordReq struct {
	g.Meta     `path:"/changepassword" method:"post" tag:"修改密码"`
	Username   string `json:"username" v:"required"`
	Password   string `json:"password" v:"required"`
	Repassword string `json:"repassword" v:"required"`
}

type ChangePasswordRes struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
