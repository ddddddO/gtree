package main

import (
	"testing"
)

func TestGen(t *testing.T) {
	tests := []struct {
		name, in, out string
	}{
		{
			name: "case 1",
			in: `
- a
	- b`,
			out: `
a
└── b`,
		},
		{
			name: "case 2",
			in: `
- a
	- b
		- c
			- d
			- e
			- f`,
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
			in: `
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`,
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
