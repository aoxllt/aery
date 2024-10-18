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
	//r.AddCookie(&http.Cookie{
	//	Name:        "uuid",
	//	Value:       setuuid,
	//	Quoted:      false,
	//	Path:        "/",
	//	Domain:      "10.22.72.216",
	//	Expires:     time.Time{},
	//	RawExpires:  "",
	//	MaxAge:      60,
	//	Secure:      false,
	//	HttpOnly:    false,
	//	SameSite:    0,
	//	Partitioned: false,
	//	Raw:         "",
	//	Unparsed:    nil,
	//})
	//r.Cookie.SetCookie("uuid", setuuid, "10.22.72.216", "/", 3600)
	r.Cookie.Set("uuid", setuuid)
	return true
}
