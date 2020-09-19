package service

import (
	"fmt"
	"log"
	"net/smtp"

	config "github.com/jacobhjustice/WeatherAlerts/model/configuration"
)

type IEmailService interface {
	SendEmail()
}

type EmailService struct {
	IEmailService
	Configuration *config.EmailConfiguration
}

func (e EmailService) getAuth() smtp.Auth {
	return smtp.PlainAuth("", e.Configuration.Email, e.Configuration.Password, e.Configuration.Host)
}

func (e EmailService) getAddress() string {
	return e.Configuration.Host + ":" + e.Configuration.Port
}

func (e EmailService) SendEmail(message string, toEmail ...string) {
	fmt.Println("")

	auth := e.getAuth()
	outBytes := []byte(message)

	err := smtp.SendMail(e.getAddress(), auth, e.Configuration.Email, toEmail, outBytes)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Sent!")
	}
}
