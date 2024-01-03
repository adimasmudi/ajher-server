package utils

import (
	"log"
	"net/mail"

	"gopkg.in/gomail.v2"
)

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Ajher <ajherteam@gmail.com>"
const CONFIG_AUTH_EMAIL = "ajherteam@gmail.com"
const CONFIG_AUTH_PASSWORD = "bucy vysh ojap iihd"

func SendMail(to []string, cc []string, subject, message string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", to...)
	mailer.SetAddressHeader("Cc", CONFIG_AUTH_EMAIL, CONFIG_AUTH_EMAIL)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	log.Println("Mail sent!")
	return nil
}
