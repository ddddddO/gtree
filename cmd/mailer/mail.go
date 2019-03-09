package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// NOTE: Senderを作って、その中にos.Stdinを持つ感じにすればよそそう？
// FIXME: あと、実行結果がFrom以外が、Body部に出力される問題

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

func NewMail() *Mail {
	return &Mail{
		From:    getFrom(),
		To:      getTo(),
		Subject: getSubject(),
		Body:    getBody(),
	}
}

func (m *Mail) send(target string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("sendmail %s", target))
	w, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	defer w.Close()

	if err := cmd.Start(); err != nil {
		return err
	}

	io.WriteString(w, fmt.Sprintf("From: %s\n", m.From))
	io.WriteString(w, fmt.Sprintf("To: %s\n", m.To))
	io.WriteString(w, fmt.Sprintf("Subject: %s\n", m.Subject))
	io.WriteString(w, m.Body+"\n")
	io.WriteString(w, ".\n")

	cmd.Wait()

	return nil
}

func getFrom() string {
	fmt.Println("From:")

	input := os.Stdin
	//defer input.Close()

	buff := make([]byte, 1024)
	_, err := input.Read(buff)
	if err != nil {
		panic(err)
	}

	return string(buff)
}

func getTo() string {
	fmt.Println("To:")

	input := os.Stdin
	//defer input.Close()

	buff := make([]byte, 1024)
	_, err := input.Read(buff)
	if err != nil {
		panic(err)
	}

	return string(buff)
}

func getSubject() string {
	fmt.Println("Subject:")

	input := os.Stdin
	//defer input.Close()

	buff := make([]byte, 1024)
	_, err := input.Read(buff)
	if err != nil {
		panic(err)
	}

	return string(buff)
}

func getBody() string {
	fmt.Println("Body:")

	input := os.Stdin
	defer input.Close()

	buff := make([]byte, 1024)
	_, err := input.Read(buff)
	if err != nil {
		panic(err)
	}

	return string(buff)
}
