package userService

import "context"

/*
service一般是和dao相关的，一个dao一个service
1.定义服务接口
2.定义一个接口变量，作为中间变量，用来接收logic返回的值并传递给controller
3.定义一个获取接口实例的函数，主要是对controller起作用
4.定义一个接口注册的方法，主要是对logic起作用
*/

// IUsers 用户服务接口定义
type IUsers interface {
	/*
		查询用entity，新增修改用do
	*/
	VerifyUser(ctx context.Context, username string, password string, captcha string, uuid string) (bool, string, error)
}

// 定义一个局部变量，保存用户服务的实例
var localUser IUsers

// User 返回已注册的 IUsers 接口实现
func User() IUsers {
	if localUser == nil {
		panic("User接口未注册或实现")
	}
	return localUser
}

// RegisterUser 注册 IUsers 接口的实现
func RegisterUser(user IUsers) {
	localUser = user
}
