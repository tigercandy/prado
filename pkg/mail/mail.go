package mail

import (
	"gopkg.in/gomail.v2"
	"strings"
)

type Options struct {
	Host     string
	Port     int
	User     string
	Password string
	To       string
	Subject  string
	Body     string
}

func SendMail(options *Options) error {
	m := gomail.NewMessage()

	// 发件人
	m.SetHeader("From", options.User)
	// 收件人及抄送
	mailTo := strings.Split(options.To, ",")
	m.SetHeader("To", mailTo...)
	// 主题
	m.SetHeader("Subject", options.Subject)
	// 正文
	m.SetBody("text/html", options.Body)

	d := gomail.NewDialer(options.Host, options.Port, options.User, options.Password)

	return d.DialAndSend(m)
}
