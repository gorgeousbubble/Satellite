package nets

import (
	"testing"
)

func TestMailSmtpSend(t *testing.T) {
	mail := MailSmtp{
		user:     "Alopex6414@outlook.com",
		password: "hyx13313541292",
		host:     "smtp.office365.com",
		port:     "587",
	}
	message := Message{
		from:        "Alopex6414@outlook.com",
		to:          []string{"1029535012@qq.com"},
		cc:          []string{},
		bcc:         []string{},
		subject:     "Satellite",
		body:        "hello,world!",
		contentType: "text/plain;charset=utf-8",
		attachment: Attachment{
			name:        "test.jpg",
			contentType: "image/jpg",
			withFile:    false,
		},
	}
	err := mail.Send(message)
	if err != nil {
		t.Errorf("Error send stmp mail:%v", err)
	}
}
