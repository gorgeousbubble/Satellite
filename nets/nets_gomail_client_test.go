package nets

import (
	"testing"
)

func TestMailGoSend(t *testing.T) {
	mail := MailGo{
		user:     "1029535012@qq.com",
		password: "amdepjytfocvbfbc",
		host:     "smtp.qq.com",
		port:     465,
	}
	message := MailMessage{
		from:    "1029535012@qq.com",
		to:      "alopex6414@outlook.com",
		cc:      []string{},
		bcc:     []string{},
		subject: "Satellite",
		body:    "hello,world!(Automatic send by Satellite gomail client.)",
	}
	err := mail.Send(message)
	if err != nil {
		t.Errorf("Error send gomail mail:%v", err)
	}
}
