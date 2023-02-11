package test

import (
	"github.com/tigercandy/prado/pkg/mail"
	"testing"
)

func TestSendMail(t *testing.T) {
	options := &mail.Options{
		Host:     "smtp.qq.com",
		Port:     465,
		User:     "xx@qq.com",
		Password: "", // 邮箱授权码
		To:       "xx@163.com",
		Subject:  "prado",
		Body:     "body",
	}
	err := mail.SendMail(options)
	if err != nil {
		t.Error("Mail send error", err)
		return
	}
	t.Log("Mail send success")
}
