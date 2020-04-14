package email

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

var (
	host     string
	port     string
	sender   string
	password string
)

var auth smtp.Auth

func send(to string, token string) error {
	host = "smtp.gmail.com"
	port = "587"
	sender = os.Getenv("GMAIL")
	password = os.Getenv("PASSWORD")

	auth = smtp.PlainAuth("Syahrul", sender, password, "smtp.gmail.com")
	templateData := struct {
		JWT string
	}{
		JWT: token,
	}
	r := NewRequest([]string{to}, "Verify your email @ OnceCard", "")
	err := r.ParseTemplate("template.html", templateData)
	if err != nil {
		return err
	}
	_, err = r.SendEmail()
	if err != nil {
		return err
	}
	log.Println("Sent verification email to " + to)
	return nil
}

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject string, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, sender, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
