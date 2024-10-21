package userService

import (
	"context"
)

type IChangePasswordService interface {
	UpadatePassword(ctx context.Context, password string, username string) (bool, string, error)
}

var localchange IChangePasswordService

func ChangePasswordServiceforController() IChangePasswordService {
	if localchange == nil {
		panic("changepassword接口未实现")
	}
	return localchange
}

func RegisterChangePasswordService(newService IChangePasswordService) {
	localchange = newService
}
