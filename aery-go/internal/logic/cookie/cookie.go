package cookie

import (
	"aery-go/internal/service"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
)

func init() {
	service.RegisterCookieService(&Cookie{})
}

type Cookie struct{}

func (c Cookie) GetCookie(ctx context.Context, s string) string {
	//TODO implement me
	r := ghttp.RequestFromCtx(ctx)
	cookie := r.Cookie.Get(s)
	return cookie.String()
}

func (c Cookie) SetCookie(ctx context.Context, s string) bool {
	//TODO implement me
	cookie := uuid.NewString()
	r := ghttp.RequestFromCtx(ctx)
	r.Cookie.Set(s, cookie)
	return true
}

func (c Cookie) Getuuid(ctx context.Context) (bool, error) {
	//TODO implement me
	r := ghttp.RequestFromCtx(ctx)
	getuuid := r.Cookie.Get("uuid")
	if getuuid.IsEmpty() {
		return false, nil
	}
	return true, nil
}

func (c Cookie) Setuuid(ctx context.Context) bool {
	//TODO implement me
	setuuid := uuid.NewString()
	r := ghttp.RequestFromCtx(ctx)
	r.Cookie.Set("uuid", setuuid)
	return true
}
