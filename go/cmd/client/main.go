package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	tc "github.com/ddddddO/work/tcpconn"
)

var (
	target string
	port   string
)

// 参考：https://qiita.com/TubAnri/items/019f8d19b91f32c878cf
func main() {
	flag.StringVar(&target, "target", "127.0.0.1", "target server ip")
	flag.StringVar(&port, "port", "8887", "target server port")
	flag.Parse()

	client := tc.NewClient(target, port)
	go client.InteractToServer()

	// ctrl c signal用
	quit := make(chan os.Signal)
	// 受け取るシグナル設定
	signal.Notify(quit, os.Interrupt)
	// シグナル受け取るまで以降の処理は実行されない
	<-quit

	client.Conn.Close()
	log.Println("connection close...")

}
