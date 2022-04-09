package gtree

import "github.com/pkg/errors"

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
		tmp := s.pop()
		if current.isDirectlyUnder(tmp) {
			tmp.addChild(current)
			current.setParent(tmp)
			s.push(tmp).push(current)
			return
		}
	}
}
