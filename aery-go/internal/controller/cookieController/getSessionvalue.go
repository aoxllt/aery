package cookieController

import (
	"aery-go/api/cookie"
	"aery-go/internal/service"
	"context"
)

type GetSessionValueController struct{}

var GetSessionValue GetSessionValueController

func (c *GetSessionValueController) Post(ctx context.Context, req *cookie.GetsessionReq) (res *cookie.GetsessionRes, err error) {
	res = &cookie.GetsessionRes{}
	value := service.CookieServiceforController().GetCookie(ctx, "aery")
	res.Value, err = service.RedisServiceforController().HGetValue(ctx, value, "username")
	if err != nil {
		return nil, err
	}
	return res, nil
}
