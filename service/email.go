package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"gopkg.in/gomail.v2"
	"html/template"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var dial *gomail.Dialer

type Credential struct {
	Server string
	Pwd    string
	Name   string
	Port   int
}

type code struct {
	Code     int
	ExpireAt time.Time
}

var emailSubject = "拾蜜注册验证码"
var tmpl *template.Template
var codes map[string]code

func EmailPrepare() {
	file, err := os.Open("./conf/email.json")
	if err != nil {
		log.Fatal("EmailSender: ", err)
	}
	var cred Credential
	tmp, _ := io.ReadAll(file)
	json.Unmarshal(tmp, &cred)
	dial = gomail.NewDialer(cred.Server, cred.Port, cred.Name, cred.Pwd)
	dial.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Email template preparation
	tmpl, err = template.ParseFiles("./blob/template.html")
	go maid()
}

// maid Delete expired code
func maid() {
	now := time.Now()
	for k, c := range codes {
		if c.ExpireAt.Before(now) {
			delete(codes, k)
		}
	}
	time.Sleep(time.Duration(time.Minute.Nanoseconds() * 5))
	go maid()
	return
}

func generateCode(email string) int {
	rand.Seed(time.Now().Unix())

	newcode := code{
		Code:     rand.Int() % 100000,
		ExpireAt: time.Now().Add(time.Minute * 5),
	}

	codes[email] = newcode
	return newcode.Code
}

func SendCode(target string) {
	email := gomail.NewMessage()
	email.SetHeader("From", "MuxiZyyyyyS@163.com")
	email.SetHeader("To", target)
	email.SetHeader("Subject", emailSubject)
	var body bytes.Buffer
	tmpl.Execute(&body, generateCode(target))
	email.SetBody("text/html", body.String())
	dial.DialAndSend(email)
}

func CodeIsValid(email string, code int) bool {
	return code == codes[email].Code && codes[email].ExpireAt.After(time.Now())
}
