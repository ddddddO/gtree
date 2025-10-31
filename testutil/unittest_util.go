package testutil

import (
	"os"
	"testing"

	"github.com/ddddddO/gtree"
)

func PrepareMarkdownFile(t *testing.T) *os.File {
	const testfilepath = "./testdata/sample6.md"
	file, err := os.Open(testfilepath)
	if err != nil {
		t.Fatal(err)
	}
	return file
}

func Prepare() *gtree.Node {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2")
	return root
}

func PrepareSameNameChild() *gtree.Node {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2")
	root.Add("child 1").Add("child 3")
	return root
}

func PrepareNotRoot() *gtree.Node {
	root := gtree.NewRoot("root")
	child1 := root.Add("child 1")
	return child1
}

func PrepareNilNode() *gtree.Node {
	var node *gtree.Node
	return node
}

func PrepareMultiNode() *gtree.Node {
	var root *gtree.Node = gtree.NewRoot("root1")
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *gtree.Node = root.Add("child 1").Add("child 2").Add("child 4")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	root.Add("child 8")
	return root
}

func PrepareMultiNodeWithDuplicationAllowed() *gtree.Node {
	var root *gtree.Node = gtree.NewRoot("root1", gtree.WithDuplicationAllowed())
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *gtree.Node = root.Add("child 1").Add("child 2").Add("child 4")
	child4.Add("child 5")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	child4.Add("child 6").Add("child 9")
	root.Add("child 8")
	return root
}

func PrepareInvalidNodeName() *gtree.Node {
	var root *gtree.Node = gtree.NewRoot("root1")
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *gtree.Node = root.Add("child 1").Add("child 2").Add("chi/ld 4")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	root.Add("child 8")
	return root
}

func PrepareExistRoot(t *testing.T) *gtree.Node {
	name := "gtreetest"

	if err := os.MkdirAll(name, 0o755); err != nil {
		t.Fatal(err)
	}

	root := gtree.NewRoot(name)
	root.Add("temp")
	return root
}

func Prepare_a() *gtree.Node {
	root := gtree.NewRoot("root8")
	root.Add("child 1").Add("child 2")
	return root
}
