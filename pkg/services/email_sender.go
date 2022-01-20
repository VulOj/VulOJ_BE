package services

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"strings"
)

func SendEmail(toEmail string, verifyCode string) (isfinished bool) {
	defer func() {
		if err := recover(); err != nil {
			emailInit()
			isfinished = false
		}
	}()
	emailInit()
	from := mail.Address{"Pivot Studio团队-楚天双创项目组", account}
	to := mail.Address{"亲爱的用户", toEmail}
	var err error
	if err = client.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = client.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	subj := "VulOJ平台注册"

	//===================================
	//Send a email template
	t, err := template.ParseFiles("vuloj_verification_mail.html")
	if err != nil {
		log.Panic(err)
	}
	buffer := new(bytes.Buffer)
	var data interface{}
	if err = t.Execute(buffer, data); err != nil {
		log.Panic(err)
	}

	//------------------------------------
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	con := strings.Replace(buffer.String(), "VerifyCodePlace", verifyCode, 1)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message += mime + con

	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write([]byte(message))

	if err != nil {
		log.Panic(err)
	}
	err = w.Close()
	if err != nil {
		log.Panic(err)
	}
	client.Quit()
	return true
}
