package main

import (
	"fmt"
	"strings"
)

type gocode string

func (g gocode) print() error {
	_, err := fmt.Print(strings.Trim(string(g), "\n"))
	return err
}

func (g gocode) println() error {
	if err := g.print(); err != nil {
		return err
	}
	_, err := fmt.Println()
	return err
}

const findToTree gocode = `
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

// $ gtree gocode > find_to_tree.go && go mod init xxx 2>/dev/null && go mod tidy 2>/dev/null && find . -type d -o -type f -print | go run find_to_tree.go
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
			if root == nil {
				root = gtree.NewRoot(s)
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

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`

const goDependencesToTree gocode = `
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

// $ ls | grep go.mod && go list -deps ./path/to/go_dir > go_dependences.txt
// $ mkdir tmp && cd tmp && gtree gocode --godeps-to-tree > godeps_to_tree.go && go mod init xxx 2>/dev/null && go mod tidy 2>/dev/null && cat ../go_dependences.txt | go run godeps_to_tree.go
func main() {
	var (
		root = gtree.NewRoot("[All Dependencies]")
		node *gtree.Node
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		splited := strings.Split(line, "/")

		for i, s := range splited {
			if i == 0 {
				node = root.Add(s)
				continue
			}
			node = node.Add(s)
		}
	}

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`
