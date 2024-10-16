package userController

import (
	"aery-go/api/user"
	"aery-go/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
)

type CaptchaController struct{}

var Captcha CaptchaController

func (c *CaptchaController) getCooike(r *ghttp.Request) string {
	return r.Cookie.Get("uuid").String()
}

// 测试用
func (c *CaptchaController) setCooike(r *ghttp.Request) {
	r.Cookie.Set("uuid", "12345")
	return
}
func (c *CaptchaController) GetLoginCaptcha(ctx context.Context, req *user.CaptchaReq) (res *user.CaptchaRes, err error) {
	res = &user.CaptchaRes{}
	r := ghttp.RequestFromCtx(ctx)
	c.setCooike(r)
	uuid := c.getCooike(r)
	img, err := service.ReturnCaptcha().Generate(uuid)
	fmt.Println(uuid)
	if err != nil {
		res.CaptchaImg = ""
		return nil, err
	}
	res.CaptchaImg = img
	return
}
