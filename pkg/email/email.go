package email

import (
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	"galihwicaksono90/musikmarching-be/pkg/config"
	"net/smtp"
)

type Email interface {
	SendPurchaseInvoice(user *model.SessionUser) error
	Send(to string, subject string, body string) error
}

type email struct {
	config config.Config
}

// Send implements Email.
func (e *email) Send(to string, subject string, body string) error {
	return e.sendEmail(to, subject, body)
}

// SendInvoice implements Email.
func (e *email) SendPurchaseInvoice(user *model.SessionUser) error {
	subject := "Invoice"
	body := "Invoice body"

	return e.sendEmail(user.Email, subject, body)
}

func (e *email) sendEmail(to string, subject string, body string) error {
	fmt.Println(e.config)
	auth := smtp.PlainAuth(
		"",
		e.config.SmtpFrom,
		e.config.SmtpFromPassword,
		e.config.SmtpHost,
	)

	message := "Subject: " + subject + "\n" + body

	return smtp.SendMail(
		e.config.SmtpHost+":"+e.config.SmptPort,
		auth,
		e.config.SmtpFrom,
		[]string{to},
		[]byte(message),
	)
}

func NewEmail(config config.Config) Email {
	return &email{
		config,
	}
}
