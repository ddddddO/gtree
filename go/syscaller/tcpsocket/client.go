package tcpsocket

import (
	"log"
	"net"
)

func RunClient() {
	conn, err := net.Dial("tcp", serverAddr+serverPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("send from client!"))
	if err != nil {
		log.Fatal(err)
	}

	bufFromServer := make([]byte, 1024)
	_, err = conn.Read(bufFromServer)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(bufFromServer))
}
