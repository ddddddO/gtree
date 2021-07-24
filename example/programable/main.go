package main

import (
	"fmt"
	"os"

	"github.com/ddddddO/gtree"
)

func main() {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	root.Add("child 4")
	if err := gtree.ExecuteProgrammably(root, os.Stdout); err != nil {
		panic(err)
	}

	// root
	// ├── child 1
	// │   └── child 2
	// │       └── child 3
	// └── child 4

	fmt.Println("-----")

	// https://ja.wikipedia.org/wiki/%E3%82%B5%E3%83%AB%E7%9B%AE
	primate := gtree.NewRoot("Primate")

	strepsirrhini := primate.Add("Strepsirrhini")
	haplorrhini := primate.Add("Haplorrhini")

	_ = strepsirrhini.Add("Lemuriformes")
	_ = strepsirrhini.Add("Lorisiformes")

	_ = haplorrhini.Add("Tarsiiformes")
	_ = haplorrhini.Add("Simiiformes")
	if err := gtree.ExecuteProgrammably(primate, os.Stdout); err != nil {
		panic(err)
	}

	// Primate
	// ├── Strepsirrhini
	// │   ├── Lemuriformes
	// │   └── Lorisiformes
	// └── Haplorrhini
	//     ├── Tarsiiformes
	//     └── Simiiformes
}
