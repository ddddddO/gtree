package gtree

import (
	"fmt"
	"io"
)

type jsonTree struct {
	defaultTree

	tmp string
}

func newJSONTree(conf *config) *jsonTree {
	return &jsonTree{}
}

func (t *jsonTree) grow() tree {
	// 一旦、1 rootで考える
	root := t.roots[0]
	t.assembleJSON(root)

	return t
}

func (t *jsonTree) assembleJSON(current *Node) {
	t.tmp += "{"
	t.assembleObject(current)

	if hasChildrenOfAnyChild(current.children) {
		t.tmp += "{"
		t.assembleObject(current.children[0])
	} else {
		t.assembleArray(current.children)
	}

	t.tmp += "}"
}

func hasChildrenOfAnyChild(children []*Node) bool {
	for _, child := range children {
		if len(child.children) != 0 {
			return true
		}
	}
	return false
}

func (t *jsonTree) assembleObject(current *Node) {
	tmp := quote(current.text) + ":"
	t.tmp += tmp
}

func (t *jsonTree) assembleArray(nodes []*Node) {
	tmp := "["
	for i, n := range nodes {
		tmp += quote(n.text)

		if i != len(nodes)-1 {
			tmp += ","
		}
	}
	tmp += "]"

	t.tmp += tmp
}

func quote(text string) string {
	return `"` + text + `"`
}

func (t *jsonTree) expand(w io.Writer) error {
	fmt.Fprint(w, t.tmp)
	return nil
}
