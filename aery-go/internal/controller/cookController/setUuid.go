package cookController

import (
	"aery-go/api/cookie"
	"aery-go/internal/service"
	"context"
)

type UuidController struct{}

var Uuid = UuidController{}

func (c *UuidController) Setuuid(ctx context.Context, req *cookie.UuidReq) (res *cookie.UuidRes, err error) {
	res = &cookie.UuidRes{}
	status, err := service.CookieServiceforController().Getuuid(ctx)
	if err != nil {
		res.Status = false
		res.Message = "失败"
		return res, err
	}
	if status {
		res.Status = false
		res.Message = "已设置"
		return res, nil
	}
	service.CookieServiceforController().Setuuid(ctx)
	res.Status = true
	res.Message = "ok"
	return
}
