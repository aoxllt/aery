package interceptor

import (
	"aery-go/internal/service"
	"context"
)

func init() {
	service.RegisterinterceptorService(&interceptorlogic{})
}

type interceptorlogic struct{}

func (i interceptorlogic) Changepassword(ctx context.Context) bool {
	//TODO implement me
	cookie := service.CookieServiceforController().GetCookie(ctx, "cpwd")
	if cookie == "" {
		return false
	}
	return true
}
