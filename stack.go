package gentree

var singleton *stack

type stack struct {
	nodes []*node
}

func getStack() *stack {
	if singleton == nil {
		singleton = &stack{}
	}
	return singleton
}

func (s *stack) push(n *node) {
	s.nodes = append(s.nodes, n)
}

func (s *stack) pop() *node {
	lastIndex := len(s.nodes) - 1
	tmp := s.nodes[lastIndex]
	s.nodes = s.nodes[:lastIndex]
	return tmp
}

func (s *stack) size() int {
	return len(s.nodes)
}
