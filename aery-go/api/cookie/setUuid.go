package cookie

import "github.com/gogf/gf/v2/frame/g"

type UuidReq struct {
	g.Meta `path:"/uuid" method:"get" tags:"获取uuid" summary:"设置uuid"`
}
type UuidRes struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
