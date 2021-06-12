package main

import (
	"io"
	"strings"
	"testing"
)

type in struct {
	input                     io.Reader
	isTwoSpaces, isFourSpaces bool
}

func TestGen(t *testing.T) {
	tests := []struct {
		name, out string
		in        in
	}{
		{
			name: "case 1",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b`)),
			},
			out: strings.TrimSpace(`
a
└── b`),
		},
		{
			name: "case 2",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c`))},

			out: strings.TrimSpace(`
a
└── b
    └── c`),
		},
		{
			name: "case 3",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
	- c`))},
			out: strings.TrimSpace(`
a
├── b
└── c`),
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
			out: strings.TrimSpace(`
a
└── b
    └── c
        ├── d
        ├── e
        └── f`),
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
			out: strings.TrimSpace(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g`),
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
			out: strings.TrimSpace(`
a
├── vvv
│   └── jjj
├── ggg
│   └── hhhh
└── ggggg`),
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
			out: strings.TrimSpace(`
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
			out: strings.TrimSpace(`
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
				isTwoSpaces: true,
			},
			out: strings.TrimSpace(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g`),
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
				isFourSpaces: true,
			},
			out: strings.TrimSpace(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g`),
		},
		{
			name: "case 11(1space & -)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root dir aaa
	- child-dir`))},
			out: strings.TrimSpace(`
root dir aaa
└── child-dir`),
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
			out: strings.TrimSpace(`
parent
├── child
│   ├── chilchil
│   ├── chilchil
│   └── chilchil
└── child`),
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)

		got := gen(tt.in.input, tt.in.isTwoSpaces, tt.in.isFourSpaces)

		if got != tt.out {
			t.Errorf("\ngot: \n%s\nwant: \n%s", got, tt.out)
		}
	}
}
