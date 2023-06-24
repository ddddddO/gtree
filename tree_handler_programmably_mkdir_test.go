package gtree_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ddddddO/gtree"
)

func TestMkdirProgrammably(t *testing.T) {
	tests := []struct {
		name    string
		root    *gtree.Node
		options []gtree.Option
		wantErr error
	}{
		{
			name: "case(succeeded)",
			root: prepare(),
		},
		{
			name:    "case(succeeded/massive)",
			root:    prepare_a(),
			options: []gtree.Option{gtree.WithMassive()},
		},
		{
			name:    "case(not root)",
			root:    prepareNotRoot(),
			wantErr: gtree.ErrNotRoot,
		},
		{
			name:    "case(nil node)",
			root:    prepareNilNode(),
			wantErr: gtree.ErrNilNode,
		},
		{
			name: "case(succeeded)",
			root: prepareMultiNode(),
		},
		{
			name: "case(dry run/invalid node name)",
			root: prepareInvalidNodeName(),
			options: []gtree.Option{
				gtree.WithDryRun(),
			},
			wantErr: fmt.Errorf("invalid node name: %s", "chi/ld 4"),
		},
		{
			name: "case(dry run/succeeded)",
			root: prepareMultiNode(),
			options: []gtree.Option{
				gtree.WithDryRun(),
			},
			wantErr: nil,
		},
		{
			name: "case(dry run/specified file/succeeded)",
			root: prepareMultiNode(),
			options: []gtree.Option{
				gtree.WithDryRun(),
				gtree.WithFileExtensions([]string{"child 3", "child 5", "child 7", "child 8"}),
			},
			wantErr: nil,
		},
		{
			name:    "case(not dry run/invalid node name)",
			root:    prepareInvalidNodeName(),
			wantErr: fmt.Errorf("invalid node name: %s", "chi/ld 4"),
		},
		{
			name:    "case(root already exists)",
			root:    prepareExistRoot(t),
			wantErr: gtree.ErrExistPath,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := gtree.MkdirProgrammably(tt.root, tt.options...)
			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
				}
			}
		})
	}
}

func prepareExistRoot(t *testing.T) *gtree.Node {
	name := "gtreetest"

	if err := os.MkdirAll(name, 0o755); err != nil {
		t.Fatal(err)
	}

	root := gtree.NewRoot(name)
	root.Add("temp")
	return root
}

func prepare_a() *gtree.Node {
	root := gtree.NewRoot("root8")
	root.Add("child 1").Add("child 2")
	return root
}
