package userService

import "context"

type ICheck interface {
	CheckUsername(ctx context.Context, username string) bool
}

var localCheckusername ICheck

func CheckUsernameforController() ICheck {
	if localCheckusername == nil {
		panic("checkusername接口未实现")
	}
	return localCheckusername
}
func RegisterCheckUsername(name ICheck) {
	localCheckusername = name
}
