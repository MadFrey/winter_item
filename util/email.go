package util

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"time"
)

const (
	MailHOST = "smtp.qq.com"
	MailPORT = 465
	MailUSER = "1802515800@qq.com"
	MailPWD  = "oestcjpeqeprbdee"
)

func SendEmail(mailAddress string, subject string) (error, string) {
	m := gomail.NewMessage()
	m.SetHeader("From", MailUSER)
	m.SetHeader("subject", subject)
	m.SetHeader("To", mailAddress)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vscode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	m.SetBody("text/html", "您的验证码是"+vscode)
	d := gomail.NewDialer(MailHOST, MailPORT, MailUSER, MailPWD)
	err := d.DialAndSend(m)
	if err != nil {
		log.Println(err)
	}
	return err, vscode
}
