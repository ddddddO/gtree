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

// $ gtree gocode > find_to_tree.go
// $ find . -type d -name .git -prune -o -type f -print | go run find_to_tree.go
func main() {
	var (
		root *gtree.Node
		node *gtree.Node
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text() // e.g.) "./example/find_pipe_programmable-gtree/main.go"
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

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`
