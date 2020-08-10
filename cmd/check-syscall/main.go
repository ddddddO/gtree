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
	count         = 1
)

func main() {
	flag.StringVar(&syscallerType, "s", syscallerType, "'file' or 'tcp'")
	flag.IntVar(&count, "c", count, "exec count")
	flag.Parse()

	switch syscallerType {
	case "tcp":
		go tcpsocket.RunServer()
		time.Sleep(duration * 2)
	}

	for i := 0; i < count; i++ {
		switch syscallerType {
		case "tcp":
			tcpsocket.RunClient()
		case "file":
			f := file.Gen()
			syscaller.Run(f)
		}
	}
}
