package uploadController

import (
	"aery-go/api/upload"
	"aery-go/utility"
	"context"
)

type StsController struct{}

var Sts StsController

func (c *StsController) Get(ctx context.Context, req *upload.StsReq) (res *upload.StsRes, err error) {
	res = &upload.StsRes{}
	res.Sts, err = utility.Oss.Getsts()
	if err != nil {
		return
	}
	return
}
