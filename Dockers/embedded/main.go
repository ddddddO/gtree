package main

import (
	"fmt"

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

	for i, _ := range vs {
		fmt.Println(i)
	}
}
