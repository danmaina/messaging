package models

import "strings"

type EmailMessage struct {
	To      []string `json:"to"`
	From    string   `json:"from"`
	Cc      []string `json:"cc"`
	Bcc     []string `json:"bcc"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

// TODO: Add Branded Templates to use

func (e EmailMessage) GenerateMessage() []byte {

	message := "From:" + e.From + "\n" +
		"To: " + strings.Join(e.To[:], ";") + "\n" +
		"Cc: " + strings.Join(e.Cc[:], ";") + "\n" +
		"Bcc: " + strings.Join(e.Bcc[:], ";") + "\n" +
		"Subject: " + e.Subject + "\n\n" +
		e.Message

	return []byte(message)

}
