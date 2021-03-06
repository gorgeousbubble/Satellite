package nets

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net"
	"net/smtp"
	"strings"
	"time"
)

type Mail interface {
	Auth()
	Send(message Message) error
	SendTLS(message Message) error
}

type MailSmtp struct {
	user     string
	password string
	host     string
	port     string
	auth     smtp.Auth
}

type Message struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	contentType string
	attachment  Attachment
}

type Attachment struct {
	name        string
	contentType string
	withFile    bool
}

func (mail *MailSmtp) Auth() {
	mail.auth = smtp.PlainAuth("", mail.user, mail.password, mail.host)
}

func (mail *MailSmtp) Send(message Message) error {
	mail.Auth()
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.from
	Header["To"] = strings.Join(message.to, ";")
	Header["Cc"] = strings.Join(message.cc, ";")
	Header["Bcc"] = strings.Join(message.bcc, ";")
	Header["Subject"] = message.subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, Header)

	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + message.contentType + "\r\n"
	body += "\r\n" + message.body + "\r\n"
	buffer.WriteString(body)

	if message.attachment.withFile {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + message.attachment.name + "\"\r\n"
		buffer.WriteString(attachment)
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln(err)
			}
		}()
		mail.writeFile(buffer, message.attachment.name)
	}

	buffer.WriteString("\r\n--" + boundary + "--")
	smtp.SendMail(mail.host+":"+mail.port, mail.auth, message.from, message.to, buffer.Bytes())
	return nil
}

func (mail *MailSmtp) SendTLS(message Message) error {
	mail.Auth()
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.from
	Header["To"] = strings.Join(message.to, ";")
	Header["Cc"] = strings.Join(message.cc, ";")
	Header["Bcc"] = strings.Join(message.bcc, ";")
	Header["Subject"] = message.subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, Header)

	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + message.contentType + "\r\n"
	body += "\r\n" + message.body + "\r\n"
	buffer.WriteString(body)

	if message.attachment.withFile {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + message.attachment.name + "\"\r\n"
		buffer.WriteString(attachment)
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln(err)
			}
		}()
		mail.writeFile(buffer, message.attachment.name)
	}
	buffer.WriteString("\r\n--" + boundary + "--")

	// send mail as tls...
	c, err := SmtpDial(mail.host + ":" + mail.port)
	if err != nil {
		log.Println("Error create client dial with server:", err)
		return err
	}
	defer c.Close()
	// fill auth
	if mail.auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(mail.auth); err != nil {
				log.Println("Error during AUTH:", err)
				return err
			}
		}
	}
	// fill from
	if err = c.Mail(message.from); err != nil {
		log.Println("Error during From:", err)
		return err
	}
	// fill to
	for _, addr := range message.to {
		if err = c.Rcpt(addr); err != nil {
			log.Println("Error during To:", err)
			return err
		}
	}
	// fetch data
	w, err := c.Data()
	if err != nil {
		log.Println("Error during Data:", err)
		return err
	}
	// write message
	_, err = w.Write(buffer.Bytes())
	if err != nil {
		log.Println("Error during write message:", err)
		return err
	}
	// close writer
	err = w.Close()
	if err != nil {
		log.Println("Error during close writer:", err)
		return err
	}
	return c.Quit()
}

func (mail *MailSmtp) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}

func (mail *MailSmtp) writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}

func SmtpDial(addr string) (c *smtp.Client, err error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Error dial:", err)
		return c, err
	}
	// parse host and port
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}
