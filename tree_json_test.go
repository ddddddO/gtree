package gtree

import (
	"bytes"
	"strings"
	"testing"
)

func TestOutput_encodeJSON(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case(tab spaces & multi root & output json)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`)),
				optFns: []OptFn{WithEncodeJSON()},
			},
			out: out{
				output: strings.TrimPrefix(`
{"value":"a","children":[{"value":"i","children":[{"value":"u","children":[{"value":"k","children":null},{"value":"kk","children":null}]},{"value":"t","children":null}]},{"value":"e","children":[{"value":"o","children":null}]},{"value":"g","children":null}]}
{"value":"a","children":[{"value":"i","children":[{"value":"u","children":[{"value":"k","children":null},{"value":"kk","children":null}]},{"value":"t","children":null}]},{"value":"e","children":[{"value":"o","children":null}]},{"value":"g","children":null}]}
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(indent 2spaces & output json)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
  - i
    - u
      - k
      - kk
    - t
  - e
    - o
  - g`)),
				optFns: []OptFn{WithIndentTwoSpaces(), WithEncodeJSON()},
			},
			out: out{
				output: strings.TrimPrefix(`
{"value":"a","children":[{"value":"i","children":[{"value":"u","children":[{"value":"k","children":null},{"value":"kk","children":null}]},{"value":"t","children":null}]},{"value":"e","children":[{"value":"o","children":null}]},{"value":"g","children":null}]}
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(indent 4spaces & output json)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
    - i
        - u
            - k
            - kk
        - t
    - e
        - o
    - g`)),
				optFns: []OptFn{WithIndentFourSpaces(), WithEncodeJSON()},
			},
			out: out{
				output: strings.TrimPrefix(`
{"value":"a","children":[{"value":"i","children":[{"value":"u","children":[{"value":"k","children":null},{"value":"kk","children":null}]},{"value":"t","children":null}]},{"value":"e","children":[{"value":"o","children":null}]},{"value":"g","children":null}]}
`, "\n"),
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := &bytes.Buffer{}
			gotErr := Output(out, tt.in.input, tt.in.optFns...)
			gotOutput := out.String()

			if gotOutput != tt.out.output {
				t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.out.output)
			}
			if gotErr != tt.out.err {
				t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.out.err)
			}
		})
	}
}
