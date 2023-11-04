package main

import (
	"fmt"
	"os"

	"github.com/ddddddO/gtree"
)

func main() {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	root.Add("child 5")
	root.Add("child 1").Add("child 2").Add("child 4")
	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// root
	// ├── child 1
	// │   └── child 2
	// │       ├── child 3
	// │       └── child 4
	// └── child 5

	primate := preparePrimate()
	// default branch format.
	if err := gtree.OutputProgrammably(os.Stdout, primate); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
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

	// output json
	if err := gtree.OutputProgrammably(os.Stdout, primate, gtree.WithEncodeJSON()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output(using 'jq'):
	// {
	// 	"value": "Primate",
	// 	"children": [
	// 	  {
	// 		"value": "Strepsirrhini",
	// 		"children": [
	// 		  {
	// 			"value": "Lemuriformes",
	// 			"children": [
	// 			  {
	// 				"value": "Lemuroidea",
	// 				"children": [
	// 				  {
	// 					"value": "Cheirogaleidae",
	// 					"children": null
	// 				  },
	// 				  {
	// 					"value": "Indriidae",
	// 					"children": null
	// 				  },
	// 				  {
	// 					"value": "Lemuridae",
	// 					"children": null
	// 				  },
	// 				  {
	// 					"value": "Lepilemuridae",
	// 					"children": null
	// 				  }
	// 				]
	// 			  },
	// 			  {
	// 				"value": "Daubentonioidea",
	// 				"children": [
	// 				  {
	// 					"value": "Daubentoniidae",
	// 					"children": null
	// 				  }
	// 				]
	// 			  }
	// 			]
	// 		  },
	// 		  {
	// 			"value": "Lorisiformes",
	// 			"children": [
	// 			  {
	// 				"value": "Galagidae",
	// 				"children": null
	// 			  },
	// 			  {
	// 				"value": "Lorisidae",
	// 				"children": null
	// 			  }
	// 			]
	// 		  }
	// 		]
	// 	  },
	// 	  {
	// 		"value": "Haplorrhini",
	// 		"children": [
	// 		  {
	// 			"value": "Tarsiiformes",
	// 			"children": [
	// 			  {
	// 				"value": "Tarsiidae",
	// 				"children": null
	// 			  }
	// 			]
	// 		  },
	// 		  {
	// 			"value": "Simiiformes",
	// 			"children": [
	// 			  {
	// 				"value": "Platyrrhini",
	// 				"children": [
	// 				  {
	// 					"value": "Ceboidea",
	// 					"children": [
	// 					  {
	// 						"value": "Atelidae",
	// 						"children": null
	// 					  },
	// 					  {
	// 						"value": "Cebidae",
	// 						"children": null
	// 					  }
	// 					]
	// 				  },
	// 				  {
	// 					"value": "Pithecioidea",
	// 					"children": [
	// 					  {
	// 						"value": "Pitheciidae",
	// 						"children": null
	// 					  }
	// 					]
	// 				  }
	// 				]
	// 			  },
	// 			  {
	// 				"value": "Catarrhini",
	// 				"children": [
	// 				  {
	// 					"value": "Cercopithecoidea",
	// 					"children": [
	// 					  {
	// 						"value": "Cercopithecidae",
	// 						"children": null
	// 					  }
	// 					]
	// 				  },
	// 				  {
	// 					"value": "Hominoidea",
	// 					"children": [
	// 					  {
	// 						"value": "Hylobatidae",
	// 						"children": null
	// 					  },
	// 					  {
	// 						"value": "Hominidae",
	// 						"children": null
	// 					  }
	// 					]
	// 				  }
	// 				]
	// 			  }
	// 			]
	// 		  }
	// 		]
	// 	  }
	// 	]
	// }

	// output yaml
	if err := gtree.OutputProgrammably(os.Stdout, primate, gtree.WithEncodeYAML()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// value: Primate
	// children:
	// - value: Strepsirrhini
	//   children:
	//   - value: Lemuriformes
	//     children:
	//     - value: Lemuroidea
	//       children:
	//       - value: Cheirogaleidae
	//         children: []
	//       - value: Indriidae
	//         children: []
	//       - value: Lemuridae
	//         children: []
	//       - value: Lepilemuridae
	//         children: []
	//     - value: Daubentonioidea
	//       children:
	//       - value: Daubentoniidae
	//         children: []
	//   - value: Lorisiformes
	//     children:
	//     - value: Galagidae
	//       children: []
	//     - value: Lorisidae
	//       children: []
	// - value: Haplorrhini
	//   children:
	//   - value: Tarsiiformes
	//     children:
	//     - value: Tarsiidae
	//       children: []
	//   - value: Simiiformes
	//     children:
	//     - value: Platyrrhini
	//       children:
	//       - value: Ceboidea
	//         children:
	//         - value: Atelidae
	//           children: []
	//         - value: Cebidae
	//           children: []
	//       - value: Pithecioidea
	//         children:
	//         - value: Pitheciidae
	//           children: []
	//     - value: Catarrhini
	//       children:
	//       - value: Cercopithecoidea
	//         children:
	//         - value: Cercopithecidae
	//           children: []
	//       - value: Hominoidea
	//         children:
	//         - value: Hylobatidae
	//           children: []
	//         - value: Hominidae
	//           children: []

	// output toml
	if err := gtree.OutputProgrammably(os.Stdout, primate, gtree.WithEncodeTOML()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// value = 'Primate'
	// [[children]]
	// value = 'Strepsirrhini'
	// [[children.children]]
	// value = 'Lemuriformes'
	// [[children.children.children]]
	// value = 'Lemuroidea'
	// [[children.children.children.children]]
	// value = 'Cheirogaleidae'
	// children = []
	// [[children.children.children.children]]
	// value = 'Indriidae'
	// children = []
	// [[children.children.children.children]]
	// value = 'Lemuridae'
	// children = []
	// [[children.children.children.children]]
	// value = 'Lepilemuridae'
	// children = []
	//
	// [[children.children.children]]
	// value = 'Daubentonioidea'
	// [[children.children.children.children]]
	// value = 'Daubentoniidae'
	// children = []
	//
	//
	// [[children.children]]
	// value = 'Lorisiformes'
	// [[children.children.children]]
	// value = 'Galagidae'
	// children = []
	// [[children.children.children]]
	// value = 'Lorisidae'
	// children = []
	//
	//
	// [[children]]
	// value = 'Haplorrhini'
	// [[children.children]]
	// value = 'Tarsiiformes'
	// [[children.children.children]]
	// value = 'Tarsiidae'
	// children = []
	//
	// [[children.children]]
	// value = 'Simiiformes'
	// [[children.children.children]]
	// value = 'Platyrrhini'
	// [[children.children.children.children]]
	// value = 'Ceboidea'
	// [[children.children.children.children.children]]
	// value = 'Atelidae'
	// children = []
	// [[children.children.children.children.children]]
	// value = 'Cebidae'
	// children = []
	//
	// [[children.children.children.children]]
	// value = 'Pithecioidea'
	// [[children.children.children.children.children]]
	// value = 'Pitheciidae'
	// children = []
	//
	//
	// [[children.children.children]]
	// value = 'Catarrhini'
	// [[children.children.children.children]]
	// value = 'Cercopithecoidea'
	// [[children.children.children.children.children]]
	// value = 'Cercopithecidae'
	// children = []
	//
	// [[children.children.children.children]]
	// value = 'Hominoidea'
	// [[children.children.children.children.children]]
	// value = 'Hylobatidae'
	// children = []
	// [[children.children.children.children.children]]
	// value = 'Hominidae'
	// children = []
	//
	//
	//
	//
	//

	sampleWalker()

	// make directories.
	if err := gtree.MkdirProgrammably(primate); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output(using Linux 'tree' command):
	// 22:20:43 > tree Primate/
	// Primate/
	// ├── Haplorrhini
	// │   ├── Simiiformes
	// │   │   ├── Catarrhini
	// │   │   │   ├── Cercopithecoidea
	// │   │   │   │   └── Cercopithecidae
	// │   │   │   └── Hominoidea
	// │   │   │       ├── Hominidae
	// │   │   │       └── Hylobatidae
	// │   │   └── Platyrrhini
	// │   │       ├── Ceboidea
	// │   │       │   ├── Atelidae
	// │   │       │   └── Cebidae
	// │   │       └── Pithecioidea
	// │   │           └── Pitheciidae
	// │   └── Tarsiiformes
	// │       └── Tarsiidae
	// └── Strepsirrhini
	// 	├── Lemuriformes
	// 	│   ├── Daubentonioidea
	// 	│   │   └── Daubentoniidae
	// 	│   └── Lemuroidea
	// 	│       ├── Cheirogaleidae
	// 	│       ├── Indriidae
	// 	│       ├── Lemuridae
	// 	│       └── Lepilemuridae
	// 	└── Lorisiformes
	// 		├── Galagidae
	// 		└── Lorisidae
	//
	// 28 directories, 0 files

	gtreeDir := gtree.NewRoot("gtree")
	gtreeDir.Add("cmd").Add("gtree").Add("main.go")
	gtreeDir.Add("Makefile")
	testdataDir := gtreeDir.Add("testdata")
	testdataDir.Add("sample1.md")
	testdataDir.Add("sample2.md")
	gtreeDir.Add("tree.go")

	// make directories and files with specific extensions.
	if err := gtree.MkdirProgrammably(
		gtreeDir,
		gtree.WithFileExtensions([]string{".go", ".md", "Makefile"}),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output(using Linux 'tree' command):
	// 09:44:50 > tree gtree/
	// gtree/
	// ├── cmd
	// │   └── gtree
	// │       └── main.go
	// ├── Makefile
	// ├── testdata
	// │   ├── sample1.md
	// │   └── sample2.md
	// └── tree.go
	//
	// 3 directories, 5 files
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

func sampleWalker() {
	fmt.Println("\nWalker Sample...")
	fmt.Println()

	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	root.Add("child 5")
	root.Add("child 1").Add("child 2").Add("child 4")

	callback := func(wn *gtree.WalkerNode) error {
		fmt.Println(wn.Row())
		return nil
	}

	if err := gtree.WalkProgrammably(root, callback); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// root
	// ├── child 1
	// │   └── child 2
	// │       ├── child 3
	// │       └── child 4
	// └── child 5

	fmt.Println("----------------------------------------")

	callback2 := func(wn *gtree.WalkerNode) error {
		fmt.Println("WalkerNode's methods called...")
		fmt.Printf("\tName   : %s\n", wn.Name())
		fmt.Printf("\tBranch : %s\n", wn.Branch())
		fmt.Printf("\tRow    : %s\n", wn.Row())
		fmt.Printf("\tLevel  : %d\n", wn.Level())
		fmt.Printf("\tPath   : %s\n", wn.Path())
		return nil
	}

	if err := gtree.WalkProgrammably(root, callback2); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// WalkerNode's methods called...
	// 	Name   : root
	// 	Branch :
	// 	Row    : root
	// 	Level  : 1
	// 	Path   : root
	// WalkerNode's methods called...
	// 	Name   : child 1
	// 	Branch : ├──
	// 	Row    : ├── child 1
	// 	Level  : 2
	// 	Path   : root/child 1
	// WalkerNode's methods called...
	// 	Name   : child 2
	// 	Branch : │   └──
	// 	Row    : │   └── child 2
	// 	Level  : 3
	// 	Path   : root/child 1/child 2
	// WalkerNode's methods called...
	// 	Name   : child 3
	// 	Branch : │       ├──
	// 	Row    : │       ├── child 3
	// 	Level  : 4
	// 	Path   : root/child 1/child 2/child 3
	// WalkerNode's methods called...
	// 	Name   : child 4
	// 	Branch : │       └──
	// 	Row    : │       └── child 4
	// 	Level  : 4
	// 	Path   : root/child 1/child 2/child 4
	// WalkerNode's methods called...
	// 	Name   : child 5
	// 	Branch : └──
	// 	Row    : └── child 5
	// 	Level  : 2
	// 	Path   : root/child 5

	fmt.Println("\nWalker Sample...end")
	fmt.Println()
}
