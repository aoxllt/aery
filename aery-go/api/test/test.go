package test

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestReq struct {
	g.Meta `path:"/test" method:"get" tags:"测试" summary:"测试"`
}
type TestRes struct {
}
