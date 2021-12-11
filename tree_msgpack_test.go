package gtree

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

func TestExecute_encodeMsgPack(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case 1(tab spaces & multi root & output message pack)",
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
				optFns: []OptFn{EncodeMsgPack()},
			},
			out: out{
				output: strings.TrimPrefix(`
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case 2(indent 2spaces & output message pack)",
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
				optFns: []OptFn{IndentTwoSpaces(), EncodeMsgPack()},
			},
			out: out{
				output: strings.TrimPrefix(`
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case 3(indent 4spaces & output message pack)",
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
				optFns: []OptFn{IndentFourSpaces(), EncodeMsgPack()},
			},
			out: out{
				output: strings.TrimPrefix(`
`, "\n"),
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := &bytes.Buffer{}
			gotErr := Execute(tmp, tt.in.input, tt.in.optFns...)
			out, err := msgpack.NewDecoder(tmp).DecodeMap()
			if err != nil {
				t.Fatal(err)
			}
			gotOutput := fmt.Sprintf("%v", out)

			if gotOutput != tt.out.output {
				t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.out.output)
			}
			if gotErr != tt.out.err {
				t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.out.err)
			}
		})
	}
}
