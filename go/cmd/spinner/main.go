package main

import (
	"fmt"
	"time"
)

func main() {
	spinner(time.Millisecond * 100)
}

func spinner(delay time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}
