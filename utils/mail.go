package utils

import (
	"log"
	"net/mail"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func SendMail(to []string, cc []string, subject, message string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", to...)
	mailer.SetAddressHeader("Cc", os.Getenv("CONFIG_AUTH_EMAIL"), os.Getenv("CONFIG_AUTH_EMAIL"))
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)
	// mailer.Attach("./sample.png")

	port, err := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))

	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(
		os.Getenv("CONFIG_SMTP_HOST"),
		port,
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	log.Println("Mail sent!")
	return nil
}
