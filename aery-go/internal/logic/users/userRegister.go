package users

import (
	"aery-go/internal/dao"
	"aery-go/internal/model/entity"
	"aery-go/internal/service/userService"
	"aery-go/utility"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"net/smtp"
	"strings"
)

const (
	SMTPHost     = "smtp.qq.com"      // SMTP 服务器地址，比如 smtp.gmail.com 或 smtp.qq.com
	SMTPPort     = "587"              // SMTP 服务器端口，常见的有 25、465、587
	SMTPUser     = "aoxllt@qq.com"    // 发件人邮箱
	SMTPPassword = "vplxwtqszdbhdieb" // 发件人邮箱密码或授权码
)

// 生成验证码
func generateVerificationCode() string {
	return gconv.String(gtime.TimestampNano() % 1000000) // 简单生成 6 位数字验证码
}

// 发送邮件验证码
func sendVerificationCodeEmail(to, code string) error {
	// 构建邮件内容
	subject := "【盛森科技】"
	body := fmt.Sprintf("你的验证码为: %s", code)
	msg := strings.Join([]string{
		"From: " + SMTPUser,
		"To: " + to,
		"Subject: " + subject,
		"",
		body,
	}, "\r\n")

	// 连接到 SMTP 服务器并发送邮件
	auth := smtp.PlainAuth("", SMTPUser, SMTPPassword, SMTPHost)
	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, SMTPUser, []string{to}, []byte(msg))
	return err
}

func init() {
	userService.RegisterRegister(&UserRegister{})
}

type UserRegister struct{}

func (u UserRegister) UserRegister(ctx context.Context, username string, password string, repassword string, email string, captcha string) (bool, string, error) {
	//TODO implement me
	if password != repassword {
		return false, "密码不一致", nil
	}
	code, err := g.Redis().Do(ctx, "HGET", "RegisterCaptcha:"+email, "ans")
	if err != nil {
		return false, "查询出错", err
	}
	if code.String() != captcha {
		return false, "验证码错误或者已过期", nil
	}
	password, err = utility.EncryptAES(password)
	if err != nil {
		return false, "错误", err
	}
	userDate := &entity.Users{
		UserName:     username,
		UserPasswd:   password,
		UserRole:     "1",
		UserStatus:   "0",
		UserEmail:    email,
		NewLoginTime: gtime.Now(),
	}
	result, err := dao.Users.Ctx(ctx).Where("user_name=?", username).One()
	if err != nil {
		return false, "错误", err
	}
	if !result.IsEmpty() {
		return false, "用户名已存在", nil
	}
	result1, err := dao.Users.Ctx(ctx).Where("user_email=?", email).One()
	if err != nil {
		return false, "错误1", err
	}
	if !result1.IsEmpty() {
		return false, "该邮箱已注册", nil
	}
	insert, err := dao.Users.Ctx(ctx).Data(userDate).Insert()
	if err != nil {
		return false, "注册失败", err
	}
	rowsAffected, err := insert.RowsAffected()
	if err != nil {
		return false, "获取影响行数失败", err
	}

	if rowsAffected > 0 {
		// 插入成功
		_, err = g.Redis().Do(ctx, "EXPIRE", "RegisterCaptcha:"+email, 0)
		if err != nil {
			return false, "错误2", nil
		}
		return true, "注册成功", nil
	} else {
		// 没有插入任何行
		return false, "注册失败", err
	}
}

func (u UserRegister) GetRegisterCaptcha(ctx context.Context, email string) (bool, string, error) {
	//TODO implement me
	code := generateVerificationCode()
	err := sendVerificationCodeEmail(email, code)
	if err != nil && !strings.Contains(err.Error(), "short response: \u0000\u0000\u0000\u001A\u0000\u0000\u0000") {
		return false, "验证码发送失败", err
	}
	//fmt.Println(code)
	_, err = g.Redis().Do(ctx, "HSET", "RegisterCaptcha:"+email, "ans", code)
	if err != nil {
		return false, "错误", err
	}
	_, err = g.Redis().Do(ctx, "EXPIRE", "RegisterCaptcha:"+email, 180)
	if err != nil {
		return false, "错误", err
	}
	return true, "ok", nil
}
