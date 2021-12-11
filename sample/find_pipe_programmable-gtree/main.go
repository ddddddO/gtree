package main

import (
	"bufio"
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
	// ├── makefile
	// ├── node.go
	// ├── programmable.go
	// ├── programmable_test.go
	// ├── README.md
	// ├── README_CLI.md
	// ├── README_Package_1.md
	// ├── README_Package_2.md
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
	// ├── tree_json.go
	// ├── tree_json_test.go
	// ├── tree_test.go
	// ├── tree_toml.go
	// ├── tree_toml_test.go
	// ├── tree_yaml.go
	// └── tree_yaml_test.go
}
