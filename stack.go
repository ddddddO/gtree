package gtree

import "github.com/pkg/errors"

var errNilStack = errors.New("nil stack")

type stack struct {
	nodes []*Node
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(n *Node) {
	s.nodes = append(s.nodes, n)
}

func (s *stack) pop() *Node {
	lastIndex := s.size() - 1
	tmp := s.nodes[lastIndex]
	s.nodes = s.nodes[:lastIndex]
	return tmp
}

func (s *stack) size() int {
	return len(s.nodes)
}
