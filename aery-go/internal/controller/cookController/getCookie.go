package cookController

import (
	"aery-go/api/cookie"
	"aery-go/internal/service"
	"context"
)

type GetCookie struct{}

var Getcookie = new(GetCookie)

func (c *GetCookie) Post(ctx context.Context, req *cookie.GetCookieReq) (res *cookie.GetCookieRes, err error) {
	res = &cookie.GetCookieRes{}
	res.CookieValue = service.CookieServiceforController().GetCookie(ctx, req.CookieKey)
	return
}
