package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ddddddO/work/cmd/embedded/lib"
)

func main() {
	dp := DockerParts{}

	lib.Run(dp)
}

type DockerParts struct {
	lib.Default
}

func (dp DockerParts) Update(vs []interface{}) {
	fmt.Println("DockerParts Update")

	wg := &sync.WaitGroup{}
	for i, _ := range vs {
		wg.Add(1)

		go dp.routine(i, wg)
	}
	wg.Wait()
}

func (dp DockerParts) routine(v interface{}, wg *sync.WaitGroup) {
	fmt.Println(v)
	time.Sleep(1 * time.Second)

	wg.Done()
}
