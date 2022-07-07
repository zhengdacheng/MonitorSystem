package service

import (
	"gopkg.in/gomail.v2"
)

type EmailConnector struct {
	User      string
	PassToken string
	Host      string
	Port      int
}

type EmailBody struct {
	MailTo  []string
	Subject string
	Content string
}

type SendUtil interface {
	Send(conn EmailConnector, body EmailBody) error
}

type SendEmail struct {
}

func (send *SendEmail) SendMail(connector EmailConnector, body EmailBody) error {
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(connector.User, "监控系统"))
	m.SetHeader("To", body.MailTo...)
	m.SetHeader("Subject", body.Subject)
	m.SetBody("text/html", body.Content)

	d := gomail.NewDialer(connector.Host, connector.Port, connector.User, connector.PassToken)

	err := d.DialAndSend(m)
	return err
}
