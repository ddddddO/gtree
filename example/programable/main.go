package main

import (
	"os"

	"github.com/ddddddO/gtree"
)

func main() {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	root.Add("child 5")
	root.Add("child 1").Add("child 2").Add("child 4")
	if err := gtree.ExecuteProgrammably(os.Stdout, root); err != nil {
		panic(err)
	}
	// root
	// ├── child 1
	// │   └── child 2
	// │       ├── child 3
	// │       └── child 4
	// └── child 5

	primate := preparePrimate()
	if err := gtree.ExecuteProgrammably(os.Stdout, primate); err != nil {
		panic(err)
	}
	// Primate
	// ├── Strepsirrhini
	// │   ├── Lemuriformes
	// │   │   ├── Lemuroidea
	// │   │   │   ├── Cheirogaleidae
	// │   │   │   ├── Indriidae
	// │   │   │   ├── Lemuridae
	// │   │   │   └── Lepilemuridae
	// │   │   └── Daubentonioidea
	// │   │       └── Daubentoniidae
	// │   └── Lorisiformes
	// │       ├── Galagidae
	// │       └── Lorisidae
	// └── Haplorrhini
	//     ├── Tarsiiformes
	//     │   └── Tarsiidae
	//     └── Simiiformes
	//         ├── Platyrrhini
	//         │   ├── Ceboidea
	//         │   │   ├── Atelidae
	//         │   │   └── Cebidae
	//         │   └── Pithecioidea
	//         │       └── Pitheciidae
	//         └── Catarrhini
	//             ├── Cercopithecoidea
	//             │   └── Cercopithecidae
	//             └── Hominoidea
	//                 ├── Hylobatidae
	//                 └── Hominidae
}

func preparePrimate() *gtree.Node {
	// https://ja.wikipedia.org/wiki/%E3%82%B5%E3%83%AB%E7%9B%AE
	primate := gtree.NewRoot("Primate")
	strepsirrhini := primate.Add("Strepsirrhini")
	haplorrhini := primate.Add("Haplorrhini")
	lemuriformes := strepsirrhini.Add("Lemuriformes")
	lorisiformes := strepsirrhini.Add("Lorisiformes")

	lemuroidea := lemuriformes.Add("Lemuroidea")
	lemuroidea.Add("Cheirogaleidae")
	lemuroidea.Add("Indriidae")
	lemuroidea.Add("Lemuridae")
	lemuroidea.Add("Lepilemuridae")
	lemuriformes.Add("Daubentonioidea").Add("Daubentoniidae")

	lorisiformes.Add("Galagidae")
	lorisiformes.Add("Lorisidae")

	haplorrhini.Add("Tarsiiformes").Add("Tarsiidae")
	simiiformes := haplorrhini.Add("Simiiformes")

	platyrrhini := simiiformes.Add("Platyrrhini")
	ceboidea := platyrrhini.Add("Ceboidea")
	ceboidea.Add("Atelidae")
	ceboidea.Add("Cebidae")
	platyrrhini.Add("Pithecioidea").Add("Pitheciidae")

	catarrhini := simiiformes.Add("Catarrhini")
	catarrhini.Add("Cercopithecoidea").Add("Cercopithecidae")
	hominoidea := catarrhini.Add("Hominoidea")
	hominoidea.Add("Hylobatidae")
	hominoidea.Add("Hominidae")

	return primate
}
