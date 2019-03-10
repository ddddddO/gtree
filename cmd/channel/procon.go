package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

// Producer/Consumer
func run() {
	// --tracer
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()
	// --tracer

	ch := make(chan string)
	wg := &sync.WaitGroup{}
	ss := []string{"A", "B", "C", "D", "X"}

	for _, s := range ss {
		wg.Add(1)
		go producer(s, ch)
	}

	go consumer(ch, wg)
	wg.Wait()
	close(ch) // NOTE: consumer内のforでch待ち受け続けてしまったまま、run()が終わってしまうから

}

func producer(s string, ch chan string) {
	// 何らかの処理
	time.Sleep(1 * time.Second)

	ch <- s + "!!"
}

func consumer(ch chan string, wg *sync.WaitGroup) {
	for c := range ch {
		fmt.Println(c)
		wg.Done()
	}
	fmt.Println("-consumer end-")
}
