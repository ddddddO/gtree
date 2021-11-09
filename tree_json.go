package gtree

import (
	"fmt"
	"io"
)

type jsonTree struct {
	defaultTree
}

func newJSONTree(conf *config) *jsonTree {
	return &jsonTree{}
}

func (t *jsonTree) grow() tree {
	return t
}

func (t *jsonTree) expand(w io.Writer) error {
	fmt.Fprint(w, "not yet impl")
	return nil
}
