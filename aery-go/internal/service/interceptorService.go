package service

import "context"

type IinterceptorService interface {
	Changepassword(ctx context.Context) bool
}

var local IinterceptorService

func InterceptorServiceforController() IinterceptorService {
	if local == nil {
		panic("interceptorService接口未实现")
	}
	return local
}

func RegisterinterceptorService(service IinterceptorService) {
	local = service
}
