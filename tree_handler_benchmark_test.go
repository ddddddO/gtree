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

var fiveRoots = strings.Repeat(singleRoot, 5)

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

func BenchmarkOutput_thousandRoots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(thousandRoots)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var thousandRoots = strings.Repeat(singleRoot, 1000)

func BenchmarkOutput_3000Roots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(threeThousandRoots)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var threeThousandRoots = strings.Repeat(singleRoot, 3000)
