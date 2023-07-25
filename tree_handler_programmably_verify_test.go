package gtree_test

import (
	"testing"

	"github.com/ddddddO/gtree"
)

func TestVerifyProgrammably(t *testing.T) {
	tests := []struct {
		name    string
		root    *gtree.Node
		options []gtree.Option
		wantErr error
	}{
		{
			name:    "case(succeeded/massive)",
			root:    prepareDirectoryRoot(),
			options: []gtree.Option{gtree.WithMassive(nil)},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := gtree.VerifyProgrammably(tt.root, tt.options...)
			if gotErr != tt.wantErr {
				t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
			}
		})
	}
}

func prepareDirectoryRoot() *gtree.Node {
	root := gtree.NewRoot("example")
	root.Add("find_pipe_programmable-gtree").Add("main.go")
	root.Add("go-list_pipe_programmable-gtree").Add("main.go")
	likeCLI := root.Add("like_cli")
	adapter := likeCLI.Add("adapter")
	adapter.Add("executor.go")
	adapter.Add("indentation.go")
	likeCLI.Add("main.go")
	root.Add("programmable").Add("main.go")
	return root
}
