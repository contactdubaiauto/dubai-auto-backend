package utils

import (
	"dubai-auto/internal/config"
	"net/smtp"
)

func SendEmail(title, subject, receiver string) error {
	message := []byte("Subject: " + title + "\r\n" +
		"\r\n" +
		subject + "\r\n")

	auth := smtp.PlainAuth("", config.ENV.SMTP_MAIL, config.ENV.SMTP_PASSWORD, config.ENV.SMTP_HOST)
	err := smtp.SendMail(config.ENV.SMTP_HOST+":"+config.ENV.SMTP_PORT, auth, config.ENV.SMTP_MAIL, []string{receiver}, message)
	return err
}
