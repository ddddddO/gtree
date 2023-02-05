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

func BenchmarkOutput_tenRoots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(tenRoots)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var tenRoots = strings.Repeat(singleRoot, 10)

func BenchmarkOutput_fiftyRoots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(fiftyRoots)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var fiftyRoots = strings.Repeat(singleRoot, 50)

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

func BenchmarkOutput_fiveHundredsRoots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(fiveHundredsRoots)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var fiveHundredsRoots = strings.Repeat(singleRoot, 500)

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

func BenchmarkOutput_6000Roots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(sixThousandRoots)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var sixThousandRoots = strings.Repeat(singleRoot, 6000)

func BenchmarkOutput_10000Roots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(tenThousandRoots)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var tenThousandRoots = strings.Repeat(singleRoot, 10000)

func BenchmarkOutput_20000Roots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		r := strings.NewReader(twentyThousandRoots)
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
	}
}

var twentyThousandRoots = strings.Repeat(singleRoot, 20000)
