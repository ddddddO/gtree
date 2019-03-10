package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v  map[string]int
	mu sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.v[key]
}

func runMu() {
	c := Counter{v: make(map[string]int)}

	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			c.Inc("Key")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(c.Value("Key"))

}
