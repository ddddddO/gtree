package main

import (
	"log"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	printValue := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(1 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		log.Printf("v1 + v2 = %d", v1.value+v2.value)
	}

	a := value{value: 22}
	b := value{value: 33}

	wg.Add(2)
	go printValue(&a, &b)
	go printValue(&b, &a)
	wg.Wait()
}
