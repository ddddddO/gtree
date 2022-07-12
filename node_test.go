package gtree

import (
	"testing"
)

func TestNode_IsDirectlyUnder(t *testing.T) {
	tests := map[string]struct {
		parent *Node
		child  *Node
		want   bool
	}{
		"true":                        {newNode("p", uint(1), uint(1)), newNode("c", uint(2), uint(2)), true},
		"false":                       {newNode("c", uint(2), uint(1)), newNode("p", uint(1), uint(2)), false},
		"false(argument node is nil)": {nil, newNode("c", uint(2), uint(2)), false},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.child.isDirectlyUnder(tt.parent)
			if got != tt.want {
				t.Errorf("\ngot: \n%t\nwant: \n%t", got, tt.want)
			}
		})
	}
}
