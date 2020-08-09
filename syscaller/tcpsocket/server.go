package tcpsocket

import (
	"log"
	"net"
)

const (
	serverAddr = "localhost"
	serverPort = ":8888"
)

func RunServer() {
	ln, err := net.Listen("tcp", serverAddr+serverPort)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(conn net.Conn) {
			//defer conn.Close()

			bufFromClient := make([]byte, 1024)
			_, err := conn.Read(bufFromClient)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("debug...: %s\n", string(bufFromClient))

			bufFromClient = append(bufFromClient, []byte(": from server!")...)
			_, err = conn.Write(bufFromClient)
			if err != nil {
				log.Fatal(err)
			}
		}(conn)
	}
}
