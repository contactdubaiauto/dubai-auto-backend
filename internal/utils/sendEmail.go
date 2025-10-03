package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail(title, subject, receiver string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := "berdalyyew99@gmail.com"
	password := "ykrg srux uiyb wjnx"
	message := []byte("Subject: " + title + "\r\n" +
		"\r\n" +
		subject + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{receiver}, message)

	if err != nil {
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
