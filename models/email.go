package models

import "strings"

type EmailMessage struct {
	To                              []string
	From, Cc, Bcc, Subject, Message string
}

func (e EmailMessage) GenerateMessage() []byte {

	message := "From:" + e.From + "\n" +
		"To: " + strings.Join(e.To[:], ";") + "\n" +
		"Subject: " + e.Subject + "\n\n" +
		e.Message

	return []byte(message)

}
