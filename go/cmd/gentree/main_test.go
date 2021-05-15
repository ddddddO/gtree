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
			in: strings.NewReader(`
- a
	- b`),
			out: `
a
└── b`,
		},
		{
			name: "case 2",
			in: strings.NewReader(`
- a
	- b
		- c
			- d
			- e
			- f`),
			out: `
a
└── b
    └── c
        ├── d
        ├── e
        └── f`,
		},
		{
			name: "case 3",
			in: strings.NewReader(`
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`),
			out: `
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g`,
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)

		got := gen(tt.in)

		if got != tt.out {
			t.Errorf("got: \n%s\nwant: \n%s", got, tt.out)
		}
	}
}
