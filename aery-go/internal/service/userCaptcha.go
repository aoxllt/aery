package service

type IUserCaptcha interface {
	Generate(string) (string, error)
}

var localCaptcha IUserCaptcha

func ReturnCaptcha() IUserCaptcha {
	if localCaptcha == nil {
		panic("验证码接口未实现")
	}
	return localCaptcha
}

func RegisterCaptcha(newCaptcha IUserCaptcha) {
	localCaptcha = newCaptcha
}
