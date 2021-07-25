package gtree

import (
	"io"

	"github.com/pkg/errors"
)

var ErrNotRoot = errors.New("not root node")

// TODO: 命名がイマイチ
func ExecuteProgrammably(root *Node, w io.Writer) error {
	if !root.isRoot() {
		return ErrNotRoot
	}

	tree := &tree{
		roots: []*Node{root},
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

func NewRoot(text string) *Node {
	programableNodeIdx++

	return newNode(text, rootHierarchyNum, programableNodeIdx)
}

func (current *Node) Add(text string) *Node {
	programableNodeIdx++

	n := newNode(text, current.hierarchy+1, programableNodeIdx)
	n.parent = current
	current.children = append(current.children, n)
	return n
}
