package userService

import "context"

type IforgotPasswordService interface {
	Check(ctx context.Context, username string, password string, captcha string) (bool, string)
}

var localforgot IforgotPasswordService

func ForgotPasswordforController() IforgotPasswordService {
	if localforgot == nil {
		panic("forgotpassword接口未实现")
	}
	return localforgot
}
func RegisterForgotPasswordService(service IforgotPasswordService) {
	localforgot = service
}
