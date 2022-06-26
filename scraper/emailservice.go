package main

import (
	"fmt"
	"net/smtp"
)

type EmailServiceSmtp struct {
	host    string
	port    int
	account string
	pass    string
}

func (s EmailServiceSmtp) SendEmailResourceChanged(resource *Resource) error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	msg := s.generateMessage(resource)
	return smtp.SendMail(addr, s.auth(), s.account, []string{resource.Email}, msg)
}

func (s EmailServiceSmtp) auth() smtp.Auth {
	return smtp.PlainAuth("", s.account, s.pass, s.host)
}

func (s EmailServiceSmtp) generateMessage(resource *Resource) []byte {
	msg := []byte(
		fmt.Sprintf(`To: %s\r\n`+
			`Subject: URL changed!\r\n`+
			`\r\n`+
			`Hi!\r\n`+
			`\r\n`+
			`The following resource has changed: \r\n`+
			`%s`, resource.Email, resource.Url))
	return msg
}
