package nets

import (
	"gopkg.in/gomail.v2"
)

type MailGo struct {
	user     string
	password string
	host     string
	port     int
}

type MailMessage struct {
	from    string
	to      string
	cc      []string
	bcc     []string
	subject string
	body    string
}

func (m *MailGo) Send(message MailMessage) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", message.from)
	mail.SetHeader("To", message.to)
	mail.SetHeader("Subject", message.subject)
	mail.SetBody("text/html", message.body)
	dialer := gomail.NewDialer(m.host, m.port, m.user, m.password)
	err := dialer.DialAndSend(mail)
	return err
}
