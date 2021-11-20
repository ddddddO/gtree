package gtree

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

type in struct {
	input  io.Reader
	optFns []OptFn
}

type out struct {
	output string
	err    error
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case 1",
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
			name: "case 2",
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
			name: "case 3",
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
			name: "case 4",
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
			name: "case 5",
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
			name: "case 6",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- vvv
		- jjj
	- ggg
		- hhhh
	- ggggg`))},
			out: out{
				output: strings.TrimPrefix(`
a
├── vvv
│   └── jjj
├── ggg
│   └── hhhh
└── ggggg
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case 7",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root
	- child1
	- child2
		- chilchil
	- dddd
		- kkkkkkk
			- lllll
				- ffff
				- ppppp
		- oooo
	- eee`))},
			out: out{
				output: strings.TrimPrefix(`
root
├── child1
├── child2
│   └── chilchil
├── dddd
│   ├── kkkkkkk
│   │   └── lllll
│   │       ├── ffff
│   │       └── ppppp
│   └── oooo
└── eee
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case 8",
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
			name: "case 9(indent 2spaces)",
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
				optFns: []OptFn{IndentTwoSpaces()},
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
			name: "case 10(indent 4spaces)",
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
				optFns: []OptFn{IndentFourSpaces()},
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
			name: "case 11(1space & -)",
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
			name: "case 12(same name)",
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
			name: "case 13(byte)",
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
			name: "case 14(multi root)",
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
			name: "case 15(empty text)",
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
			//       そのため、case 17 ~ case 21を予約
			name: "case 16(incorrect input format(input 4spaces indent / execute tab mode))",
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
			name: "case 22(bufio.Scanner err)",
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
			name: "case 23(input markdown file)",
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
			name: "case 24(invalid options)",
			in: in{
				input:  prepareMarkdownFile(t),
				optFns: []OptFn{IndentTwoSpaces(), IndentFourSpaces()},
			},
			out: out{
				output: "",
				err:    errInvalidOption,
			},
		},
		{
			name: "case 25(indent 2spaces and cutom branch format)",
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
					IndentTwoSpaces(),
					BranchFormatIntermedialNode("+->", ":   "),
					BranchFormatLastNode("+->", "    "),
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
		// json
		{
			name: "case 26(tab spaces & multi root & output json)",
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
				optFns: []OptFn{EncodeJSON()},
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
			name: "case 27(indent 2spaces & output json)",
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
				optFns: []OptFn{IndentTwoSpaces(), EncodeJSON()},
			},
			out: out{
				output: strings.TrimPrefix(`
{"value":"a","children":[{"value":"i","children":[{"value":"u","children":[{"value":"k","children":null},{"value":"kk","children":null}]},{"value":"t","children":null}]},{"value":"e","children":[{"value":"o","children":null}]},{"value":"g","children":null}]}
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case 28(indent 4spaces & output json)",
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
				optFns: []OptFn{IndentFourSpaces(), EncodeJSON()},
			},
			out: out{
				output: strings.TrimPrefix(`
{"value":"a","children":[{"value":"i","children":[{"value":"u","children":[{"value":"k","children":null},{"value":"kk","children":null}]},{"value":"t","children":null}]},{"value":"e","children":[{"value":"o","children":null}]},{"value":"g","children":null}]}
`, "\n"),
				err: nil,
			},
		},
		// yaml
		{
			name: "case 29(tab spaces & multi root & output yaml)",
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
				optFns: []OptFn{EncodeYAML()},
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
			name: "case 30(indent 2spaces & output json)",
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
				optFns: []OptFn{IndentTwoSpaces(), EncodeYAML()},
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
			name: "case 31(indent 4spaces & output json)",
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
				optFns: []OptFn{IndentFourSpaces(), EncodeYAML()},
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
		// toml
		{
			name: "case 32(tab spaces & multi root & output toml)",
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
				optFns: []OptFn{EncodeTOML()},
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
			name: "case 33(indent 2spaces & output toml)",
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
				optFns: []OptFn{IndentTwoSpaces(), EncodeTOML()},
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
			name: "case 34(indent 4spaces & output toml)",
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
				optFns: []OptFn{IndentFourSpaces(), EncodeTOML()},
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
			gotErr := Execute(out, tt.in.input, tt.in.optFns...)
			gotOutput := out.String()

			if gotOutput != tt.out.output {
				t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.out.output)
			}
			if gotErr != tt.out.err {
				t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.out.err)
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
