package gtree_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

func TestVerify(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case(succeeded)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- example
	- find_pipe_programmable-gtree
		- main.go
	- go-list_pipe_programmable-gtree
		- main.go
	- like_cli
		- adapter
			- executor.go
			- indentation.go
		- main.go
	- programmable
		- main.go`)),
			},
			out: out{
				err: nil,
			},
		},
		{
			name: "case(error/no exist root)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- no_exist_root_dir`)),
			},
			out: out{
				err: fmt.Errorf("Required paths does not exist:\n%s", "\tno_exist_root_dir"),
			},
		},
		{
			name: "case(error/strict mode/Extra paths exist and Required paths does not exist)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- example
	- find_pipe_programmable-gtree
		- main.go
		- README.md
		- go.mod
		- go.sum
	- go-list_pipe_programmable-gtree
		- main.go
		- README.md
		- go.mod
		- go.sum
	- like_cli
		- adapter
			- executor.go
			- indentation.go
		- main.go
		- xxxx
	- programmable
		- main.go
	- README.md`)),
				options: []gtree.Option{gtree.WithStrictVerify()},
			},
			out: out{
				err: fmt.Errorf("Extra paths exist:\n%s\nRequired paths does not exist:\n%s", "\texample/noexist\n\texample/noexist/xxx", "\texample/like_cli/xxxx"),
			},
		},
		{
			name: "case(error/strict mode/specify target dir/Extra paths exist and Required paths does not exist)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- go-list_pipe_programmable-gtree
	- main.go
	- README.md
	- go.sum
	- xxxx`)),
				options: []gtree.Option{gtree.WithStrictVerify(), gtree.WithTargetDir("example")},
			},
			out: out{
				err: fmt.Errorf("Extra paths exist:\n%s\nRequired paths does not exist:\n%s", "\texample/go-list_pipe_programmable-gtree/go.mod", "\texample/go-list_pipe_programmable-gtree/xxxx"),
			},
		},
		{
			name: "case(error/specify target dir/Required paths does not exist)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- go-list_pipe_programmable-gtree
	- main.go
	- README.md
	- go.sum
	- xxxx`)),
				options: []gtree.Option{gtree.WithTargetDir("example")},
			},
			out: out{
				err: fmt.Errorf("Required paths does not exist:\n%s", "\texample/go-list_pipe_programmable-gtree/xxxx"),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := gtree.Verify(tt.in.input, tt.in.options...)
			if gotErr != nil || tt.out.err != nil {
				if gotErr.Error() != tt.out.err.Error() {
					t.Errorf("\ngotErr: \n%s\nwantErr: \n%s", gotErr, tt.out.err)
				}
			}
		})
	}
}
