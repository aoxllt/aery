package cmd

import (
	"aery-go/internal/controller/cookController"
	"aery-go/internal/controller/interceptorController"
	"aery-go/internal/controller/test_c"
	"aery-go/internal/controller/userController"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
)

var (
	Main = gcmd.Command{
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 创建一个HTTP服务器实例
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.SetSessionStorage(gsession.NewStorageRedisHashTable(g.Redis()))
			// 路由分组
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareCORS)
				group.Bind(
					userController.Login,
					userController.Captcha,
					test_c.Test,
					userController.Register,
					cookController.Uuid,
					userController.CheckUsername,
					userController.ForgotPassword,
					interceptorController.ChangePasswordController,
					userController.ChangePasswordController,
					cookController.Getcookie,
				)

			})
			// 启动服务器
			s.Run()
			return nil
		},
	}
)
