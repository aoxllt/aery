package test

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestReq struct {
	g.Meta `path:"/" method:"get" tags:"测试" summary:"数据库测试"`
}
type TestRes struct {
}
