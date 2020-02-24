package main

import (
	"flag"
	"log"

	"github.com/ddddddO/work/dot"
)

func main() {
	var isDigraph bool
	flag.BoolVar(&isDigraph, "digraph", false, "有向グラフの場合は'digraph'を指定する")
	flag.Parse()

	g := dot.NewGraph(
		"test",
		isDigraph,
		[]string{"A", "B", "C", "D"},
		[][]string{
			[]string{"A", "B"},
			[]string{"B", "C"},
			[]string{"C", "D"},
			[]string{"D", "A"},
		},
	)

	if err := dot.Gen(g); err != nil {
		log.Fatal(err)
	}
}
