package gtree_test

import (
	"bytes"
	"fmt"

	"github.com/ddddddO/gtree/v6"
)

func Example() {
	var root *gtree.Node
	root = gtree.NewRoot("root")
	root.Add("child 1").Add("child 2")
	root.Add("child 1").Add("child 3")
	child4 := root.Add("child 4")

	var child7 *gtree.Node
	child7 = child4.Add("child 5").Add("child 6").Add("child 7")
	child7.Add("child 8")

	buf := &bytes.Buffer{}
	if err := gtree.ExecuteProgrammably(buf, root); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
	// Output:
	// root
	// ├── child 1
	// │   ├── child 2
	// │   └── child 3
	// └── child 4
	//     └── child 5
	//         └── child 6
	//             └── child 7
	//                 └── child 8
}
