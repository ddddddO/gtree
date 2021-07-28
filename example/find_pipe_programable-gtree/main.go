package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

// cd github.com/ddddddO/gtree
// find . -type d -name .git -prune -o -type f -print | go run example/find_pipe_programable-gtree/main.go
func main() {
	var (
		root  *gtree.Node
		child *gtree.Node
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		splited := strings.Split(line, "/")

		for i, s := range splited {
			if i == 0 {
				if root == nil {
					root = gtree.NewRoot(s)
				}
				child = root
				continue
			}
			tmp := child.Add(s)
			child = tmp
		}
		child = root
	}

	if err := gtree.ExecuteProgrammably(os.Stdout, root); err != nil {
		panic(err)
	}
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
	// ├── example
	// │   ├── find_pipe_programable-gtree
	// │   │   └── main.go
	// │   ├── like_cli
	// │   │   ├── adapter
	// │   │   │   ├── executor.go
	// │   │   │   └── indentation.go
	// │   │   └── main.go
	// │   └── programable
	// │       └── main.go
	// ├── go.mod
	// ├── go.sum
	// ├── LICENSE
	// ├── node.go
	// ├── programable.go
	// ├── programable_test.go
	// ├── README.md
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
