package gtree_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
	tu "github.com/ddddddO/gtree/testutil"
)

// TODO: 何パターンかのcallbackを用意してWalkerNode用メソッドのテストもしたい
func TestWalkProgrammably(t *testing.T) {
	tests := []struct {
		name    string
		root    *gtree.Node
		options []gtree.Option
		out     out
	}{
		{
			name: "case(succeeded)",
			root: tu.PrepareMultiNode(),
			options: []gtree.Option{
				gtree.WithBranchFormatIntermedialNode("+--", ":   "),
				gtree.WithBranchFormatLastNode("+--", "    "),
			},
			out: out{
				output: strings.TrimLeft(`
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
				err: nil,
			},
		},
		{
			name:    "case(succeeded/massive)",
			root:    tu.Prepare(),
			options: []gtree.Option{gtree.WithMassive(nil)},
			out: out{
				output: strings.TrimLeft(`
root
└── child 1
    └── child 2
`, "\n"),
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			buf := &bytes.Buffer{}
			callback := func(wn *gtree.WalkerNode) error {
				fmt.Fprintln(buf, wn.Row())
				return nil
			}
			gotErr := gtree.WalkProgrammably(tt.root, callback, tt.options...)
			if gotErr != nil || tt.out.err != nil {
				if gotErr.Error() != tt.out.err.Error() {
					t.Errorf("\ngotErr: \n%s\nwantErr: \n%s", gotErr, tt.out.err.Error())
				}
			}
			got := buf.String()
			if got != tt.out.output {
				t.Errorf("\ngot: \n%s\nwant: \n%s", got, tt.out.output)
			}
		})
	}
}

func TestWalkIterProgrammably(t *testing.T) {
	tests := []struct {
		name    string
		root    *gtree.Node
		options []gtree.Option
		out     out
	}{
		{
			name: "case(succeeded)",
			root: tu.PrepareMultiNode(),
			options: []gtree.Option{
				gtree.WithBranchFormatIntermedialNode("+--", ":   "),
				gtree.WithBranchFormatLastNode("+--", "    "),
			},
			out: out{
				output: strings.TrimLeft(`
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
				err: nil,
			},
		},
		{
			name:    "case(succeeded/massive)",
			root:    tu.Prepare(),
			options: []gtree.Option{gtree.WithMassive(nil)},
			out: out{
				output: strings.TrimLeft(`
root
└── child 1
    └── child 2
`, "\n"),
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ""
			for walkerNode, gotErr := range gtree.WalkIterProgrammably(tt.root, tt.options...) {
				if gotErr != nil && gotErr.Error() != tt.out.err.Error() {
					t.Errorf("\ngotErr: \n%s\nwantErr: \n%s", gotErr, tt.out.err.Error())
				}
				got += walkerNode.Row() + "\n"
			}
			if got != tt.out.output {
				t.Errorf("\ngot: \n%s\nwant: \n%s", got, tt.out.output)
			}
		})
	}
}
