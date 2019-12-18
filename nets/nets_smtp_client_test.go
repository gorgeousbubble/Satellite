package nets

import (
	"testing"
)

func TestMailSmtpSend(t *testing.T) {
	mail := MailSmtp{
		user:     "1029535012@qq.com",
		password: "amdepjytfocvbfbc",
		host:     "smtp.qq.com",
		port:     "465",
	}
	message := Message{
		from:        "1029535012@qq.com",
		to:          []string{"alopex6414@outlook.com", "1029535012@qq.com"},
		cc:          []string{},
		bcc:         []string{},
		subject:     "Satellite",
		body:        "hello,world!(Automatic send by Satellite smpt client.)",
		contentType: "text/plain;charset=utf-8",
		attachment: Attachment{
			name:        "test.jpg",
			contentType: "image/jpg",
			withFile:    false,
		},
	}
	err := mail.SendTLS(message)
	if err != nil {
		t.Errorf("Error send stmp mail:%v", err)
	}
}
