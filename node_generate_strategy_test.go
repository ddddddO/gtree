package gtree

import (
	"testing"
)

type want struct {
	name      string
	hierarchy uint
	index     uint
}

const fixedIndex uint = 1

func TestTabStrategy_Generate(t *testing.T) {
	// ref: https://zenn.dev/kimuson13/articles/go_table_driven_test
	tests := map[string]struct {
		row     string
		want    *want
		wantErr error
	}{
		"root/hierarchy=1": {"- aaa bb", &want{name: "aaa bb", hierarchy: 1, index: fixedIndex}, nil},
		"child/hierarchy=2": {"	- aaa bb", &want{name: "aaa bb", hierarchy: 2, index: fixedIndex}, nil},
		"child/hierarchy=2/tab on the way": {"	- aaa	bb", &want{name: "aaa	bb", hierarchy: 2, index: fixedIndex}, nil},
		"invalid/hierarchy=0/prefix space": {" - aaa bb", nil, errIncorrectFormat},
		"invalid/hierarchy=0/prefix chars": {"xx- aaa bb", nil, errIncorrectFormat},
		"invalid/hierarchy=0/no hyphen":    {"xx aaa bb", nil, errIncorrectFormat},
		"invalid/hierarchy=0/tab only": {"			", nil, errIncorrectFormat},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			node, err := (*tabStrategy)(nil).generate(tt.row, fixedIndex)
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("\ngot: \n%v\nwant: \n%v", err, tt.wantErr)
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

func TestTwoSpacesStrategy_Generate(t *testing.T) {
	tests := map[string]struct {
		row     string
		want    *want
		wantErr error
	}{
		"root/hierarchy=1":                     {"- aaa bb", &want{name: "aaa bb", hierarchy: 1, index: fixedIndex}, nil},
		"child/hierarchy=2":                    {"  - aaa bb", &want{name: "aaa bb", hierarchy: 2, index: fixedIndex}, nil},
		"invalid/hierarchy=0/prefix odd space": {" - aaa bb", nil, errIncorrectFormat},
		"invalid/hierarchy=0/prefix chars":     {"xx- aaa bb", nil, errIncorrectFormat},
		"invalid/hierarchy=0/no hyphen":        {"xx aaa bb", nil, errIncorrectFormat},
		"invalid/hierarchy=0/space only":       {"  ", nil, errIncorrectFormat},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			node, err := (*twoSpacesStrategy)(nil).generate(tt.row, fixedIndex)
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("\ngot: \n%v\nwant: \n%v", err, tt.wantErr)
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

func TestFourSpacesStrategy_Generate(t *testing.T) {
	tests := map[string]struct {
		row     string
		want    *want
		wantErr error
	}{
		"root/hierarchy=1":                     {"- aaa bb", &want{name: "aaa bb", hierarchy: 1, index: fixedIndex}, nil},
		"child/hierarchy=2":                    {"    - aaa bb", &want{name: "aaa bb", hierarchy: 2, index: fixedIndex}, nil},
		"child/hierarchy=3":                    {"        - aaa    bb", &want{name: "aaa    bb", hierarchy: 3, index: fixedIndex}, nil},
		"invalid/hierarchy=0/prefix odd space": {" - aaa bb", nil, errIncorrectFormat},
		"invalid/hierarchy=0/prefix chars":     {"xx- aaa bb", nil, errIncorrectFormat},
		"invalid/hierarchy=0/no hyphen":        {"xx aaa bb", nil, errIncorrectFormat},
		"invalid/hierarchy=0/space only":       {"    ", nil, errIncorrectFormat},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			node, err := (*fourSpacesStrategy)(nil).generate(tt.row, fixedIndex)
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("\ngot: \n%v\nwant: \n%v", err, tt.wantErr)
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
