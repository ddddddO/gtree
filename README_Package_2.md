# Package(2) / generate a tree programmatically

## Installation

Go version requires 1.18 or later.

```console
go get github.com/ddddddO/gtree
```

## Usage

## *OutputProgrammably* func

```go
package main

import (
	"fmt"
	"os"

	"github.com/ddddddO/gtree"
)

func main() {
	var root *gtree.Node = gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *gtree.Node = root.Add("child 1").Add("child 2").Add("child 4")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	root.Add("child 8")
	// you can customize branch format.
	if err := gtree.OutputProgrammably(os.Stdout, root,
		gtree.WithBranchFormatIntermedialNode("+--", ":   "),
		gtree.WithBranchFormatLastNode("+--", "    "),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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

- You can also output JSON.

[link](https://github.com/ddddddO/gtree/blob/master/sample/programmable/main.go#L61)

- You can also output YAML.

[link](https://github.com/ddddddO/gtree/blob/master/sample/programmable/main.go#L198)

- You can also output TOML.

[link](https://github.com/ddddddO/gtree/blob/master/sample/programmable/main.go#L262)


---

- The program below converts the result of `find` into a tree.
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
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

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// .
	// ├── .github
	// │   ├── dependabot.yml
	// │   └── workflows
	// │       ├── cd.yaml
	// │       └── ci.yaml
	// ├── .gitignore
	// ├── .goreleaser.yml
	// ├── cli_mkdir_dryrun.png
	// ├── cmd
	// │   └── gtree
	// │       ├── indent.go
	// │       ├── main.go
	// │       ├── mkdir.go
	// │       └── output.go
	// ├── config.go
	// ├── counter.go
	// ├── CREDITS
	// ├── example_tree_handler_programmable_test.go
	// ├── file_consider.go
	// ├── go.mod
	// ├── go.sum
	// ├── LICENSE
	// ├── Makefile
	// ├── node.go
	// ├── node_generate.go
	// ├── node_generate_strategy.go
	// ├── node_generate_strategy_test.go
	// ├── README.md
	// ├── README_CLI.md
	// ├── README_Package_1.md
	// ├── README_Package_2.md
	// ├── sample
	// │   ├── find_pipe_programmable-gtree
	// │   │   └── main.go
	// │   ├── go-list_pipe_programmable-gtree
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
	// ├── tree_grow.go
	// ├── tree_handler.go
	// ├── tree_handler_programmable.go
	// ├── tree_handler_programmable_test.go
	// ├── tree_handler_test.go
	// ├── tree_mkdir.go
	// └── tree_spread.go
}

```

- Convert `go list -deps ./...` to tree. (inspired by https://github.com/nikolaydubina/go-recipes#readme)

[link](https://github.com/ddddddO/gtree/blob/master/sample/go-list_pipe_programmable-gtree/main.go)

## *MkdirProgrammably* func

```go
package main

import (
	"fmt"

	"github.com/ddddddO/gtree"
)

func main() {
	primate := preparePrimate()
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
}
```

[details](https://github.com/ddddddO/gtree/blob/master/sample/programmable/main.go#L354)

---

- Make directories and files with specific extensions.

```go
package main

import (
	"fmt"

	"github.com/ddddddO/gtree"
)

func main() {
	gtreeDir := gtree.NewRoot("gtree")
	gtreeDir.Add("cmd").Add("main.go")
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
```