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
