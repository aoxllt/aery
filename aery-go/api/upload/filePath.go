package upload

import "github.com/gogf/gf/v2/frame/g"

type FilePathReq struct {
	g.Meta `path:"/upload" method:"post" tags:"获取文件路径"`
	Paths  string `json:"path"`
}
type FilePathRes struct {
	Status bool `json:"status"`
}
