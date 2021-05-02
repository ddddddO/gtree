package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	configPath = "./config.txt"
)

func main() {
	log.Println("start")

	sig := make(chan os.Signal, 1)
	// 設定ファイル更新後、SIGHUPでプロセスを再起動して設定ファイルを再読み込みしたりする。
	signal.Notify(sig, syscall.SIGHUP)

	ticker := time.NewTicker(2 * time.Second)
	dots := "."
	for {
		select {
		case <-ticker.C:
			fmt.Println(dots)
			dots += "."
		case <-sig:
			log.Println("restart")
			dots = "."
			continue
		}
	}
}
