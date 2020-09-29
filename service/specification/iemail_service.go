package specification

type IEmailService interface {
	SendEmail(message string, toEmail ...string) error
}
