package main

import (
	"fmt"
	"sync"
	"time"
)

// NOTE:配列要素増やして速度比較したい。で、onMutexと普通のfor文との速度比べる
// 予想：早　offMutex > for文 > onMutex　遅
func main() {
	alphs := []string{"a", "b", "c", "d", "e"}

	onMutex(alphs)

	time.Sleep(1 * time.Second)

	offMutex(alphs)

	// 全goroutinの実行終了を待つため
	time.Sleep(2 * time.Second)
}

func onMutex(as []string) {
	fmt.Println("-on Mutex-")

	mu := sync.Mutex{}
	for _, a := range as {
		mu.Lock()

		go func(s string) {
			defer mu.Unlock()
			fmt.Println(s)
		}(a)
	}
}

func offMutex(as []string) {
	fmt.Println("-off Mutex-")

	for _, a := range as {
		go func(s string) {
			fmt.Println(s)
		}(a)
	}
}
