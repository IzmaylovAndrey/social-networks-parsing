package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEmail(login string, name string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "containerum@exonlab.ru", "Containerum Go Course")
	m.SetAddressHeader("To", login, name)
	m.SetHeader("Subject", "You are successfully registered!")
	m.SetBody("text/plain", "We are glad to meet you. Enjoy our project.")

	d := gomail.NewDialer("smtp.yandex.ru", 465, "containerum@exonlab.ru", "aAyPbDUX")

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Error sending email: %s", err)
		return err
	}
	return nil
}
