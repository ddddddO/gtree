package gtree

import (
	"testing"
)

type want struct {
	Name      string
	hierarchy uint
	index     uint
}

const fixedIndex uint = 1

func TestTabStrategy_Generate(t *testing.T) {
	tests := []struct {
		name string
		row  string
		want want
	}{
		{"root/hierarchy=1", "- aaa bb", want{Name: "aaa bb", hierarchy: 1, index: fixedIndex}},
		{"child/hierarchy=2", "	- aaa bb", want{Name: "aaa bb", hierarchy: 2, index: fixedIndex}},
		{"child/hierarchy=2/tab on the way", "	- aaa	bb", want{Name: "aaa	bb", hierarchy: 2, index: fixedIndex}},
		// {"invalid/hierarchy=0/prefix space", " - aaa bb", want{Name: "", hierarchy: invalidHierarchyNum, index: fixedIndex}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			node := (*tabStrategy)(nil).generate(tt.row, fixedIndex)
			if node.Name != tt.want.Name {
				t.Errorf("\ngot: \n%s\nwant: \n%s", node.Name, tt.want.Name)
			}
			if node.hierarchy != tt.want.hierarchy {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.hierarchy, tt.want.hierarchy)
			}
			if node.index != tt.want.index {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.index, tt.want.index)
			}
		})
	}
}

func TestTwoSpacesStrategy_Generate(t *testing.T) {
	tests := []struct {
		name string
		row  string
		want want
	}{
		{"root/hierarchy=1", "- aaa bb", want{Name: "aaa bb", hierarchy: 1, index: fixedIndex}},
		{"child/hierarchy=2", "  - aaa bb", want{Name: "aaa bb", hierarchy: 2, index: fixedIndex}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			node := (*twoSpacesStrategy)(nil).generate(tt.row, fixedIndex)
			if node.Name != tt.want.Name {
				t.Errorf("\ngot: \n%s\nwant: \n%s", node.Name, tt.want.Name)
			}
			if node.hierarchy != tt.want.hierarchy {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.hierarchy, tt.want.hierarchy)
			}
			if node.index != tt.want.index {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.index, tt.want.index)
			}
		})
	}
}

func TestFourSpacesStrategy_Generate(t *testing.T) {
	tests := []struct {
		name string
		row  string
		want want
	}{
		{"root/hierarchy=1", "- aaa bb", want{Name: "aaa bb", hierarchy: 1, index: fixedIndex}},
		{"child/hierarchy=2", "    - aaa bb", want{Name: "aaa bb", hierarchy: 2, index: fixedIndex}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			node := (*fourSpacesStrategy)(nil).generate(tt.row, fixedIndex)
			if node.Name != tt.want.Name {
				t.Errorf("\ngot: \n%s\nwant: \n%s", node.Name, tt.want.Name)
			}
			if node.hierarchy != tt.want.hierarchy {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.hierarchy, tt.want.hierarchy)
			}
			if node.index != tt.want.index {
				t.Errorf("\ngot: \n%d\nwant: \n%d", node.index, tt.want.index)
			}
		})
	}
}
