package messaging

import (
	"errors"
	"log"

	gomail "gopkg.in/mail.v2"
)

type emailMessenger struct {
	From string
	Key  string
}

func New(from, key string) Sender {
	if key == "" || from == "" {
		log.Fatal("key or from is empty")
	}

	return &emailMessenger{
		From: from,
		Key:  key,
	}
}

func (e *emailMessenger) Send(to string, message string) error {
	msg := gomail.NewMessage()

	msg.SetHeader("From", e.From)
	msg.SetHeader("To", to)
	// msg.SetHeader("Subject", message.Header)
	msg.SetBody("text/plain", message)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, e.From, e.Key)

	if err := dialer.DialAndSend(msg); err != nil {
		log.Println(err)
		return errors.New("can't this error")
	}
	
	return nil
}