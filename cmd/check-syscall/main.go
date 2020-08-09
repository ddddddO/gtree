package main

import (
	"flag"
	"time"

	"github.com/ddddddO/work/syscaller"
	"github.com/ddddddO/work/syscaller/file"
	"github.com/ddddddO/work/syscaller/tcpsocket"
)

const (
	duration = 1 * time.Second
)

var (
	syscallerType = "file"
)

func main() {
	flag.StringVar(&syscallerType, "s", syscallerType, "'file' or 'tcp'")
	flag.Parse()

	interval := time.NewTicker(duration)
	end := time.NewTicker(duration * 5)

	switch syscallerType {
	case "tcp":
		go tcpsocket.RunServer()
		time.Sleep(duration * 2)
	}

END:
	for {
		select {
		case <-interval.C:
			switch syscallerType {
			case "tcp":
				tcpsocket.RunClient()
			case "file":
				f := file.Gen()
				syscaller.Run(f)
			}
		case <-end.C:
			break END
		}
	}
}
