package gtree

import (
	"bytes"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestOutputProgrammably(t *testing.T) {
	tests := []struct {
		name    string
		root    *Node
		optFns  []OptFn
		want    string
		wantErr error
	}{
		{
			name: "case1(succeeded)",
			root: prepare(),
			want: strings.TrimPrefix(`
root
└── child 1
    └── child 2
`, "\n"),
			wantErr: nil,
		},
		{
			name: "case2(succeeded / added same name)",
			root: prepareSameNameChild(),
			want: strings.TrimPrefix(`
root
└── child 1
    ├── child 2
    └── child 3
`, "\n"),
			wantErr: nil,
		},
		{
			name:    "case3(not root)",
			root:    prepareNotRoot(),
			want:    "",
			wantErr: ErrNotRoot,
		},
		{
			name:    "case4(nil node)",
			root:    prepareNilNode(),
			want:    "",
			wantErr: ErrNilNode,
		},
		{
			name: "case5(succeeded / branch format)",
			root: prepareMultiNode(),
			optFns: []OptFn{
				WithBranchFormatIntermedialNode("+--", ":   "),
				WithBranchFormatLastNode("+--", "    "),
			},
			want: strings.TrimPrefix(`
root1
+-- child 1
:   +-- child 2
:       +-- child 3
:       +-- child 4
:           +-- child 5
:           +-- child 6
:               +-- child 7
+-- child 8
`, "\n"),
		},
		{
			name:   "case6(succeeded / output json)",
			root:   prepareMultiNode(),
			optFns: []OptFn{WithEncodeJSON()},
			want: strings.TrimPrefix(`
{"value":"root1","children":[{"value":"child 1","children":[{"value":"child 2","children":[{"value":"child 3","children":null},{"value":"child 4","children":[{"value":"child 5","children":null},{"value":"child 6","children":[{"value":"child 7","children":null}]}]}]}]},{"value":"child 8","children":null}]}
`, "\n"),
		},
		{
			name:   "case7(succeeded / output yaml)",
			root:   prepareMultiNode(),
			optFns: []OptFn{WithEncodeYAML()},
			want: strings.TrimPrefix(`
value: root1
children:
- value: child 1
  children:
  - value: child 2
    children:
    - value: child 3
      children: []
    - value: child 4
      children:
      - value: child 5
        children: []
      - value: child 6
        children:
        - value: child 7
          children: []
- value: child 8
  children: []
`, "\n"),
		},
		{
			name:   "case8(succeeded / output toml)",
			root:   prepareMultiNode(),
			optFns: []OptFn{WithEncodeTOML()},
			want: strings.TrimPrefix(`
value = 'root1'
[[children]]
value = 'child 1'
[[children.children]]
value = 'child 2'
[[children.children.children]]
value = 'child 3'
children = []
[[children.children.children]]
value = 'child 4'
[[children.children.children.children]]
value = 'child 5'
children = []
[[children.children.children.children]]
value = 'child 6'
[[children.children.children.children.children]]
value = 'child 7'
children = []




[[children]]
value = 'child 8'
children = []

`, "\n"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			buf := &bytes.Buffer{}
			gotErr := OutputProgrammably(buf, tt.root, tt.optFns...)
			got := buf.String()

			if got != tt.want {
				t.Errorf("\ngot: \n%s\nwant: \n%s", got, tt.want)
			}
			if gotErr != tt.wantErr {
				t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
			}
		})
	}
}

func prepare() *Node {
	root := NewRoot("root")
	root.Add("child 1").Add("child 2")
	return root
}

func prepareSameNameChild() *Node {
	root := NewRoot("root")
	root.Add("child 1").Add("child 2")
	root.Add("child 1").Add("child 3")
	return root
}

func prepareNotRoot() *Node {
	root := NewRoot("root")
	child1 := root.Add("child 1")
	return child1
}

func prepareNilNode() *Node {
	var node *Node
	return node
}

func prepareMultiNode() *Node {
	var root *Node = NewRoot("root1")
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *Node = root.Add("child 1").Add("child 2").Add("child 4")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	root.Add("child 8")
	return root
}

func prepareInvalidNodeName() *Node {
	var root *Node = NewRoot("root1")
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *Node = root.Add("child 1").Add("child 2").Add("chi/ld 4")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	root.Add("child 8")
	return root
}

func TestMkdirProgrammably(t *testing.T) {
	tests := []struct {
		name    string
		root    *Node
		optFns  []OptFn
		wantErr error
	}{
		{
			name: "case1(succeeded)",
			root: prepare(),
		},
		{
			name:    "case2(not root)",
			root:    prepareNotRoot(),
			wantErr: ErrNotRoot,
		},
		{
			name:    "case3(nil node)",
			root:    prepareNilNode(),
			wantErr: ErrNilNode,
		},
		{
			name: "case4(succeeded)",
			root: prepareMultiNode(),
		},
		{
			name: "case5(dry run/invalid node name)",
			root: prepareInvalidNodeName(),
			optFns: []OptFn{
				WithDryRun(),
			},
			wantErr: errors.Errorf("invalid node name: %s", "chi/ld 4"),
		},
		{
			name: "case6(dry run/succeeded)",
			root: prepareMultiNode(),
			optFns: []OptFn{
				WithDryRun(),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := MkdirProgrammably(tt.root, tt.optFns...)
			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
				}
			}
		})
	}
}
