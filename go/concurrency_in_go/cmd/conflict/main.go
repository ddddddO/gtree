package main

import (
	"log"
	//"time"
)

func main() {
	data := 0

	go func() {
		data++
	}()

	//time.Sleep(1 * time.Second)

	if data == 0 {
		log.Printf("data is %d", data)
	}
	log.Println("end")
}