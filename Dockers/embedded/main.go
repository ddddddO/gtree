package main

import (
	"fmt"
	"time"

	"github.com/ddddddO/work/cmd/embedded/lib"
)

type data struct {
	num   int
	isErr bool
}

func main() {
	dp := DockerParts{}

	lib.Run(dp)
}

type DockerParts struct {
	lib.Default
}

func (dp DockerParts) Update(vs []interface{}) {
	fmt.Println("DockerParts Update")

	ch := make(chan data, len(vs))
	for i, _ := range vs {
		go dp.routine(i, ch)
	}

	for i := 0; i < len(vs); i++ {
		d := <-ch
		if d.isErr {
			fmt.Printf("faild num: %d\n", d.num)
		} else {
			fmt.Printf("suceeded num: %d\n", d.num)
		}
	}
}

func (dp DockerParts) routine(v interface{}, ch chan data) {
	fmt.Println(v)
	time.Sleep(1 * time.Second)

	num := v.(int)
	d := data{
		num:   num,
		isErr: false,
	}

	if num == 2 {
		d.isErr = true
	}

	ch <- d
}
