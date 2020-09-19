package service

import (
    "log"
    "net/smtp"
    "fmt"
)

type IEmailService interface {
	SendEmail()
}

type EmailService struct {
	IEmailService
	email string 
	host string
	password string
	port string
} 

func (e EmailService) getAuth() smtp.Auth {
	return smtp.PlainAuth("", e.email, e.password, e.host)
}

func (e EmailService) getAddress() string {
	return e.host + ":" + e.port
}


func (e EmailService) SendEmail(message string, toEmail ...string) {
    fmt.Println("")
    
	
	auth := e.getAuth()
	outBytes := []byte(message)

	err := smtp.SendMail(e.host + ":" + e.port, auth, e.email, toEmail, outBytes)
	if err != nil {
		log.Fatal(err)
	} else {
        fmt.Println("Sent!")
    }
}