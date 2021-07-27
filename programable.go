package gtree

import (
	"io"

	"github.com/pkg/errors"
)

var ErrNotRoot = errors.New("not root node")

// TODO: 命名がイマイチ
func ExecuteProgrammably(w io.Writer, root *Node) error {
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

func (parent *Node) Add(text string) *Node {
	for _, child := range parent.children {
		if text == child.text {
			return child
		}
	}

	programableNodeIdx++

	current := newNode(text, parent.hierarchy+1, programableNodeIdx)
	current.parent = parent
	parent.children = append(parent.children, current)
	return current
}
