package main

import (
	"fmt"
	"time"
)

func runSel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go gor1(ch1)
	go gor2(ch2)

	for {
		select {
		case s1 := <-ch1:
			fmt.Println(s1)
		case s2 := <-ch2:
			fmt.Println(s2)
		}
	}

}

func gor1(ch chan string) {
	for {
		ch <- "gor1"
		time.Sleep(1 * time.Second)
	}
}

func gor2(ch chan string) {
	for {
		ch <- "gor2"
		time.Sleep(2 * time.Second)
	}
}
