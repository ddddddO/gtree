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
	Children : [i kk]
	Ancestors: []
WalkerNode's methods called...
	Name     : i
	Branch   : ├──
	Row      : ├── i
	Level    : 2
	Path     : a/i
	HasChild : true
	Children : [u]
	Ancestors: [a]
WalkerNode's methods called...
	Name     : u
	Branch   : │   └──
	Row      : │   └── u
	Level    : 3
	Path     : a/i/u
	HasChild : true
	Children : [k]
	Ancestors: [a i]
WalkerNode's methods called...
	Name     : k
	Branch   : │       └──
	Row      : │       └── k
	Level    : 4
	Path     : a/i/u/k
	HasChild : false
	Children : []
	Ancestors: [a i u]
WalkerNode's methods called...
	Name     : kk
	Branch   : └──
	Row      : └── kk
	Level    : 2
	Path     : a/kk
	HasChild : true
	Children : [t]
	Ancestors: [a]
WalkerNode's methods called...
	Name     : t
	Branch   :     └──
	Row      :     └── t
	Level    : 3
	Path     : a/kk/t
	HasChild : false
	Children : []
	Ancestors: [a kk]
WalkerNode's methods called...
	Name     : e
	Branch   : 
	Row      : e
	Level    : 1
	Path     : e
	HasChild : true
	Children : [o]
	Ancestors: []
WalkerNode's methods called...
	Name     : o
	Branch   : └──
	Row      : └── o
	Level    : 2
	Path     : e/o
	HasChild : true
	Children : [g]
	Ancestors: [e]
WalkerNode's methods called...
	Name     : g
	Branch   :     └──
	Row      :     └── g
	Level    : 3
	Path     : e/o/g
	HasChild : false
	Children : []
	Ancestors: [e o]
`, "\n"),
				err: nil,
			},
		},
	}

	for _, tt := range tests {
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

				childNames := func(wn *gtree.WalkerNode) []string {
					names := make([]string, 0, len(wn.Children()))
					for _, child := range wn.Children() {
						names = append(names, child.Name())
					}
					return names
				}(wn)
				fmt.Fprintf(buf, "\tChildren : %v\n", childNames)

				ancestorNames := func(wn *gtree.WalkerNode) []string {
					names := []string{}
					for _, parent := range wn.Ancestors() {
						names = append(names, parent.Name())
					}
					return names
				}(wn)
				fmt.Fprintf(buf, "\tAncestors: %v\n", ancestorNames)

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
