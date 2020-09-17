package messager
import (
    "log"
    "net/smtp"
    "fmt"
    "github.com/spf13/viper"
)


func Send() {
    fmt.Println("Starting..")
    sender := viper.GetString("messager.sender.email")
    host := viper.GetString("messager.sender.host")
    password := viper.GetString("messager.sender.pass")
    port := viper.GetString("messager.sender.port")

    auth := smtp.PlainAuth("", sender, password, host)

    // Here we do it all: connect to our server, set up a message and send it
	to := []string{"test@gmail.com"}
	msg := []byte("To: test@gmail.com\r\n" +
		"Subject:What's up??\r\n" +
		"\r\n" +
		"Woah! I sent this from a computer\r\n")
	err := smtp.SendMail(host + ":" + port, auth, sender, to, msg)
	if err != nil {
		log.Fatal(err)
	} else {
        fmt.Println("Sent!")
    }
}