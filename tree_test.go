package gtree

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

type in struct {
	input  io.Reader
	optFns []OptFn
}

type out struct {
	output string
	err    error
}

func TestOutput(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case(succeeded/has a child)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b`)),
			},
			out: out{
				output: strings.TrimPrefix(`
a
└── b
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/has a child nest)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c`))},

			out: out{
				output: strings.TrimPrefix(`
a
└── b
    └── c
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/has children)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
	- c`))},
			out: out{
				output: strings.TrimPrefix(`
a
├── b
└── c
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/has children deeply)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c
			- d
			- e
			- f`))},
			out: out{
				output: strings.TrimPrefix(`
a
└── b
    └── c
        ├── d
        ├── e
        └── f
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/has children complexly)",
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
	- g`))},
			out: out{
				output: strings.TrimPrefix(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/very deeply)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root
	- dddd
		- kkkkkkk
			- lllll
				- ffff
				- LLL
					- WWWWW
						- ZZZZZ
				- ppppp
					- KKK
						- 1111111
							- AAAAAAA
	- eee`))},
			out: out{
				output: strings.TrimPrefix(`
root
├── dddd
│   └── kkkkkkk
│       └── lllll
│           ├── ffff
│           ├── LLL
│           │   └── WWWWW
│           │       └── ZZZZZ
│           └── ppppp
│               └── KKK
│                   └── 1111111
│                       └── AAAAAAA
└── eee
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 2spaces)",
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
				optFns: []OptFn{WithIndentTwoSpaces()},
			},
			out: out{
				output: strings.TrimPrefix(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 4spaces)",
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
				optFns: []OptFn{WithIndentFourSpaces()},
			},
			out: out{
				output: strings.TrimPrefix(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/node name 1space & -)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root dir aaa
	- child-dir`))},
			out: out{
				output: strings.TrimPrefix(`
root dir aaa
└── child-dir
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/same node name)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- parent
	- child
		- chilchil
		- chilchil
		- chilchil
	- child`))},
			out: out{
				output: strings.TrimPrefix(`
parent
├── child
│   ├── chilchil
│   ├── chilchil
│   └── chilchil
└── child
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/input byte)",
			in: in{
				input: bytes.NewBufferString(strings.TrimSpace(`
- a
	- b`)),
			},
			out: out{
				output: strings.TrimPrefix(`
a
└── b
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/multi root)",
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
	- g`))},
			out: out{
				output: strings.TrimPrefix(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(empty node name)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	-`)),
			},
			out: out{
				output: "",
				err:    errEmptyText,
			},
		},
		{
			// TODO: inputのパターンが3つ(tab/ts/fs)と実行時のモードが3つで、それぞれの正常系(3つ)を上でしてるから、このパターン含めると、
			//       3*3-3=6パターンのケースが必要
			name: "case(incorrect input format(input 4spaces indent / tab mode))",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
    - b`)),
			},
			out: out{
				output: "",
				err:    errIncorrectFormat,
			},
		},
		{
			name: "case(bufio.Scanner err)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(fmt.Sprintf(`
- a
	- %s`, strings.Repeat("a", 64*1024)))),
			},
			out: out{
				output: "",
				err:    bufio.ErrTooLong,
			},
		},
		{
			name: "case(succeeded/input markdown file)",
			in: in{
				input: prepareMarkdownFile(t)},
			out: out{
				output: strings.TrimPrefix(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 2spaces and cutom branch format)",
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
				optFns: []OptFn{
					WithIndentTwoSpaces(),
					WithBranchFormatIntermedialNode("+->", ":   "),
					WithBranchFormatLastNode("+->", "    "),
				},
			},
			out: out{
				output: strings.TrimPrefix(`
a
+-> i
:   +-> u
:   :   +-> k
:   :   +-> kk
:   +-> t
+-> e
:   +-> o
+-> g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/dry run/no error)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b`)),
				optFns: []OptFn{
					WithDryRun(),
				},
			},
			out: out{
				output: strings.TrimPrefix(`
a
└── b
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(dry run/invalid node name)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b/c`)),
				optFns: []OptFn{
					WithDryRun(),
				},
			},
			out: out{
				output: "",
				err:    errors.New("invalid node name: b/c"),
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
			if gotErr != nil {
				if gotErr.Error() != tt.out.err.Error() {
					t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.out.err)
				}
			}

			if file, ok := tt.in.input.(*os.File); ok {
				file.Close()
			}
		})
	}
}

func prepareMarkdownFile(t *testing.T) *os.File {
	const testfilepath = "./testdata/sample6.md"
	file, err := os.Open(testfilepath)
	if err != nil {
		t.Fatal(err)
	}
	return file
}

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
				optFns: []OptFn{WithDryRun()},
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
				optFns: []OptFn{WithDryRun()},
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
				optFns: []OptFn{WithDryRun()},
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
				optFns: []OptFn{WithFileExtension([]string{".go", "makefile"})},
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
				optFns: []OptFn{WithFileExtension([]string{".go"})},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := Mkdir(tt.in.input, tt.in.optFns...)
			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
				}
			}
		})
	}
}

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

func TestOutput_encodeTOML(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case(succeeded/tab spaces & multi root & output toml)",
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
				optFns: []OptFn{WithEncodeTOML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value = 'a'
[[children]]
value = 'i'
[[children.children]]
value = 'u'
[[children.children.children]]
value = 'k'
children = []
[[children.children.children]]
value = 'kk'
children = []

[[children.children]]
value = 't'
children = []

[[children]]
value = 'e'
[[children.children]]
value = 'o'
children = []

[[children]]
value = 'g'
children = []

value = 'a'
[[children]]
value = 'i'
[[children.children]]
value = 'u'
[[children.children.children]]
value = 'k'
children = []
[[children.children.children]]
value = 'kk'
children = []

[[children.children]]
value = 't'
children = []

[[children]]
value = 'e'
[[children.children]]
value = 'o'
children = []

[[children]]
value = 'g'
children = []

`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 2spaces & output toml)",
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
				optFns: []OptFn{WithIndentTwoSpaces(), WithEncodeTOML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value = 'a'
[[children]]
value = 'i'
[[children.children]]
value = 'u'
[[children.children.children]]
value = 'k'
children = []
[[children.children.children]]
value = 'kk'
children = []

[[children.children]]
value = 't'
children = []

[[children]]
value = 'e'
[[children.children]]
value = 'o'
children = []

[[children]]
value = 'g'
children = []

`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 4spaces & output toml)",
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
				optFns: []OptFn{WithIndentFourSpaces(), WithEncodeTOML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value = 'a'
[[children]]
value = 'i'
[[children.children]]
value = 'u'
[[children.children.children]]
value = 'k'
children = []
[[children.children.children]]
value = 'kk'
children = []

[[children.children]]
value = 't'
children = []

[[children]]
value = 'e'
[[children.children]]
value = 'o'
children = []

[[children]]
value = 'g'
children = []

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

func TestOutput_encodeYAML(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case(succeeded/tab spaces & multi root & output yaml)",
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
				optFns: []OptFn{WithEncodeYAML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value: a
children:
- value: i
  children:
  - value: u
    children:
    - value: k
      children: []
    - value: kk
      children: []
  - value: t
    children: []
- value: e
  children:
  - value: o
    children: []
- value: g
  children: []
---
value: a
children:
- value: i
  children:
  - value: u
    children:
    - value: k
      children: []
    - value: kk
      children: []
  - value: t
    children: []
- value: e
  children:
  - value: o
    children: []
- value: g
  children: []
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 2spaces & output yaml)",
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
				optFns: []OptFn{WithIndentTwoSpaces(), WithEncodeYAML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value: a
children:
- value: i
  children:
  - value: u
    children:
    - value: k
      children: []
    - value: kk
      children: []
  - value: t
    children: []
- value: e
  children:
  - value: o
    children: []
- value: g
  children: []
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 4spaces & output yaml)",
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
				optFns: []OptFn{WithIndentFourSpaces(), WithEncodeYAML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value: a
children:
- value: i
  children:
  - value: u
    children:
    - value: k
      children: []
    - value: kk
      children: []
  - value: t
    children: []
- value: e
  children:
  - value: o
    children: []
- value: g
  children: []
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
