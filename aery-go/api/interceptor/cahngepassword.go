package interceptor

import "github.com/gogf/gf/v2/frame/g"

type ChangePasswordReq struct {
	g.Meta `path:"/interceptor/changepassword" method:"get" tags:"修改密码拦截器"`
}
type ChangePasswordRes struct {
	Status bool `json:"status"`
}
