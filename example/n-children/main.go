package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ddddddO/gtree"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 1000, "number of children")
	flag.Parse()

	fmt.Printf("=== %d children ===\n", n)

	root := gtree.NewRoot("Parent")
	addChildren(root, n)
	if err := gtree.OutputFromRoot(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addChildren(root *gtree.Node, n int) {
	for i := range n {
		child := root.Add(fmt.Sprintf("child %d", i))
		grandchild1 := child.Add("aaaaaaaaaa")
		grandchild1.Add("ccccccccccccccccc")
		grandchild2 := child.Add("bbbbbbbbbbbbbbb")
		g := grandchild2.Add("ddddddddddddddddddd")
		g.Add("eeeeeeeeeeeeeeeeeeeeeeeeee")
	}
}
