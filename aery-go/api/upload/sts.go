package upload

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gogf/gf/v2/frame/g"
)

type StsReq struct {
	g.Meta `path:"/sts" method:"get" tags:"获取sts"`
}
type StsRes struct {
	Sts *sts.AssumeRoleResponse `json:"sts"`
}
