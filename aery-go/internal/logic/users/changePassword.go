package users

import (
	"aery-go/internal/dao"
	"aery-go/internal/model/entity"
	"aery-go/internal/service/userService"
	"aery-go/utility"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	userService.RegisterChangePasswordService(&Change{})
}

type Change struct{}

func (c Change) UpadatePassword(ctx context.Context, password string, username string) (bool, string, error) {
	// 查询用户信息
	user := entity.Users{}
	err := dao.Users.Ctx(ctx).Where("user_name=?", username).Scan(&user)
	if err != nil {
		return false, "错误", err
	}

	// 检查用户是否存在
	if user.UserName == "" {
		return false, "用户不存在", nil
	}

	// 解密数据库中的密码进行比较
	decryptedPasswd, err := utility.DecryptAES(user.UserPasswd)
	if err != nil {
		return false, "错误", err
	}

	// 检查新密码是否与旧密码相同
	if decryptedPasswd == password {
		return false, "新密码不能和原密码相同", nil
	}

	// 对新密码进行加密
	encryptedPasswd, err := utility.EncryptAES(password)
	if err != nil {
		return false, "错误", err
	}

	// 更新数据库中的密码
	_, err = dao.Users.Ctx(ctx).Where("user_name=?", username).Data(g.Map{
		"user_passwd": encryptedPasswd,
	}).Update()
	if err != nil {
		return false, "错误", err
	}

	return true, "密码更新成功", nil
}
