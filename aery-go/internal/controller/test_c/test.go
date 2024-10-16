package test_c

import (
	"aery-go/api/test"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

type TestController struct{}

var Test TestController

func (c *TestController) MysqlTest(ctx context.Context, req *test.TestReq) (res *test.TestRes, err error) {
	// 查询数据库
	md := g.Model("users")
	user, err := md.Where("user_name=?", "盛森").All()
	if err != nil {
		// 如果查询失败，返回错误
		return nil, err
	}
	// 或者将结果转换为您需要的格式
	fmt.Println(user)
	return
}
