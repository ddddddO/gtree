package gtree_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

func TestWalkFromMarkdown(t *testing.T) {
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
				options: []gtree.Option{
					gtree.WithMidBranch("+--"),
					gtree.WithLastBranch("+--"),
					gtree.WithHLine(""),
					gtree.WithVLine(":"),
				},
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
			gotErr := gtree.WalkFromMarkdown(tt.in.input, callback, tt.in.options...)
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

func TestWalk_WalkerNode(t *testing.T) {
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
WalkerNode's methods called...
	Name     : a
	Branch   : 
	Row      : a
	Level    : 1
	Path     : a
	HasChild : true
WalkerNode's methods called...
	Name     : i
	Branch   : ├──
	Row      : ├── i
	Level    : 2
	Path     : a/i
	HasChild : true
WalkerNode's methods called...
	Name     : u
	Branch   : │   └──
	Row      : │   └── u
	Level    : 3
	Path     : a/i/u
	HasChild : true
WalkerNode's methods called...
	Name     : k
	Branch   : │       └──
	Row      : │       └── k
	Level    : 4
	Path     : a/i/u/k
	HasChild : false
WalkerNode's methods called...
	Name     : kk
	Branch   : └──
	Row      : └── kk
	Level    : 2
	Path     : a/kk
	HasChild : true
WalkerNode's methods called...
	Name     : t
	Branch   :     └──
	Row      :     └── t
	Level    : 3
	Path     : a/kk/t
	HasChild : false
WalkerNode's methods called...
	Name     : e
	Branch   : 
	Row      : e
	Level    : 1
	Path     : e
	HasChild : true
WalkerNode's methods called...
	Name     : o
	Branch   : └──
	Row      : └── o
	Level    : 2
	Path     : e/o
	HasChild : true
WalkerNode's methods called...
	Name     : g
	Branch   :     └──
	Row      :     └── g
	Level    : 3
	Path     : e/o/g
	HasChild : false
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
				fmt.Fprintln(buf, "WalkerNode's methods called...")
				fmt.Fprintf(buf, "\tName     : %s\n", wn.Name())
				fmt.Fprintf(buf, "\tBranch   : %s\n", wn.Branch())
				fmt.Fprintf(buf, "\tRow      : %s\n", wn.Row())
				fmt.Fprintf(buf, "\tLevel    : %d\n", wn.Level())
				fmt.Fprintf(buf, "\tPath     : %s\n", wn.Path())
				fmt.Fprintf(buf, "\tHasChild : %t\n", wn.HasChild())
				return nil
			}
			gotErr := gtree.WalkFromMarkdown(tt.in.input, callback, tt.in.options...)
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
