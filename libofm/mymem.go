package libofm

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	HOST        = "smtp.gmail.com"
	SERVER_ADDR = "smtp.gmail.com:587"
	USER        = "seafooler@gmail.com"
	PASSWD      = "yxwrtdxsfvwtocxe"
)

type Email struct {
	to      string "to"
	subject string "subject"
	msg     string "msg"
}

func NewEmail(to, subject, msg string) *Email {
	return &Email{to: to, subject: subject, msg: msg}
}

func SendEmail(email *Email) error {
	auth := smtp.PlainAuth("", USER, PASSWD, HOST)
	sendTo := strings.Split(email.to, ";")
	fmt.Println("Count of sendTo:", len(sendTo))
	done := make(chan error, 1024)

	go func() {
		defer close(done)
		for _, v := range sendTo {
			str := strings.Replace("From: "+USER+"~To:"+v+"~Subject:"+email.subject+"~~", "~", "\r\n", -1) + email.msg

			err := smtp.SendMail(
				SERVER_ADDR,
				auth,
				USER,
				[]string{v},
				[]byte(str),
			)
			done <- err
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Send well")
		}
	}()
	for i := 0; i < len(sendTo); i++ {
		<-done
	}
	return nil
}
