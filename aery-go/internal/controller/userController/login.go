package userController

import (
	"aery-go/api/user"
	"aery-go/internal/service"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
)

/*
1.定义控制器
2.实例化一个对象
3.方法实现
*/
type LoginController struct{}

var Login = LoginController{}

func (c *LoginController) loginGetCooike(r *ghttp.Request) string {
	return r.Cookie.Get("uuid").String()
}

func (c *LoginController) UserLogin(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error) {
	// 初始化响应结构体
	res = &user.LoginRes{}

	r := ghttp.RequestFromCtx(ctx)
	uuid := c.loginGetCooike(r)
	// 根据用户名和密码验证用户信息
	outcome, info, err := service.User().VerifyUser(ctx, req.Username, req.Password, req.Captcha, uuid)
	if err != nil {
		// 打印错误信息并返回适当的错误
		fmt.Println("Error verifying user:", err)
		return nil, errors.New("用户不存在或信息获取失败")
	}

	// 填充响应数据
	res.Status = outcome
	res.Message = info
	return
}
