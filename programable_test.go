package gtree

import (
	"bytes"
	"strings"
	"testing"
)

func TestExecuteProgrammably(t *testing.T) {
	tests := []struct {
		name    string
		root    *Node
		want    string
		wantErr error
	}{
		{
			name: "case1(succeeded)",
			root: prepare(),
			want: strings.TrimPrefix(`
root
└── child 1
    └── child 2
`, "\n"),
			wantErr: nil,
		},
		{
			name:    "case2(not root)",
			root:    prepareNotRoot(),
			want:    "",
			wantErr: ErrNotRoot,
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)

		buf := &bytes.Buffer{}
		gotErr := ExecuteProgrammably(tt.root, buf)
		got := buf.String()

		if got != tt.want {
			t.Errorf("\ngot: \n%s\nwant: \n%s", got, tt.want)
		}
		if gotErr != tt.wantErr {
			t.Errorf("\ngot: \n%v\nwant: \n%v", gotErr, tt.wantErr)
		}
	}
}

func prepare() *Node {
	root := NewRoot("root")
	root.Add("child 1").Add("child 2")
	return root
}

func prepareNotRoot() *Node {
	root := NewRoot("root")
	child1 := root.Add("child 1")
	return child1
}
