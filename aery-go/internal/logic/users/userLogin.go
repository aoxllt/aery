package users

import (
	"aery-go/internal/dao"
	"aery-go/internal/model/entity"
	"aery-go/internal/service/userService"
	"aery-go/utility"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

/*
1.创建init函数调用service中获取接口的函数
2.创建一个结构体用来实现service中的要求
3.自动注入service中定义的方法
4.完成方法的逻辑
*/
func init() {
	userService.RegisterUser(&sUser{})
}

type sUser struct{}

func (s sUser) VerifyUser(ctx context.Context, username string, password string, captcha string, uuid string) (bool, string, error) {
	//TODO implement me
	// 查询用户信息
	// 定义用户结构体来存储查询结果
	var user entity.Users
	// 从数据库中查询用户名对应的用户信息
	err := dao.Users.Ctx(ctx).Where("user_name=?", username).Scan(&user)
	if err != nil {
		return false, "用户名或密码错误", err
	}
	//fmt.Println(cookie)
	ans, err := g.Redis().Do(ctx, "HGET", "captcha:"+uuid, "ans")
	if err != nil {
		return false, "验证码查询出错", err
	}
	//fmt.Println(ans)
	if strings.ToLower(captcha) != ans.String() {
		return false, "验证码错误或者已过期", nil
	}
	user.UserPasswd, err = utility.DecryptAES(user.UserPasswd)
	if err != nil {
		return false, "错误", err
	}
	// 获取数据库中的密码字段
	dbPassword := user.UserPasswd

	// 输出数据库中的密码
	if dbPassword != password {
		return false, "用户名或密码错误", nil
	}
	return true, "ok", nil
}
