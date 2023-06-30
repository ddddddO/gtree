package gtree_test

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
	tu "github.com/ddddddO/gtree/testutil"
)

func Example() {
	var root *gtree.Node = gtree.NewRoot("root")
	root.Add("child 1").Add("child 2")
	root.Add("child 1").Add("child 3")
	child4 := root.Add("child 4")

	var child7 *gtree.Node = child4.Add("child 5").Add("child 6").Add("child 7")
	child7.Add("child 8")

	buf := &bytes.Buffer{}
	if err := gtree.OutputProgrammably(buf, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(buf.String())
	// Output:
	// root
	// ├── child 1
	// │   ├── child 2
	// │   └── child 3
	// └── child 4
	//     └── child 5
	//         └── child 6
	//             └── child 7
	//                 └── child 8
}

func TestOutputProgrammably(t *testing.T) {
	tests := []struct {
		name    string
		root    *gtree.Node
		options []gtree.Option
		want    string
		wantErr error
	}{
		{
			name: "case(succeeded)",
			root: tu.Prepare(),
			want: strings.TrimPrefix(`
root
└── child 1
    └── child 2
`, "\n"),
			wantErr: nil,
		},
		{
			name:    "case(succeeded/massive)",
			root:    tu.Prepare(),
			options: []gtree.Option{gtree.WithMassive(context.Background())},
			want: strings.TrimPrefix(`
root
└── child 1
    └── child 2
`, "\n"),
			wantErr: nil,
		},
		{
			name: "case(succeeded / added same name)",
			root: tu.PrepareSameNameChild(),
			want: strings.TrimPrefix(`
root
└── child 1
    ├── child 2
    └── child 3
`, "\n"),
			wantErr: nil,
		},
		{
			name:    "case(not root)",
			root:    tu.PrepareNotRoot(),
			want:    "",
			wantErr: gtree.ErrNotRoot,
		},
		{
			name:    "case(nil node)",
			root:    tu.PrepareNilNode(),
			want:    "",
			wantErr: gtree.ErrNilNode,
		},
		{
			name: "case(succeeded / branch format)",
			root: tu.PrepareMultiNode(),
			options: []gtree.Option{
				gtree.WithBranchFormatIntermedialNode("+--", ":   "),
				gtree.WithBranchFormatLastNode("+--", "    "),
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
			name:    "case(succeeded / output json)",
			root:    tu.PrepareMultiNode(),
			options: []gtree.Option{gtree.WithEncodeJSON()},
			want: strings.TrimPrefix(`
{"value":"root1","children":[{"value":"child 1","children":[{"value":"child 2","children":[{"value":"child 3","children":null},{"value":"child 4","children":[{"value":"child 5","children":null},{"value":"child 6","children":[{"value":"child 7","children":null}]}]}]}]},{"value":"child 8","children":null}]}
`, "\n"),
		},
		{
			name:    "case(succeeded / output yaml)",
			root:    tu.PrepareMultiNode(),
			options: []gtree.Option{gtree.WithEncodeYAML()},
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
			name:    "case(succeeded / output toml)",
			root:    tu.PrepareMultiNode(),
			options: []gtree.Option{gtree.WithEncodeTOML()},
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
			gotErr := gtree.OutputProgrammably(buf, tt.root, tt.options...)
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
