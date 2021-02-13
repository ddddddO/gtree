package main

import (
	"context"
	"log"
	"time"
	"sync"
)

func main() {
	log.Println("start")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := run1(ctx); err != nil {
			log.Printf("run1 err: %v\n", err)
			cancel() // ここをコメントアウトで、run1のcontextのタイムアウトによって、親のcontextがキャンセルされないため、run2は最後まで実行される。
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := run2(ctx); err != nil {
			log.Printf("run2 err: %v\n", err)
		}
	}()
	wg.Wait()

	log.Println("end")
}

func run1(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3 * time.Second)
	defer cancel()
	log.Println("in run1")

	go run1child(ctx)

	return exec(ctx, 1)
}

func run2(ctx context.Context) error {
	log.Println("in run2")

	go run2child(ctx)

	return exec(ctx, 2)
}

func run1child(ctx context.Context) {
	log.Println("in run1child")
	select {
	case <- ctx.Done():
		log.Printf("ctx done! in run1child: %v", ctx.Err())
	case <- time.After(5 * time.Second):
		log.Println("end run1child")
	}
}

func run2child(ctx context.Context) {
	log.Println("in run2child")
	time.Sleep(5 * time.Second)
	log.Println("end run2child")
}

func exec(ctx context.Context, num int) error {
	log.Printf("in exec: %d\n", num)
	select {
	case <-ctx.Done():
		log.Printf("ctx done!: %d\n", num)
		return ctx.Err()
	case <-time.After(7 * time.Second):
		return nil
	}
	return nil
}