package gtree

import "io"

// TODO: 命名がイマイチ
func ExecuteProgrammably(root *node, w io.Writer) error {
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

func NewRoot(name string) *node {
	programableNodeIdx++

	return newNode(name, rootHierarchyNum, programableNodeIdx)
}

func (current *node) Add(name string) *node {
	programableNodeIdx++

	n := newNode(name, current.hierarchy+1, programableNodeIdx)
	n.parent = current
	current.children = append(current.children, n)
	return n
}
