package service

import "context"

type ICookieService interface {
	Getuuid(ctx context.Context) (bool, error)
	Setuuid(ctx context.Context) bool
}

var localCookieService ICookieService

func CookieServiceforController() ICookieService {
	if localCookieService == nil {
		panic("cookie接口未实现")
	}
	return localCookieService
}

func RegisterCookieService(c ICookieService) {
	localCookieService = c
}
