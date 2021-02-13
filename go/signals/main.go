package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// signalとは？https://techblog.istyle.co.jp/archives/425
// 主なsignal https://qiita.com/bluesDD/items/43a255bcf0dee6798967#%E3%82%B7%E3%82%B0%E3%83%8A%E3%83%AB%E3%81%AE%E7%A8%AE%E9%A1%9E
func main() {
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGINT)

	for {
		select {
		case sig := <-signalCh:
			switch sig {
			case syscall.SIGINT:
				log.Println("end")
				os.Exit(1)
			}
		}
	}
}
