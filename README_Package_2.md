## Package(2) / generate a tree programmatically

### Installation
```console
go get github.com/ddddddO/gtree/v6
```

### Usage

```go
package main

import (
	"os"

	"github.com/ddddddO/gtree/v6"
)

func main() {
	var root *gtree.Node = gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *gtree.Node = root.Add("child 1").Add("child 2").Add("child 4")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	root.Add("child 8")
	// you can customize branch format.
	if err := gtree.ExecuteProgrammably(os.Stdout, root,
		gtree.BranchFormatIntermedialNode("+--", ":   "),
		gtree.BranchFormatLastNode("+--", "    "),
	); err != nil {
		panic(err)
	}
	// Output:
	// root
	// +-- child 1
	// :   +-- child 2
	// :       +-- child 3
	// :       +-- child 4
	// :           +-- child 5
	// :           +-- child 6
	// :               +-- child 7
	// +-- child 8

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
}

func preparePrimate() *gtree.Node {
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

	platyrrhini := haplorrhini.Add("Platyrrhini")
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

```

---

- The program below converts the result of `find` into a tree.
```go
package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/ddddddO/gtree/v6"
)

// cd github.com/ddddddO/gtree
// find . -type d -name .git -prune -o -type f -print | go run sample/find_pipe_programmable-gtree/main.go
func main() {
	var (
		root *gtree.Node
		node *gtree.Node
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		splited := strings.Split(line, "/")

		for i, s := range splited {
			if i == 0 {
				if root == nil {
					root = gtree.NewRoot(s)
					node = root
				}
				continue
			}
			tmp := node.Add(s)
			node = tmp
		}
		node = root
	}

	if err := gtree.ExecuteProgrammably(os.Stdout, root); err != nil {
		panic(err)
	}
	// Output:
	// .
	// ├── .github
	// │   └── workflows
	// │       ├── cd.yaml
	// │       └── ci.yaml
	// ├── .gitignore
	// ├── .goreleaser.yml
	// ├── cmd
	// │   └── gtree
	// │       └── main.go
	// ├── config.go
	// ├── example_programmable_test.go
	// ├── go.mod
	// ├── go.sum
	// ├── LICENSE
	// ├── node.go
	// ├── programmable.go
	// ├── programmable_test.go
	// ├── README.md
	// ├── sample
	// │   ├── find_pipe_programmable-gtree
	// │   │   └── main.go
	// │   ├── like_cli
	// │   │   ├── adapter
	// │   │   │   ├── executor.go
	// │   │   │   └── indentation.go
	// │   │   └── main.go
	// │   └── programmable
	// │       └── main.go
	// ├── stack.go
	// ├── testdata
	// │   ├── demo.md
	// │   ├── sample0.md
	// │   ├── sample1.md
	// │   ├── sample2.md
	// │   ├── sample3.md
	// │   ├── sample4.md
	// │   ├── sample5.md
	// │   └── sample6.md
	// ├── tmp.md
	// ├── tree.go
	// └── tree_test.go
}

```

- You can also output JSON.

[link](https://github.com/ddddddO/gtree/blob/master/sample/programmable/main.go#L61)

- You can also output YAML.

[link](https://github.com/ddddddO/gtree/blob/master/sample/programmable/main.go#L198)

- You can also output TOML.

[link](https://github.com/ddddddO/gtree/blob/master/sample/programmable/main.go#L262)

---
