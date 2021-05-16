package main

import (
	"io"
	"strings"
	"testing"
)

func TestGen(t *testing.T) {
	tests := []struct {
		name, out string
		in        io.Reader
	}{
		{
			name: "case 1",
			in: strings.NewReader(strings.TrimSpace(`
- a
	- b`)),
			out: strings.TrimSpace(`
a
└── b`),
		},
		{
			name: "case 2",
			in: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c`)),
			out: strings.TrimSpace(`
a
└── b
    └── c`),
		},
		{
			name: "case 3",
			in: strings.NewReader(strings.TrimSpace(`
- a
	- b
	- c`)),
			out: strings.TrimSpace(`
a
├── b
└── c`),
		},

		{
			name: "case 4",
			in: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c
			- d
			- e
			- f`)),
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
			in: strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`)),
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
	}

	for _, tt := range tests {
		t.Log(tt.name)

		got := gen(tt.in)

		if got != tt.out {
			t.Errorf("\ngot: \n%s\nwant: \n%s", got, tt.out)
		}
	}
}
