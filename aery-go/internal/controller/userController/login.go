package userController

import (
	"aery-go/api/user"
	"aery-go/internal/service/userService"
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
	return r.Cookie.Get("cookie").String()
}

func (c *LoginController) UserLogin(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error) {
	// 初始化响应结构体
	res = &user.LoginRes{}

	r := ghttp.RequestFromCtx(ctx)
	uuid := c.loginGetCooike(r)
	// 根据用户名和密码验证用户信息
	outcome, info, err := userService.User().VerifyUser(ctx, req.Username, req.Password, req.Captcha, uuid)
	if err != nil {
		// 打印错误信息并返回适当的错误
		fmt.Println("Error verifying user:", err)
		res.Status = false
		res.Message = "用户不存在或信息获取失败"
		return res, errors.New("用户不存在或信息获取失败")
	}
	// 检查是否已有 session
	existingSession, err := r.Session.Get("username")
	if err != nil {
		res.Status = false
		res.Message = "获取现有 session 失败"
		return res, errors.New("获取现有 session 失败")
	}

	// 如果已有 session，删除它
	if existingSession != nil {
		err = r.Session.Remove("username")
		if err != nil {
			res.Status = false
			res.Message = "删除现有 session 失败"
			return res, errors.New("删除现有 session 失败")
		}
	}

	// 创建新的 session
	err = r.Session.Set("username", req.Username)
	if err != nil {
		res.Status = false
		res.Message = "session 设置失败"
		return res, errors.New("session 设置失败")
	}
	// 填充响应数据
	res.Status = outcome
	res.Message = info
	return
}
