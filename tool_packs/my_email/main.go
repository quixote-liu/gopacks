package main

import (
	"bytes"
	"log"
	"net/smtp"
	"os"

	"github.com/quixote-liu/email"
)

func main() {
	e := email.Email{
		From:     "lcs.shun@foxmail.com",
		To:       []string{"1351169665@qq.com", "lcs_shun@163.com"},
		Subject:  "email testing",
		Comments: "this is comments",
		Addr:     "smtp.qq.com:25",
	}
	auth := smtp.PlainAuth("", "lcs.shun@foxmail.com", "ikpfnntluodtbadh", "smtp.qq.com")

	e.SetAuth(auth)
	e.WriteText(text2())

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

func text() []byte {
	f, err := os.Open("./photo.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bytes.NewBufferString("hello, world")
	if _, err := buf.ReadFrom(f); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func text2() []byte {
	return []byte("hello, world")
}
