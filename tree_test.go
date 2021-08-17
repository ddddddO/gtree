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
	input io.Reader
	conf  Config
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
				conf: Config{
					IsTwoSpaces: true,
				},
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
				conf: Config{
					IsFourSpaces: true,
				},
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
	}

	for _, tt := range tests {
		t.Log(tt.name)

		buf := &bytes.Buffer{}
		gotErr := Execute(buf, tt.in.input, tt.in.conf)
		gotOutput := buf.String()

		if gotOutput != tt.out.output {
			t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.out.output)
		}
		if gotErr != tt.out.err {
			t.Errorf("\ngot: \n%v\nwant: \n%v", gotErr, tt.out.err)
		}

		if file, ok := tt.in.input.(*os.File); ok {
			file.Close()
		}
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
