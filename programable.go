package gtree

import (
	"github.com/pkg/errors"
	"io"
)

var ErrNotRoot = errors.New("not root node")

// TODO: 命名がイマイチ
func ExecuteProgrammably(root *node, w io.Writer) error {
	if !root.isRoot() {
		return ErrNotRoot
	}

	tree := &tree{
		roots: []*node{root},
		lastNodeFormat: lastNodeFormat{
			directly:   "└──",
			indirectly: "    ",
		},
		intermedialNodeFormat: intermedialNodeFormat{
			directly:   "├──",
			indirectly: "│   ",
		},
	}

	tree.grow()
	return tree.expand(w)
}

var programableNodeIdx int

func NewRoot(text string) *node {
	programableNodeIdx++

	return newNode(text, rootHierarchyNum, programableNodeIdx)
}

func (current *node) Add(text string) *node {
	programableNodeIdx++

	n := newNode(text, current.hierarchy+1, programableNodeIdx)
	n.parent = current
	current.children = append(current.children, n)
	return n
}
