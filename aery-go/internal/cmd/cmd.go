package cmd

import (
	"aery-go/internal/controller/test_c"
	"aery-go/internal/controller/userController"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 创建一个HTTP服务器实例
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			// 路由分组
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareCORS)

				group.Bind(
					userController.Login,
				)
				group.Bind(
					userController.Captcha,
					test_c.Test,
					userController.Register,
				)

			})
			// 设置服务器端口（如果你有特定的端口需求）

			// 启动服务器
			s.Run()
			return nil
		},
	}
)
