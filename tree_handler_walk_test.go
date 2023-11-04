package gtree_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

// TODO: 何パターンかのcallbackを用意してWalkerNode用メソッドのテストもしたい
func TestWalk(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case(succeeded)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
	- kk
		- t
- e
	- o
		- g`)),
			},
			out: out{
				output: strings.TrimLeft(`
a
├── i
│   └── u
│       └── k
└── kk
    └── t
e
└── o
    └── g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/change branch)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
	- kk
		- t`)),
				options: []gtree.Option{gtree.WithBranchFormatIntermedialNode("+--", ":   "), gtree.WithBranchFormatLastNode("+--", "    ")},
			},
			out: out{
				output: strings.TrimLeft(`
a
+-- i
:   +-- u
:       +-- k
+-- kk
    +-- t
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
			gotErr := gtree.Walk(tt.in.input, callback, tt.in.options...)
			if gotErr != nil || tt.out.err != nil {
				if gotErr.Error() != tt.out.err.Error() {
					t.Errorf("\ngotErr: \n%s\nwantErr: \n%s", gotErr, tt.out.err)
				}
			}
			got := buf.String()
			if got != tt.out.output {
				t.Errorf("\ngot: \n%s\nwant: \n%s", got, tt.out.output)
			}
		})
	}
}
