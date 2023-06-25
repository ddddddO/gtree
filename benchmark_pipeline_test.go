package gtree_test

import (
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

func BenchmarkOutput_pipeline_singleRoot(b *testing.B) {
	r := strings.NewReader(singleRootP)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(singleRootP)
	}
}

func BenchmarkOutput_pipeline_tenRoots(b *testing.B) {
	r := strings.NewReader(tenRootsP)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(tenRootsP)
	}
}

func BenchmarkOutput_pipeline_fiftyRoots(b *testing.B) {
	r := strings.NewReader(fiftyRootsP)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(fiftyRootsP)
	}
}

func BenchmarkOutput_pipeline_hundredRoots(b *testing.B) {
	r := strings.NewReader(hundredRootsP)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(hundredRootsP)
	}
}

func BenchmarkOutput_pipeline_fiveHundredsRoots(b *testing.B) {
	r := strings.NewReader(fiveHundredsRootsP)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(fiveHundredsRootsP)
	}
}

func BenchmarkOutput_pipeline_thousandRoots(b *testing.B) {
	r := strings.NewReader(thousandRootsP)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(thousandRootsP)
	}
}

func BenchmarkOutput_pipeline_3000Roots(b *testing.B) {
	r := strings.NewReader(threeThousandRootsP)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(threeThousandRootsP)
	}
}

func BenchmarkOutput_pipeline_6000Roots(b *testing.B) {
	r := strings.NewReader(sixThousandRootsP)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(sixThousandRootsP)
	}
}

func BenchmarkOutput_pipeline_10000Roots(b *testing.B) {
	r := strings.NewReader(tenThousandRootsP)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(tenThousandRootsP)
	}
}

// func BenchmarkOutput_pipeline_20000Roots(b *testing.B) {
// 	r := strings.NewReader(twentyThousandRootsP)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		w := &strings.Builder{}
// 		b.StartTimer()
// 		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
// 			b.Fatal(err)
// 		}
// 		b.StopTimer()
// 		r.Reset(twentyThousandRootsP)
// 	}
// }

var (
	singleRootP          = singleP
	tenRootsP            = strings.Repeat(singleRootP, 10)
	fiftyRootsP          = strings.Repeat(singleRootP, 50)
	hundredRootsP        = strings.Repeat(singleRootP, 100)
	fiveHundredsRootsP   = strings.Repeat(singleRootP, 500)
	thousandRootsP       = strings.Repeat(singleRootP, 1000)
	threeThousandRootsP  = strings.Repeat(singleRootP, 3000)
	sixThousandRootsP    = strings.Repeat(singleRootP, 6000)
	tenThousandRootsP    = strings.Repeat(singleRootP, 10000)
	twentyThousandRootsP = strings.Repeat(singleRootP, 20000)
)

var singleP = strings.TrimPrefix(`
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
