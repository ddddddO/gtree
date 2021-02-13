package main

import (
	"fmt"
	"time"
)

func runTic() {
	tic := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

OUT:
	for {
		select {
		case <-tic:
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("Boom!")
			//break // NOTE: 抜けられない
			break OUT // NOTE: ラベル + breakで、for抜ける
			//return
		default:
			fmt.Println(".")
			time.Sleep(50 * time.Millisecond)
		}
	}

	fmt.Println("break for")

}
