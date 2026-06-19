package gtree

import (
	"testing"
)

func TestNode_IsDirectlyUnder(t *testing.T) {
	idxCounter := newCounter()
	tests := map[string]struct {
		parent *Node
		child  *Node
		want   bool
	}{
		"true":                        {newNode("p", uint(1), idxCounter), newNode("c", uint(2), idxCounter), true},
		"false":                       {newNode("c", uint(2), idxCounter), newNode("p", uint(1), idxCounter), false},
		"false(argument node is nil)": {nil, newNode("c", uint(2), idxCounter), false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.child.isDirectlyUnder(tt.parent)
			if got != tt.want {
				t.Errorf("\ngot: \n%t\nwant: \n%t", got, tt.want)
			}
		})
	}
}

func TestNode_IsLastOfHierarchy(t *testing.T) {
	trueNode := func() *Node {
		idxCounter := newCounter()
		parent := newNode("p", uint(1), idxCounter)
		child := newNode("c", uint(2), idxCounter)
		parent.addChild(child)
		child.setParent(parent)
		return child
	}

	falseNode := func() *Node {
		idxCounter := newCounter()
		parent := newNode("p", uint(1), idxCounter)
		child1 := newNode("c1", uint(2), idxCounter)
		child2 := newNode("c2", uint(2), idxCounter)
		parent.addChild(child1)
		child1.setParent(parent)
		parent.addChild(child2)
		child2.setParent(parent)
		return child1
	}

	falseNodeNotSetParent := func() *Node {
		idxCounter := newCounter()
		child := newNode("c", uint(2), idxCounter)
		return child
	}

	tests := map[string]struct {
		node *Node
		want bool
	}{
		"true":                  {trueNode(), true},
		"false":                 {falseNode(), false},
		"false(not set parent)": {falseNodeNotSetParent(), false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.node.isLastOfHierarchy()
			if got != tt.want {
				t.Errorf("\ngot: \n%t\nwant: \n%t", got, tt.want)
			}
		})
	}
}
