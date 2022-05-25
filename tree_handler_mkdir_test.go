package gtree_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

func TestMkdir(t *testing.T) {
	tests := []struct {
		name    string
		in      in
		wantErr error
	}{
		{
			name: "case(succeeded)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root2
	- b
	- bb
		- lll
	-ff`)),
			},
			wantErr: nil,
		},
		{
			name: "case(dry-run/no error)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root2
	- b
	- bb
		- lll
	-ff`)),
				options: []gtree.Option{gtree.WithDryRun()},
			},
			wantErr: nil,
		},
		{
			name: "case(dry-run/invalid node name)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root2
	- b
	- b/b
		- lll
	-ff`)),
				options: []gtree.Option{gtree.WithDryRun()},
			},
			wantErr: errors.New("invalid node name: b/b"),
		},
		{
			name: "case(dry-run/invalid path)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- /root2
	- b
	- bb
		- lll
	-ff`)),
				options: []gtree.Option{gtree.WithDryRun()},
			},
			wantErr: errors.New("invalid path: /root2/b"),
		},
		{
			name: "case(succeeded/only root)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root3`)),
			},
			wantErr: nil,
		},
		{
			name: "case(succeeded/only multi roots)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root4
- root5`)),
			},
			wantErr: nil,
		},
		{
			name: "case(succeeded/make directories and files)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root6
	- b.go
	- bb
		- lll
	-makefile`)),
				options: []gtree.Option{gtree.WithFileExtensions([]string{".go", "makefile"})},
			},
			wantErr: nil,
		},
		{
			name: "case(succeeded/make directories and files/even if the extension is specified, it must be created as a directory)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root7
	- b.go
	- bb.go
		- lll`)),
				options: []gtree.Option{gtree.WithFileExtensions([]string{".go"})},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := gtree.Mkdir(tt.in.input, tt.in.options...)
			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
				}
			}
		})
	}
}
