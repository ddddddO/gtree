package main

import (
	"flag"
	"log"

	tc "github.com/ddddddO/work/tcpconn"
)

var (
	port    string
	handler string
)

var handlers = map[string]tc.Handler{
	"janken": jankenHandler,
	"euc":    convEUCHandler,
}

func main() {
	flag.StringVar(&port, "port", "8887", "server port open")
	flag.StringVar(&handler, "handler", "janken", "handler prefix")
	flag.Parse()

	h, ok := handlers[handler]
	if !ok {
		log.Fatal("not in handlers")
	}

	tc.Run(port, h)
}
