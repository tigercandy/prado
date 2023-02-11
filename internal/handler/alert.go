package handler

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tigercandy/prado/global"
	"github.com/tigercandy/prado/internal/pkg/logger"
	"github.com/tigercandy/prado/internal/proposal"
	"github.com/tigercandy/prado/pkg/mail"
	"html/template"
	"time"
)

type BodyData struct {
	URL   string
	ID    string
	Msg   string
	Stack string
	Year  int
}

func NotifyHandler() func(msg *proposal.AlertMessage) {
	return func(msg *proposal.AlertMessage) {
		cfg := global.App.Config.Mail
		if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Password == "" || cfg.To == "" {
			logger.Error("Mail config error")
			return
		}
		subject, body, err := parsingTemplate(
			msg.Method,
			msg.Host,
			msg.Uri,
			msg.TraceID,
			msg.ErrMsg,
			msg.ErrStack,
		)
		if err != nil {
			logger.Error("email template err", err)
			return
		}
		options := &mail.Options{
			Host:     cfg.Host,
			Port:     cfg.Port,
			User:     cfg.User,
			Password: cfg.Password,
			To:       cfg.To,
			Subject:  subject,
			Body:     body,
		}
		if err := mail.SendMail(options); err != nil {
			logger.Error("send mail failed", errors.WithStack(err))
		}
		return
	}
}

func parsingTemplate(method, host, uri, id string, msg interface{}, stack string) (subject, body string, err error) {
	mailData := &BodyData{
		URL:   fmt.Sprintf("%s %s%s", method, host, uri),
		ID:    id,
		Msg:   fmt.Sprintf("%+v", msg),
		Stack: stack,
		Year:  time.Now().Year(),
	}

	subject = fmt.Sprintf("【系统告警】-%s", uri)

	tmpl, err := template.ParseFiles("./assets/template/alert_mail.html")
	if err != nil {
		return
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, mailData)
	if err != nil {
		return
	}

	return subject, buf.String(), nil
}
