package cookie

import "github.com/gogf/gf/v2/frame/g"

type GetsessionReq struct {
	g.Meta `path:"/getsession" method:"post" tags:"获取session值"`
}
type GetsessionRes struct {
	Value string `json:"value"`
}
