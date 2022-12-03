package gtree_test

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ddddddO/gtree"
)

func ExampleProgram_Output() {
	var root *gtree.Node = gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *gtree.Node = root.Add("child 1").Add("child 2").Add("child 4")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	root.Add("child 8")

	buf := &bytes.Buffer{}
	err := gtree.NewProgram(buf, os.Stderr).
		FormatIntermedialNode("+--", ":   ").
		FormatLastNode("+--", "    ").
		Output(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(buf.String())
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
}
