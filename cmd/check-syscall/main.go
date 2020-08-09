package main

import (
	"time"

	"github.com/ddddddO/work/syscaller/tcpsocket"
)

const (
	duration = 1 * time.Second
)

func main() {
	interval := time.NewTicker(duration)
	end := time.NewTicker(duration * 5)

	go tcpsocket.RunServer()
	time.Sleep(duration * 3)

END:
	for {
		select {
		case <-interval.C:
			tcpsocket.RunClient()
			//sc := file.Gen()
			//syscaller.Run(sc)
		case <-end.C:
			break END
		}
	}
}
