package tcpconn

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Handler func(s string) string

// 参考：https://qiita.com/Imyslx/items/2c11fb75a8ce1e6aef8b
func Run(port string, h Handler) {
	log.Printf("sever start. open port:%s\n", port)
	// ポートを開く
	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", port))
	if err != nil {
		log.Fatal(err)
	}

	var n int

	for {
		// clientから接続待ち
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		n++

		// 接続後処理
		go func(conn net.Conn, n int) {
			bw := bufio.NewWriter(conn)
			br := bufio.NewReader(conn)

			bw.WriteString("connect server!\n")
			bw.Flush()
			log.Printf("client %d\n", n)

			// clientと対話
			for {
				msg, err := br.ReadString('\n')
				if err == io.EOF {
					continue
				}
				if err != nil {
					log.Fatal(err)
				}

				bw.WriteString("response from server:\n")
				bw.WriteString(h(msg))
				bw.Flush()
			}
		}(conn, n)
	}
}
