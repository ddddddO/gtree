package main

import (
	"fmt"
)

// fan-out fan-in
func runFan() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producerF(first)
	go times2(first, second)
	go times3(second, third)

	for t := range third {
		fmt.Println(t)
	}
}

func producerF(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func times2(first <-chan int /*受信用*/, second chan<- int /*送信用*/) {
	defer close(second)
	for f := range first {
		second <- f * 2
	}
}

func times3(second, third chan int) {
	defer close(third)
	for s := range second {
		third <- s * 3
	}
}
