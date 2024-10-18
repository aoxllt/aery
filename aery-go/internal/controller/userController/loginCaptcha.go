package userController

import (
	"aery-go/api/user"
	"aery-go/internal/service/userService"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
)

type CaptchaController struct{}

var Captcha CaptchaController

func (c *CaptchaController) getCooike(r *ghttp.Request) string {
	return r.Cookie.Get("cookie").String()
}

// 测试用
//
//	func (c *CaptchaController) setCooike(r *ghttp.Request) {
//		r.Cookie.Set("cookie", "12345")
//		return
//	}
func (c *CaptchaController) GetLoginCaptcha(ctx context.Context, req *user.CaptchaReq) (res *user.CaptchaRes, err error) {
	res = &user.CaptchaRes{}
	r := ghttp.RequestFromCtx(ctx)
	//c.setCooike(r)
	uuid := c.getCooike(r)
	fmt.Println(r.Cookie.Get("1uuid").String())
	fmt.Println(uuid)
	img, err := userService.ReturnCaptcha().Generate(uuid)
	fmt.Println(uuid)
	if err != nil {
		res.CaptchaImg = ""
		return nil, err
	}
	res.CaptchaImg = img
	return
}
