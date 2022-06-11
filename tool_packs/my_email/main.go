package main

import (
	"log"
	"net/smtp"
	"github.com/quixote-liu/email"
)

func main() {
	e := email.Email{
		From:     "lcs.shun@foxmail.com",
		To:       []string{"1351169665@qq.com"},
		Subject:  "email testing",
		Comments: "this is comments",
		Addr:     "smtp.qq.com:25",
	}
	auth := smtp.PlainAuth("", "lcs.shun@foxmail.com", "ikpfnntluodtbadh", "smtp.qq.com")

	text := `<h1>testing email !</h1>`
	e.SetAuth(auth).WriteHTML([]byte(text))
	e.WriteText([]byte("heihei, this is text"))

	if err := e.AttachFile("./shenghuo.txt"); err != nil {
		log.Fatal(err)
	}
	if err := e.AttachFile("./html.html"); err != nil {
		log.Fatal(err)
	}


	err := e.Send()
	if err != nil {
		log.Fatal(err)
	}

	e.Reset()
}
