package utils

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

// SendEmail fromEmail is senders email address
// fromAccount is senders account
// fromEmailPasswd is senders email passwd
// host is address of server
// port is the default port of smtp server
// receiveEmail is email address to be sent
func SendEmail(isHtml bool, fromEmail string, fromAccount string, fromEmailPasswd string, host string, port int, receiveEmail string, message string) (bool, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", receiveEmail)
	m.SetHeader("Subject", "Lee")
	if isHtml {
		m.SetBody("text/html", message)
	} else {
		m.SetBody("text/plain", message)
	}
	d := gomail.NewDialer(host, port, fromAccount, fromEmailPasswd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}
