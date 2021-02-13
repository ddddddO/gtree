package tcpconn

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type Client struct {
	Conn    net.Conn
	stdInSc *bufio.Scanner
	bw      *bufio.Writer
	br      *bufio.Reader
}


// client生成
func NewClient(target, port string) *Client {
	c := connect(target, port)

	return &Client{
		Conn:    c,
		stdInSc: bufio.NewScanner(os.Stdin), // client側の標準入力を取得
		bw:      bufio.NewWriter(c),
		br:      bufio.NewReader(c),
	}
}

// serverとconnection確立
func connect(target, port string) net.Conn {
	log.Printf("start connect to %s:%s\n", target, port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", target, port))
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func (c *Client) InteractToServer() {
	// serverから接続成功メッセージ取得
	line, err := c.br.ReadString('\n')
	if err == io.EOF {
		log.Fatal(err)
	}
	log.Print(line)

	// serverと対話
REMESSAGE:
	fmt.Println("please input message:")
	text := ""
	if c.stdInSc.Scan() {
		text = c.stdInSc.Text()
	}

	c.bw.WriteString(text + "\n")
	c.bw.Flush()

	for {
		sMsgBuf := make([]byte, 1024) // 予めバッファ確保しないとダメっぽい
		_, err := c.br.Read(sMsgBuf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(sMsgBuf))

		goto REMESSAGE
	}
}
