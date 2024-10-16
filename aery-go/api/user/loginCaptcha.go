package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tag:"Captcha"`
}
type CaptchaRes struct {
	CaptchaImg string `json:"captcha"`
}
