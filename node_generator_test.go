package gtree

// ref: https://zenn.dev/kimuson13/articles/go_table_driven_test

import (
	"testing"
)

type want struct {
	name      string
	hierarchy uint
	index     uint
	err       error
}

const fixedIndex uint = 1

func TestGenerateTab(t *testing.T) {
	tests := map[string]struct {
		row  string
		want *want
	}{
		"root/hierarchy=1":                 {"- aaa bb", &want{name: "aaa bb", hierarchy: 1, index: fixedIndex, err: nil}},
		"child/hierarchy=2":                {"	- aaa bb", &want{name: "aaa bb", hierarchy: 2, index: fixedIndex, err: nil}},
		"child/hierarchy=2/tab on the way": {"	- aaa	bb", &want{name: "aaa	bb", hierarchy: 2, index: fixedIndex, err: nil}},
		"invalid/hierarchy=0/prefix chars": {"xx- aaa bb", &want{err: &inputFormatError{row: "xx- aaa bb"}}},
		"invalid/hierarchy=0/no hyphen":    {"xx aaa bb", &want{err: &inputFormatError{row: "xx aaa bb"}}},
		"invalid/hierarchy=0/tab only":     {"			", nil},
	}

	nodeGenerator := newNodeGenerator()
	_, _ = nodeGenerator.generate("	- xxx", fixedIndex) // Parserのインデントスペース数を決めるために必要

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			node, err := nodeGenerator.generate(tt.row, fixedIndex)
			if node == nil && err == nil {
				return
			}
			if tt.want.err != nil {
				if e, ok := err.(*inputFormatError); !ok {
					t.Errorf("\ngot: \n%v\nwant: \n%v", e, tt.want.err)
				}
				return
			}

			if node.name != tt.want.name {
				t.Errorf("\ngot: \n%s\nwant: \n%s", node.name, tt.want.name)
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

func TestGenerateTwoSpaces(t *testing.T) {
	tests := map[string]struct {
		row  string
		want *want
	}{
		"root/hierarchy=1":                     {"- aaa bb", &want{name: "aaa bb", hierarchy: 1, index: fixedIndex, err: nil}},
		"child/hierarchy=2":                    {"  - aaa bb", &want{name: "aaa bb", hierarchy: 2, index: fixedIndex, err: nil}},
		"invalid/hierarchy=0/prefix odd space": {" - aaa bb", &want{err: &inputFormatError{row: " - aaa bb"}}},
		"invalid/hierarchy=0/prefix chars":     {"xx- aaa bb", &want{err: &inputFormatError{row: "xx- aaa bb"}}},
		"invalid/hierarchy=0/no hyphen":        {"xx aaa bb", &want{err: &inputFormatError{row: "xx aaa bb"}}},
		"invalid/hierarchy=0/space only":       {"  ", nil},
	}

	nodeGenerator := newNodeGenerator()
	_, _ = nodeGenerator.generate("  - xxx", fixedIndex) // Parserのインデントスペース数を決めるために必要

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			node, err := nodeGenerator.generate(tt.row, fixedIndex)
			if node == nil && err == nil {
				return
			}
			if tt.want.err != nil {
				if e, ok := err.(*inputFormatError); !ok {
					t.Errorf("\ngot: \n%v\nwant: \n%v", e, tt.want.err)
				}
				return
			}

			if node.name != tt.want.name {
				t.Errorf("\ngot: \n%s\nwant: \n%s", node.name, tt.want.name)
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

func TestGenerateFourSpaces(t *testing.T) {
	tests := map[string]struct {
		row  string
		want *want
	}{
		"root/hierarchy=1":                     {"- aaa bb", &want{name: "aaa bb", hierarchy: 1, index: fixedIndex, err: nil}},
		"child/hierarchy=2":                    {"    - aaa bb", &want{name: "aaa bb", hierarchy: 2, index: fixedIndex, err: nil}},
		"child/hierarchy=3":                    {"        - aaa    bb", &want{name: "aaa    bb", hierarchy: 3, index: fixedIndex, err: nil}},
		"invalid/hierarchy=0/prefix odd space": {" - aaa bb", &want{err: &inputFormatError{row: " - aaa bb"}}},
		"invalid/hierarchy=0/prefix chars":     {"xx- aaa bb", &want{err: &inputFormatError{row: "xx- aaa bb"}}},
		"invalid/hierarchy=0/no hyphen":        {"xx aaa bb", &want{err: &inputFormatError{row: "xx aaa bb"}}},
		"invalid/hierarchy=0/space only":       {"    ", nil},
	}

	nodeGenerator := newNodeGenerator()
	_, _ = nodeGenerator.generate("    - xxx", fixedIndex) // Parserのインデントスペース数を決めるために必要

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			node, err := nodeGenerator.generate(tt.row, fixedIndex)
			if node == nil && err == nil {
				return
			}
			if tt.want.err != nil {
				if e, ok := err.(*inputFormatError); !ok {
					t.Errorf("\ngot: \n%v\nwant: \n%v", e, tt.want.err)
				}
				return
			}

			if node.name != tt.want.name {
				t.Errorf("\ngot: \n%s\nwant: \n%s", node.name, tt.want.name)
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
