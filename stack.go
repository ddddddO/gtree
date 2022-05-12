package gtree

import "errors"

var errNilStack = errors.New("nil stack")

type stack struct {
	nodes []*Node
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(n *Node) *stack {
	s.nodes = append(s.nodes, n)
	return s
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

// depth-first search
func (s *stack) dfs(current *Node) {
	size := s.size()
	for i := 0; i < size; i++ {
		parent := s.pop()
		if !current.isDirectlyUnder(parent) {
			continue
		}

		parent.addChild(current)
		current.setParent(parent)
		s.push(parent).push(current)
		return
	}
}
