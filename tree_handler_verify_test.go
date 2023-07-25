package gtree_test

import (
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
			name: "case(error/strict mode/Extra paths exist and Required paths does not exist)",
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
		- xxxx
	- programmable
		- main.go`)),
				options: []gtree.Option{gtree.WithStrictVerify()},
			},
			out: out{
				err: nil, // TODO: not nil
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := gtree.Verify(tt.in.input, tt.in.options...)
			if gotErr != nil {
				// TODO:
				// if gotErr.Error() != tt.out.err.Error() {
				// 	t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.out.err)
				// }
			}
		})
	}
}
