package userService

import "context"

type IUserRegister interface {
	UserRegister(ctx context.Context, username string, password string, repassword string, email string, captcha string) (bool, string, error)
	GetRegisterCaptcha(ctx context.Context, email string) (bool, string, error)
}

var localRegister IUserRegister

func RegisterforController() IUserRegister {
	if localRegister == nil {
		panic("注册接口未实现")
	}
	return localRegister
}
func RegisterRegister(r IUserRegister) {
	localRegister = r
}
