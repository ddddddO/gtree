package gtree_test

import (
	"fmt"
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
			name:    "case(not dry run/invalid node name)",
			root:    prepareInvalidNodeName(),
			wantErr: fmt.Errorf("invalid node name: %s", "chi/ld 4"),
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
