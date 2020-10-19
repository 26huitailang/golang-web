package notify

import (
	"log"
	"net/smtp"
)

// Message ...
type Message struct {
	Title   string
	From    string
	To      []string
	Content []byte
}

var (
	username = "username"
	password = "password"
	host     = "127.0.0.1"
	port     = ":1025"
)

// Example ...
func Example(msg *Message) {
	auth := smtp.PlainAuth("", username, password, host)
	err := smtp.SendMail(host+port, auth, msg.From, msg.To, msg.Content)
	if err != nil {
		log.Fatal(err)
	}
}
