package gtree_test

import (
	"bufio"
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

func ExampleOutputFromMarkdown() {
	md := bytes.NewBufferString(strings.TrimSpace(`
- root
	- dddd
		- kkkkkkk
			- lllll
				- ffff
				- LLL
					- WWWWW
						- ZZZZZ
				- ppppp
					- KKK
						- 1111111
							- AAAAAAA
	- eee`))
	if err := gtree.OutputFromMarkdown(os.Stdout, md); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// root
	// ├── dddd
	// │   └── kkkkkkk
	// │       └── lllll
	// │           ├── ffff
	// │           ├── LLL
	// │           │   └── WWWWW
	// │           │       └── ZZZZZ
	// │           └── ppppp
	// │               └── KKK
	// │                   └── 1111111
	// │                       └── AAAAAAA
	// └── eee
}

func ExampleOutputFromMarkdown_second() {
	md := bytes.NewBufferString(strings.TrimSpace(`
- a
  - i
    - u
      - k
      - kk
    - t
  - e
    - o
  - g`))

	// You can customize branch format.
	if err := gtree.OutputFromMarkdown(os.Stdout, md,
		gtree.WithMidBranch("+->"),
		gtree.WithLastBranch("+->"),
		gtree.WithHLine(""),
		gtree.WithVLine(":"),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// a
	// +-> i
	// :   +-> u
	// :   :   +-> k
	// :   :   +-> kk
	// :   +-> t
	// +-> e
	// :   +-> o
	// +-> g
}

func ExampleWalkFromMarkdown() {
	md := strings.TrimSpace(`
- a
	- i
		- u
			- k
	- kk
		- t
- e
	- o
		- g`)

	callback := func(wn *gtree.WalkerNode) error {
		fmt.Println(wn.Row())
		return nil
	}

	if err := gtree.WalkFromMarkdown(strings.NewReader(md), callback); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// a
	// ├── i
	// │   └── u
	// │       └── k
	// └── kk
	//     └── t
	// e
	// └── o
	//     └── g

}

func ExampleWalkFromMarkdown_second() {
	md := strings.TrimSpace(`
- a
	- i
		- u
			- k
	- kk
		- t
- e
	- o
		- g`)

	callback := func(wn *gtree.WalkerNode) error {
		fmt.Println("WalkerNode's methods called...")
		fmt.Printf("\tName     : %s\n", wn.Name())
		fmt.Printf("\tBranch   : %s\n", wn.Branch())
		fmt.Printf("\tRow      : %s\n", wn.Row())
		fmt.Printf("\tLevel    : %d\n", wn.Level())
		fmt.Printf("\tPath     : %s\n", wn.Path())
		fmt.Printf("\tHasChild : %t\n", wn.HasChild())

		childNames := func(wn *gtree.WalkerNode) []string {
			names := make([]string, 0, len(wn.Children()))
			for _, child := range wn.Children() {
				names = append(names, child.Name())
			}
			return names
		}(wn)
		fmt.Printf("\tChildren : %v\n", childNames)

		ancestorNames := func(wn *gtree.WalkerNode) []string {
			names := []string{}
			for _, parent := range wn.Ancestors() {
				names = append(names, parent.Name())
			}
			return names
		}(wn)
		fmt.Printf("\tAncestors: %v\n", ancestorNames)
		return nil
	}

	if err := gtree.WalkFromMarkdown(strings.NewReader(md), callback); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// want:
	// WalkerNode's methods called...
	//	Name     : a
	//	Branch   :
	//	Row      : a
	//	Level    : 1
	//	Path     : a
	//	HasChild : true
	//	Children : [i kk]
	//	Ancestors: []
	// WalkerNode's methods called...
	//	Name     : i
	//	Branch   : ├──
	//	Row      : ├── i
	//	Level    : 2
	//	Path     : a/i
	//	HasChild : true
	//	Children : [u]
	//	Ancestors: [a]
	// WalkerNode's methods called...
	//	Name     : u
	//	Branch   : │   └──
	//	Row      : │   └── u
	//	Level    : 3
	//	Path     : a/i/u
	//	HasChild : true
	//	Children : [k]
	//	Ancestors: [a i]
	// WalkerNode's methods called...
	//	Name     : k
	//	Branch   : │       └──
	//	Row      : │       └── k
	//	Level    : 4
	//	Path     : a/i/u/k
	//	HasChild : false
	//	Children : []
	//	Ancestors: [a i u]
	// WalkerNode's methods called...
	//	Name     : kk
	//	Branch   : └──
	//	Row      : └── kk
	//	Level    : 2
	//	Path     : a/kk
	//	HasChild : true
	//	Children : [t]
	//	Ancestors: [a]
	// WalkerNode's methods called...
	//	Name     : t
	//	Branch   :     └──
	//	Row      :     └── t
	//	Level    : 3
	//	Path     : a/kk/t
	//	HasChild : false
	//	Children : []
	//	Ancestors: [a kk]
	// WalkerNode's methods called...
	//	Name     : e
	//	Branch   :
	//	Row      : e
	//	Level    : 1
	//	Path     : e
	//	HasChild : true
	//	Children : [o]
	//	Ancestors: []
	// WalkerNode's methods called...
	//	Name     : o
	//	Branch   : └──
	//	Row      : └── o
	//	Level    : 2
	//	Path     : e/o
	//	HasChild : true
	//	Children : [g]
	//	Ancestors: [e]
	// WalkerNode's methods called...
	//	Name     : g
	//	Branch   :     └──
	//	Row      :     └── g
	//	Level    : 3
	//	Path     : e/o/g
	//	HasChild : false
	//	Children : []
	//	Ancestors: [e o]
}

func ExampleOutputFromRoot() {
	var root *gtree.Node = gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *gtree.Node = root.Add("child 1").Add("child 2").Add("child 4")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	root.Add("child 8")
	// you can customize branch format.
	if err := gtree.OutputFromRoot(os.Stdout, root,
		gtree.WithMidBranch("+"),
		gtree.WithLastBranch("-"),
		gtree.WithVLine(":"),
		gtree.WithHLine(">>>"),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// root
	// +>>> child 1
	// :   ->>> child 2
	// :       +>>> child 3
	// :       ->>> child 4
	// :           +>>> child 5
	// :           ->>> child 6
	// :               ->>> child 7
	// ->>> child 8
}

type CustomWriter struct {
	logger *slog.Logger
}

func (c *CustomWriter) Write(bytes []byte) (int, error) {
	row := strings.TrimSuffix(string(bytes), "\n")
	c.logger.Info("[TREE]", "row", row)
	return len(bytes), nil
}

func ExampleOutputFromRoot_second() {
	w := &CustomWriter{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	var root *gtree.Node = gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *gtree.Node = root.Add("child 1").Add("child 2").Add("child 4")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	root.Add("child 8")

	if err := gtree.OutputFromRoot(w, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// want:
	// time=2025-05-17T10:51:35.902+09:00 level=INFO msg=[TREE] row=root
	// time=2025-05-17T10:51:35.902+09:00 level=INFO msg=[TREE] row="├── child 1"
	// time=2025-05-17T10:51:35.902+09:00 level=INFO msg=[TREE] row="│   └── child 2"
	// time=2025-05-17T10:51:35.902+09:00 level=INFO msg=[TREE] row="│       ├── child 3"
	// time=2025-05-17T10:51:35.902+09:00 level=INFO msg=[TREE] row="│       └── child 4"
	// time=2025-05-17T10:51:35.902+09:00 level=INFO msg=[TREE] row="│           ├── child 5"
	// time=2025-05-17T10:51:35.902+09:00 level=INFO msg=[TREE] row="│           └── child 6"
	// time=2025-05-17T10:51:35.902+09:00 level=INFO msg=[TREE] row="│               └── child 7"
	// time=2025-05-17T10:51:35.902+09:00 level=INFO msg=[TREE] row="└── child 8"
}

func ExampleOutputFromRoot_third() {
	preparePrimate := func() *gtree.Node {
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

	primate := preparePrimate()
	// default branch format.
	if err := gtree.OutputFromRoot(os.Stdout, primate); err != nil {
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
	//     ├── Simiiformes
	//     │   └── Catarrhini
	//     │       ├── Cercopithecoidea
	//     │       │   └── Cercopithecidae
	//     │       └── Hominoidea
	//     │           ├── Hylobatidae
	//     │           └── Hominidae
	//     └── Platyrrhini
	//         ├── Ceboidea
	//         │   ├── Atelidae
	//         │   └── Cebidae
	//         └── Pithecioidea
	//             └── Pitheciidae
}

func ExampleOutputFromRoot_fourth() {
	// Example: The program below converts the result of `find` into a tree.
	//
	// $ cd github.com/ddddddO/gtree
	// $ find . -type d -name .git -prune -o -type f -print
	// ./config.go
	// ./node_generator_test.go
	// ./example/like_cli/adapter/indentation.go
	// ./example/like_cli/adapter/executor.go
	// ./example/like_cli/main.go
	// ./example/find_pipe_programmable-gtree/main.go
	// ...
	// $ find . -type d -name .git -prune -o -type f -print | go run example/find_pipe_programmable-gtree/main.go
	// << See "want:" below. >>
	var (
		root *gtree.Node
		node *gtree.Node
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()              // e.g.) "./example/find_pipe_programmable-gtree/main.go"
		splited := strings.Split(line, "/") // e.g.) [. example find_pipe_programmable-gtree main.go]

		for i, s := range splited {
			if root == nil {
				root = gtree.NewRoot(s) // s := "."
				node = root
				continue
			}
			if i == 0 {
				continue
			}

			tmp := node.Add(s)
			node = tmp
		}
		node = root
	}

	if err := gtree.OutputFromRoot(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// want:
	// .
	// ├── config.go
	// ├── node_generator_test.go
	// ├── example
	// │   ├── like_cli
	// │   │   ├── adapter
	// │   │   │   ├── indentation.go
	// │   │   │   └── executor.go
	// │   │   └── main.go
	// │   ├── find_pipe_programmable-gtree
	// │   │   └── main.go
	// │   ├── go-list_pipe_programmable-gtree
	// │   │   └── main.go
	// │   └── programmable
	// │       └── main.go
	// ├── file_considerer.go
	// ├── node.go
	// ├── node_generator.go
	// ├── .gitignore
	// ...
}

func ExampleMkdirFromRoot() {
	preparePrimate := func() *gtree.Node {
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

	if err := gtree.MkdirFromRoot(preparePrimate()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// want(using Linux 'tree' command):
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

func ExampleMkdirFromRoot_second() {
	gtreeDir := gtree.NewRoot("gtree")
	gtreeDir.Add("cmd").Add("gtree").Add("main.go")
	gtreeDir.Add("Makefile")
	testdataDir := gtreeDir.Add("testdata")
	testdataDir.Add("sample1.md")
	testdataDir.Add("sample2.md")
	gtreeDir.Add("tree.go")

	// make directories and files with specific extensions.
	if err := gtree.MkdirFromRoot(
		gtreeDir,
		gtree.WithFileExtensions([]string{".go", ".md", "Makefile"}),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// want(using Linux 'tree' command):
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

func ExampleWalkFromRoot() {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	root.Add("child 5")
	root.Add("child 1").Add("child 2").Add("child 4")

	callback := func(wn *gtree.WalkerNode) error {
		fmt.Println(wn.Row())
		return nil
	}

	if err := gtree.WalkFromRoot(root, callback); err != nil {
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
}

func ExampleWalkFromRoot_second() {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	root.Add("child 5")
	root.Add("child 1").Add("child 2").Add("child 4")

	callback := func(wn *gtree.WalkerNode) error {
		fmt.Println("WalkerNode's methods called...")
		fmt.Printf("\tName     : %s\n", wn.Name())
		fmt.Printf("\tBranch   : %s\n", wn.Branch())
		fmt.Printf("\tRow      : %s\n", wn.Row())
		fmt.Printf("\tLevel    : %d\n", wn.Level())
		fmt.Printf("\tPath     : %s\n", wn.Path())
		fmt.Printf("\tHasChild : %t\n", wn.HasChild())

		childNames := func(wn *gtree.WalkerNode) []string {
			names := make([]string, 0, len(wn.Children()))
			for _, child := range wn.Children() {
				names = append(names, child.Name())
			}
			return names
		}(wn)
		fmt.Printf("\tChildren : %v\n", childNames)

		ancestorNames := func(wn *gtree.WalkerNode) []string {
			names := []string{}
			for _, parent := range wn.Ancestors() {
				names = append(names, parent.Name())
			}
			return names
		}(wn)
		fmt.Printf("\tAncestors: %v\n", ancestorNames)

		return nil
	}

	if err := gtree.WalkFromRoot(root, callback); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// want:
	// WalkerNode's methods called...
	//         Name     : root
	//         Branch   :
	//         Row      : root
	//         Level    : 1
	//         Path     : root
	//         HasChild : true
	//         Children : [child 1 child 5]
	//         Ancestors: []
	// WalkerNode's methods called...
	//         Name     : child 1
	//         Branch   : ├──
	//         Row      : ├── child 1
	//         Level    : 2
	//         Path     : root/child 1
	//         HasChild : true
	//         Children : [child 2]
	//         Ancestors: [root]
	// WalkerNode's methods called...
	//         Name     : child 2
	//         Branch   : │   └──
	//         Row      : │   └── child 2
	//         Level    : 3
	//         Path     : root/child 1/child 2
	//         HasChild : true
	//         Children : [child 3 child 4]
	//         Ancestors: [root child 1]
	// WalkerNode's methods called...
	//         Name     : child 3
	//         Branch   : │       ├──
	//         Row      : │       ├── child 3
	//         Level    : 4
	//         Path     : root/child 1/child 2/child 3
	//         HasChild : false
	//         Children : []
	//         Ancestors: [root child 1 child 2]
	// WalkerNode's methods called...
	//         Name     : child 4
	//         Branch   : │       └──
	//         Row      : │       └── child 4
	//         Level    : 4
	//         Path     : root/child 1/child 2/child 4
	//         HasChild : false
	//         Children : []
	//         Ancestors: [root child 1 child 2]
	// WalkerNode's methods called...
	//         Name     : child 5
	//         Branch   : └──
	//         Row      : └── child 5
	//         Level    : 2
	//         Path     : root/child 5
	//         HasChild : false
	//         Children : []
	//         Ancestors: [root]
}

func ExampleWalkIterFromRoot() {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	root.Add("child 5")
	root.Add("child 1").Add("child 2").Add("child 4")

	for wn, err := range gtree.WalkIterFromRoot(root) {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(wn.Row())
	}
	// Output:
	// root
	// ├── child 1
	// │   └── child 2
	// │       ├── child 3
	// │       └── child 4
	// └── child 5
}

func ExampleWalkIterFromRoot_second() {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	root.Add("child 5")
	root.Add("child 1").Add("child 2").Add("child 4")

	for wn, err := range gtree.WalkIterFromRoot(root) {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("WalkerNode's methods called...")
		fmt.Printf("\tName     : %s\n", wn.Name())
		fmt.Printf("\tBranch   : %s\n", wn.Branch())
		fmt.Printf("\tRow      : %s\n", wn.Row())
		fmt.Printf("\tLevel    : %d\n", wn.Level())
		fmt.Printf("\tPath     : %s\n", wn.Path())
		fmt.Printf("\tHasChild : %t\n", wn.HasChild())

		childNames := func(wn *gtree.WalkerNode) []string {
			names := make([]string, 0, len(wn.Children()))
			for _, child := range wn.Children() {
				names = append(names, child.Name())
			}
			return names
		}(wn)
		fmt.Printf("\tChildren : %v\n", childNames)

		ancestorNames := func(wn *gtree.WalkerNode) []string {
			names := []string{}
			for _, parent := range wn.Ancestors() {
				names = append(names, parent.Name())
			}
			return names
		}(wn)
		fmt.Printf("\tAncestors: %v\n", ancestorNames)
	}
	// want:
	// WalkerNode's methods called...
	//         Name     : root
	//         Branch   :
	//         Row      : root
	//         Level    : 1
	//         Path     : root
	//         HasChild : true
	//         Children : [child 1 child 5]
	//         Ancestors: []
	// WalkerNode's methods called...
	//         Name     : child 1
	//         Branch   : ├──
	//         Row      : ├── child 1
	//         Level    : 2
	//         Path     : root/child 1
	//         HasChild : true
	//         Children : [child 2]
	//         Ancestors: [root]
	// WalkerNode's methods called...
	//         Name     : child 2
	//         Branch   : │   └──
	//         Row      : │   └── child 2
	//         Level    : 3
	//         Path     : root/child 1/child 2
	//         HasChild : true
	//         Children : [child 3 child 4]
	//         Ancestors: [root child 1]
	// WalkerNode's methods called...
	//         Name     : child 3
	//         Branch   : │       ├──
	//         Row      : │       ├── child 3
	//         Level    : 4
	//         Path     : root/child 1/child 2/child 3
	//         HasChild : false
	//         Children : []
	//         Ancestors: [root child 1 child 2]
	// WalkerNode's methods called...
	//         Name     : child 4
	//         Branch   : │       └──
	//         Row      : │       └── child 4
	//         Level    : 4
	//         Path     : root/child 1/child 2/child 4
	//         HasChild : false
	//         Children : []
	//         Ancestors: [root child 1 child 2]
	// WalkerNode's methods called...
	//         Name     : child 5
	//         Branch   : └──
	//         Row      : └── child 5
	//         Level    : 2
	//         Path     : root/child 5
	//         HasChild : false
	//         Children : []
	//         Ancestors: [root]
}
