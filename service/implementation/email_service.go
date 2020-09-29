package implementation

import (
	"fmt"
	"log"
	"net/smtp"

	config "github.com/jacobhjustice/WeatherAlerts/model/configuration"
	"github.com/jacobhjustice/WeatherAlerts/service/specification"
)

type EmailService struct {
	specification.IEmailService
	Configuration *config.EmailConfiguration
	LogService    *specification.ILogService
}

func (e *EmailService) getAuth() smtp.Auth {
	return smtp.PlainAuth("", e.Configuration.Email, e.Configuration.Password, e.Configuration.Host)
}

func (e *EmailService) getAddress() string {
	return e.Configuration.Host + ":" + e.Configuration.Port
}

func (e EmailService) SendEmail(message string, toEmail ...string) error {
	fmt.Println("")

	auth := e.getAuth()
	outBytes := []byte(message)

	err := smtp.SendMail(e.getAddress(), auth, e.Configuration.Email, toEmail, outBytes)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Sent!")
	}

	return err
}
