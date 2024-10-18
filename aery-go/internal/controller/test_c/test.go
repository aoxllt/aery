package test_c

import (
	"aery-go/api/test"
	"context"
	"fmt"
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

type TestController struct{}

var Test TestController

func (c *TestController) MysqlTest(ctx context.Context, req *test.TestReq) (res *test.TestRes, err error) {

	return
}
