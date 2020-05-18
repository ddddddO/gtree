package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.Println("start")

	ctx := context.Background()
	ctx1, _ := context.WithTimeout(ctx, time.Second * 5)
	go ctx1F(ctx1)
	go ctx2F(ctx1)

	select {
	case <- ctx1.Done():
		log.Println("ctx1 done:", ctx1.Err())
	}

	//time.Sleep(time.Second * 5) // ctx2Fは実行される
	log.Println("end")
}

func ctx1F(ctx context.Context) {
	time.Sleep(3 * time.Second)
	log.Println("in ctx1")
}

func ctx2F(ctx context.Context) {
	time.Sleep(7 * time.Second)
	log.Println("in ctx2")
}