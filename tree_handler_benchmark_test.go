package gtree_test

import (
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

func BenchmarkOutput_singleRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(singleRoot)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var singleRoot = strings.TrimPrefix(`
- a
	- b
		- c
			- d
		- e
			- f
			- g
				- h
	- i
		- j
			- k
	- l
		- m
`, "\n")

func BenchmarkOutput_fiveRoots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(fiveRoots)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var fiveRoots = strings.TrimPrefix(`
- a1
	- b
		- c
			- d
		- e
			- f
			- g
				- h
	- i
		- j
			- k
	- l
		- m
- a2
	- b
		- c
			- d
		- e
			- f
			- g
				- h
	- i
		- j
			- k
	- l
		- m
- a3
	- b
		- c
			- d
		- e
			- f
			- g
				- h
	- i
		- j
			- k
	- l
		- m
- a4
	- b
		- c
			- d
		- e
			- f
			- g
				- h
	- i
		- j
			- k
	- l
		- m
- a5
	- b
		- c
			- d
		- e
			- f
			- g
				- h
	- i
		- j
			- k
	- l
		- m
`, "\n")

func BenchmarkOutput_hundredRoots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(hundredRoots)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var hundredRoots = strings.Repeat(singleRoot, 100)
