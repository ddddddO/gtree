package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	src, err := os.Open("./in")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=============")
	fmt.Fprintln(os.Stdout, b)
	fmt.Println("=============")

	dst, err := os.Create("./out")
	if err != nil {
		log.Fatal(err)
	}

	// 97 -> 122
	b[0] = 122

	br := bytes.NewBuffer(b)
	_, err = io.Copy(dst, br)
	if err != nil {
		log.Fatal(err)
	}
}
