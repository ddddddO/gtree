package gtree

import (
	"bytes"
	"io"
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
				output: strings.TrimSpace(`
a
└── b`),
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
				output: strings.TrimSpace(`
a
└── b
    └── c`),
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
				output: strings.TrimSpace(`
a
├── b
└── c`),
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
				output: strings.TrimSpace(`
a
└── b
    └── c
        ├── d
        ├── e
        └── f`),
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
				output: strings.TrimSpace(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g`),
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
				output: strings.TrimSpace(`
a
├── vvv
│   └── jjj
├── ggg
│   └── hhhh
└── ggggg`),
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
				output: strings.TrimSpace(`
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
└── eee`),
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
				output: strings.TrimSpace(`
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
└── eee`),
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
				output: strings.TrimSpace(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g`),
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
				output: strings.TrimSpace(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g`),
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
				output: strings.TrimSpace(`
root dir aaa
└── child-dir`),
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
				output: strings.TrimSpace(`
parent
├── child
│   ├── chilchil
│   ├── chilchil
│   └── chilchil
└── child`),
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
				output: strings.TrimSpace(`
a
└── b`),
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
				output: strings.TrimSpace(`
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
└── g`),
				err: nil,
			},
		},
		{
			name: "case 15(empty name)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	-`)),
			},
			out: out{
				output: "",
				err:    ErrEmptyName,
			},
		},
		{
			// TODO: inputのパターンが3つ(tab/ts/fs)と実行時のモードが3つで、それぞれの正常系(3つ)を上でしてるから、このパターン含めると、
			//       3*3-3=6パターンのケースが必要
			name: "case 16(incorrect input format(input 4spaces indent / execute tab mode))",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
    - b`)),
			},
			out: out{
				output: "",
				err:    ErrIncorrectFormat,
			},
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)

		gotOutput, gotErr := Execute(tt.in.input, tt.in.conf)

		if gotOutput != tt.out.output {
			t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.out.output)
		}
		if gotErr != tt.out.err {
			t.Errorf("\ngot: \n%v\nwant: \n%v", gotErr, tt.out.err)
		}
	}
}
