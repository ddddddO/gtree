package gtree

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestExecute_jsonTree(t *testing.T) {
	tests := []struct {
		name    string
		in      io.Reader
		want    string
		wantErr error
	}{
		{
			name: "case 1",
			in: strings.NewReader(strings.TrimSpace(`
- a
	- b`)),
			want:    `{"a":["b"]}`,
			wantErr: nil,
		},
		{
			name: "case 2",
			in: strings.NewReader(strings.TrimSpace(`
- a
	- b
	- c`)),
			want:    `{"a":["b","c"]}`,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}

			gotErr := Execute(out, tt.in, ModeJSON())
			gotOutput := out.String()

			if gotOutput != tt.want {
				t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.want)
			}
			if gotErr != tt.wantErr {
				t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
			}
		})
	}
}
