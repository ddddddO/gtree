package gtree_test

import (
	"fmt"
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
		{
			name:    "case(succeeded/specify target dir/strict mode)",
			root:    prepareDirectoryFindPipe(),
			options: []gtree.Option{gtree.WithTargetDir("example"), gtree.WithStrictVerify(), gtree.WithMassive(nil)},
			wantErr: nil,
		},
		{
			name:    "case(error/specify target dir)",
			root:    prepareDirectoryFindPipeWithNoExistDirAndRequiredDir(),
			options: []gtree.Option{gtree.WithTargetDir("example"), gtree.WithMassive(nil)},
			wantErr: fmt.Errorf("Required paths does not exist:\n%s", "\texample/find_pipe_programmable-gtree/required!"),
		},
		{
			name:    "case(error/specify target dir/strict mode)",
			root:    prepareDirectoryFindPipeWithNoExistDirAndRequiredDir(),
			options: []gtree.Option{gtree.WithTargetDir("example"), gtree.WithStrictVerify(), gtree.WithMassive(nil)},
			wantErr: fmt.Errorf("Extra paths exist:\n%s\nRequired paths does not exist:\n%s", "\texample/find_pipe_programmable-gtree/main.go", "\texample/find_pipe_programmable-gtree/required!"),
		},
		{
			name:    "case(error/no exist root)",
			root:    gtree.NewRoot("no_exist_root_dir"),
			options: []gtree.Option{gtree.WithMassive(nil)},
			wantErr: fmt.Errorf("Required paths does not exist:\n%s", "\tno_exist_root_dir"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := gtree.VerifyProgrammably(tt.root, tt.options...)
			if gotErr != nil || tt.wantErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("\ngotErr: \n%s\nwantErr: \n%s", gotErr, tt.wantErr)
				}
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

func prepareDirectoryFindPipe() *gtree.Node {
	root := gtree.NewRoot("find_pipe_programmable-gtree")
	root.Add("go.mod")
	root.Add("go.sum")
	root.Add("main.go")
	root.Add("README.md")
	return root
}

func prepareDirectoryFindPipeWithNoExistDirAndRequiredDir() *gtree.Node {
	root := gtree.NewRoot("find_pipe_programmable-gtree")
	root.Add("go.mod")
	root.Add("go.sum")
	root.Add("required!")
	root.Add("README.md")
	return root
}
