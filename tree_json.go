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
	t.tmp += "{"

	// 一旦、1 rootで考える
	root := t.roots[0]
	t.assembleJSON(root)

	t.tmp += "}"
	return t
}

func (t *jsonTree) assembleJSON(current *Node) {
	t.assembleObject(current)

	for _, child := range current.children {
		t.assembleArray(child)
	}
}

func (t *jsonTree) assembleObject(current *Node) {
	tmp := quote(current.text) + ":"
	t.tmp += tmp
}

func (t *jsonTree) assembleArray(current *Node) {
	tmp := "[" + quote(current.text) + "]"
	t.tmp += tmp
}

func quote(text string) string {
	return `"` + text + `"`
}

func (t *jsonTree) expand(w io.Writer) error {
	fmt.Fprint(w, t.tmp)
	return nil
}
