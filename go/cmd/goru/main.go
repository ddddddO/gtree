package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	s := []string{"a", "b", "c", "d", "e"}
	ch := make(chan string)

	log.Println("start")
	for i := 0; i < len(s); i++ {
		go func(i int) {
			time.Sleep(time.Duration(i) * 2 * time.Second)
			ch <- s[i]
		}(i)
	}

	f := func() {
		for {
			select {
			case c := <-ch:
				fmt.Printf("%s desu\n", c)
			case <-time.After(3 * time.Second):
				fmt.Println("afafa")
				go func() {
					ch <- "z"
				}()
			}
		}
	}

	// 無限ループ使わずとも
	for {
		f()
	}

	log.Println("end")
}
