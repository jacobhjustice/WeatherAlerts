package messager
import (
    "log"
    "net/smtp"
    "os"
    "fmt"
)


func Send() {
    // smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
    // auth := smtp.PlainAuth("", "jhj.program@gmail.com", "testpassword123", "smtp.gmail.com")
    fmt.Println("Starting..")
    SetEnv()
    sender := os.Getenv("SENDER_EMAIL")
    host := os.Getenv("SENDER_HOST")
    password := os.Getenv("SENDER_PASS")
    port := os.Getenv("SENDER_PORT")

    auth := smtp.PlainAuth("", sender, password, host)

    // Here we do it all: connect to our server, set up a message and send it
	to := []string{"test@gmail.com"}
	msg := []byte("To: test@gmail.com\r\n" +
		"Subject:Test email\r\n" +
		"\r\n" +
		"Here’s a test email! I love you bug eyes\r\n")
	err := smtp.SendMail(host + ":" + port, auth, sender, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}