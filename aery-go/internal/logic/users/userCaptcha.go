package users

import (
	"aery-go/internal/service/userService"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"math/rand"
	"strings"
)

func init() {
	userService.RegisterCaptcha(&sCaptcha{})
}

type sCaptcha struct {
}

func (s sCaptcha) Generate(uuid string) (string, error) {
	// 配置验证码
	driver := base64Captcha.NewDriverString(
		160,                               // 高度
		480,                               // 宽度
		0,                                 // 噪声数量
		base64Captcha.OptionShowSlimeLine, // 线条选项
		4,                                 // 随机字符串长度
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",                                     // 字符来源
		&color.RGBA{uint8(rand.Intn(56) + 200), uint8(rand.Intn(56) + 200), uint8(rand.Intn(56) + 200), 255}, // 背景颜色
		base64Captcha.DefaultEmbeddedFonts,                                                                   // 字体存储
		[]string{},                                                                                           // 使用的字体
	)
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	_, img, ans, err := c.Generate()
	if err != nil {
		return "", err
	}
	var ctx = gctx.New()
	//将生成的验证码和uuid存放在redis中，设置90秒过期
	_, err = g.Redis().Do(ctx, "HSET", "captcha:"+uuid, "ans", strings.ToLower(string(ans)))
	if err != nil {
		return "", err
	}
	_, err = g.Redis().Do(ctx, "EXPIRE", "captcha:"+uuid, 90)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	return img, nil
}
