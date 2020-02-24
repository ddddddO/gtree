package main

import (
	"flag"
	"log"

	"github.com/ddddddO/work/dot"
)

func main() {
	var (
		isDigraph bool
		cnt       int
	)
	flag.BoolVar(&isDigraph, "digraph", false, "有向グラフの場合は'digraph'を指定する")
	flag.IntVar(&cnt, "cnt", 11, "ランキングの昇順からcnt番目までの記事を対象にする")
	flag.Parse()

	nodes, edges, err := dot.Scrape(cnt)
	if err != nil {
		log.Fatal(err)
	}

	g := dot.NewGraph(
		"test",
		isDigraph,
		nodes,
		edges,
	)

	if err := dot.Gen(g); err != nil {
		log.Fatal(err)
	}
}
