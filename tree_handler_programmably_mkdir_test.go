package gtree_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ddddddO/gtree"
	tu "github.com/ddddddO/gtree/testutil"
)

func TestMkdirFromRoot(t *testing.T) {
	tests := []struct {
		name    string
		root    *gtree.Node
		options []gtree.Option
		wantErr error
	}{
		{
			name: "case(succeeded)",
			root: tu.Prepare(),
		},
		{
			name:    "case(succeeded/massive)",
			root:    tu.Prepare_a(),
			options: []gtree.Option{gtree.WithMassive(context.Background())},
		},
		{
			name:    "case(not root)",
			root:    tu.PrepareNotRoot(),
			wantErr: gtree.ErrNotRoot,
		},
		{
			name:    "case(nil node)",
			root:    tu.PrepareNilNode(),
			wantErr: gtree.ErrNilNode,
		},
		{
			name: "case(succeeded)",
			root: tu.PrepareMultiNode(),
		},
		{
			name: "case(dry run/invalid node name)",
			root: tu.PrepareInvalidNodeName(),
			options: []gtree.Option{
				gtree.WithDryRun(),
			},
			wantErr: fmt.Errorf("invalid node name: %s", "chi/ld 4"),
		},
		{
			name: "case(dry run/succeeded)",
			root: tu.PrepareMultiNode(),
			options: []gtree.Option{
				gtree.WithDryRun(),
			},
			wantErr: nil,
		},
		{
			name: "case(dry run/specified file/succeeded)",
			root: tu.PrepareMultiNode(),
			options: []gtree.Option{
				gtree.WithDryRun(),
				gtree.WithFileExtensions([]string{"child 3", "child 5", "child 7", "child 8"}),
			},
			wantErr: nil,
		},
		{
			name:    "case(not dry run/invalid node name)",
			root:    tu.PrepareInvalidNodeName(),
			wantErr: fmt.Errorf("invalid node name: %s", "chi/ld 4"),
		},
		{
			name:    "case(root already exists)",
			root:    tu.PrepareExistRoot(t),
			wantErr: gtree.ErrExistPath,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := gtree.MkdirFromRoot(tt.root, tt.options...)
			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
				}
			}
		})
	}
}
