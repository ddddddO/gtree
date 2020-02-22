package main

import (
	"flag"
	"log"

	"github.com/ddddddO/work/dot"
)

func main() {
	var isDigraph bool
	flag.BoolVar(&isDigraph, "digraph", false, "有効グラフの場合は'digraph'を指定する")
	flag.Parse()

	generator := dot.NewDotGenerator(
		isDigraph,
		"dumy/path/xxx.dot",
	)

	err := dot.Run(generator)
	if err != nil {
		log.Fatal(err)
	}
}
