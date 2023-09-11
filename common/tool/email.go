package tool

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/pkg/errors"
	"net/smtp"
	"zero-shop/common/xerr"
)

const (
	sendEmailSubject = "邮箱验证码"
	sendEmailContent = "您的邮箱验证码为"
)

func SendEmailCode(fromEmail, toEmail, code, password, host string, port int64) error {
	e := email.NewEmail()
	//设置发送方邮件
	e.From = fmt.Sprintf("dj <%s>", fromEmail)
	// 设置接收方邮件
	e.To = []string{fmt.Sprintf("%s", toEmail)}
	// 设置主题
	e.Subject = sendEmailSubject
	// 设置邮件内容
	e.Text = []byte(fmt.Sprintf("%s : %s", sendEmailContent, code))
	// 设置相关配置
	err := e.Send(fmt.Sprintf("%s:%d", host, port),
		smtp.PlainAuth("", fromEmail, password, host))
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.USER_SEND_EMAIL_ERROR), "SendEmailCode ERROR: %+v", err)
	}
	return nil
}
