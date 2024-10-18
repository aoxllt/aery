package users

import (
	"aery-go/internal/dao"
	"aery-go/internal/model/entity"
	"aery-go/internal/service/userService"
	"context"
)

func init() {
	userService.RegisterCheckUsername(&LusenameCheck{})
}

type LusenameCheck struct{}

func (l LusenameCheck) CheckUsername(ctx context.Context, username string) bool {
	//TODO implement me
	user := entity.Users{
		UserId:       0,
		UserName:     "",
		UserPasswd:   "",
		UserRole:     "",
		UserStatus:   "",
		UserEmail:    "",
		NewLoginTime: nil,
	}
	err := dao.Users.Ctx(ctx).Where("user_name=?", username).Scan(&user)
	if err != nil {
		return false
	}
	return true
}
