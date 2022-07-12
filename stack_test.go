package gtree

import (
	"testing"
)

func TestStack_Pop(t *testing.T) {
	s := newStack()
	n1 := newNode("p", uint(1), uint(1))
	n2 := newNode("c", uint(2), uint(2))
	s.push(n1).push(n2)

	tests := map[string]struct {
		s    *stack
		want *Node
	}{
		"exists in stack":         {s, n2},
		"does not exist in stack": {newStack(), nil},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.s.pop()
			if got != tt.want {
				t.Errorf("\ngot: \n%+v\nwant: \n%+v", got, tt.want)
			}
		})
	}
}
