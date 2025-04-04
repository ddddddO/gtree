package gtree_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

func TestMkdirFromMarkdown(t *testing.T) {
	tests := []struct {
		name    string
		in      in
		wantErr error
	}{
		{
			name: "case(succeeded)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root_a
	- b
	- bb
		- lll
	- ff`)),
			},
			wantErr: nil,
		},
		{
			name: "case(dry-run/path exist err)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- example
	- b
	- bb
		- lll
	- ff`)),
				options: []gtree.Option{gtree.WithDryRun()},
			},
			wantErr: gtree.ErrExistPath,
		},
		{
			name: "case(massive/path exist err)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- example
	- b
	- bb
		- lll
	- ff`)),
				options: []gtree.Option{gtree.WithMassive(context.Background())},
			},
			wantErr: gtree.ErrExistPath,
		},
		{
			name: "case(dry-run/invalid node name)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root_c
	- b
	- b/b
		- lll
	- ff`)),
				options: []gtree.Option{gtree.WithDryRun()},
			},
			wantErr: errors.New("invalid node name: b/b"),
		},
		// NOTE: 上のパターンでエラーとして返すようになった
		// 		{
		// 			name: "case(dry-run/invalid path)",
		// 			in: in{
		// 				input: strings.NewReader(strings.TrimSpace(`
		// - /root_d
		// 	- b
		// 	- bb
		// 		- lll
		// 	-ff`)),
		// 				options: []gtree.Option{gtree.WithDryRun()},
		// 			},
		// 			wantErr: errors.New("invalid path: /root2/b"),
		// 		},
		{
			name: "case(succeeded/only root)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root_e`)),
			},
			wantErr: nil,
		},
		{
			name: "case(succeeded/only multi roots)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root_f
- root_g`)),
			},
			wantErr: nil,
		},
		{
			name: "case(succeeded/make directories and files)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root_h
	- b.go
	- bb
		- lll
	- Makefile`)),
				options: []gtree.Option{gtree.WithFileExtensions([]string{".go", "Makefile"})},
			},
			wantErr: nil,
		},
		{
			name: "case(succeeded/make directories and files/even if the extension is specified, it must be created as a directory)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root_i
	- b.go
	- bb.go
		- lll`)),
				options: []gtree.Option{gtree.WithFileExtensions([]string{".go"})},
			},
			wantErr: nil,
		},
		{
			name: "case(succeeded/make directories and files/massive root)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root_j
	- b.go
	- bb.go
		- lll`)),
				options: []gtree.Option{
					gtree.WithFileExtensions([]string{".go"}),
					gtree.WithMassive(context.Background()),
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := gtree.MkdirFromMarkdown(tt.in.input, tt.in.options...)
			if gotErr != nil {
				t.Log(gotErr.Error())
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
				}
			}
		})
	}
}
