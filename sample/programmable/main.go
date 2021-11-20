package main

import (
	"os"

	"github.com/ddddddO/gtree/v6"
)

func main() {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	root.Add("child 5")
	root.Add("child 1").Add("child 2").Add("child 4")
	if err := gtree.ExecuteProgrammably(os.Stdout, root); err != nil {
		panic(err)
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
	if err := gtree.ExecuteProgrammably(os.Stdout, primate); err != nil {
		panic(err)
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
	if err := gtree.ExecuteProgrammably(os.Stdout, primate, gtree.EncodeJSON()); err != nil {
		panic(err)
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
	if err := gtree.ExecuteProgrammably(os.Stdout, primate, gtree.EncodeYAML()); err != nil {
		panic(err)
	}
	// Output:
	// value: Primate
	// children:
	// - value: Strepsirrhini
	//   children:
	//   - value: Lemuriformes
	// 	children:
	// 	- value: Lemuroidea
	// 	  children:
	// 	  - value: Cheirogaleidae
	// 		children: []
	// 	  - value: Indriidae
	// 		children: []
	// 	  - value: Lemuridae
	// 		children: []
	// 	  - value: Lepilemuridae
	// 		children: []
	// 	- value: Daubentonioidea
	// 	  children:
	// 	  - value: Daubentoniidae
	// 		children: []
	//   - value: Lorisiformes
	// 	children:
	// 	- value: Galagidae
	// 	  children: []
	// 	- value: Lorisidae
	// 	  children: []
	// - value: Haplorrhini
	//   children:
	//   - value: Tarsiiformes
	// 	children:
	// 	- value: Tarsiidae
	// 	  children: []
	//   - value: Simiiformes
	// 	children:
	// 	- value: Platyrrhini
	// 	  children:
	// 	  - value: Ceboidea
	// 		children:
	// 		- value: Atelidae
	// 		  children: []
	// 		- value: Cebidae
	// 		  children: []
	// 	  - value: Pithecioidea
	// 		children:
	// 		- value: Pitheciidae
	// 		  children: []
	// 	- value: Catarrhini
	// 	  children:
	// 	  - value: Cercopithecoidea
	// 		children:
	// 		- value: Cercopithecidae
	// 		  children: []
	// 	  - value: Hominoidea
	// 		children:
	// 		- value: Hylobatidae
	// 		  children: []
	// 		- value: Hominidae
	// 		  children: []

	// output toml
	if err := gtree.ExecuteProgrammably(os.Stdout, primate, gtree.EncodeTOML()); err != nil {
		panic(err)
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
