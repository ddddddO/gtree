package gtree_test

import (
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

func BenchmarkOutput_pipeline_singleRoot(b *testing.B) {
	baseBenchmarkWithMassive(singleRootP, b)
}

func BenchmarkOutput_pipeline_tenRoots(b *testing.B) {
	baseBenchmarkWithMassive(tenRootsP, b)
}

func BenchmarkOutput_pipeline_fiftyRoots(b *testing.B) {
	baseBenchmarkWithMassive(fiftyRootsP, b)
}

func BenchmarkOutput_pipeline_hundredRoots(b *testing.B) {
	baseBenchmarkWithMassive(hundredRootsP, b)
}

func BenchmarkOutput_pipeline_fiveHundredsRoots(b *testing.B) {
	baseBenchmarkWithMassive(fiveHundredsRootsP, b)
}

func BenchmarkOutput_pipeline_thousandRoots(b *testing.B) {
	baseBenchmarkWithMassive(thousandRootsP, b)
}

func BenchmarkOutput_pipeline_3000Roots(b *testing.B) {
	baseBenchmarkWithMassive(threeThousandRootsP, b)
}

func BenchmarkOutput_pipeline_6000Roots(b *testing.B) {
	baseBenchmarkWithMassive(sixThousandRootsP, b)
}

func BenchmarkOutput_pipeline_10000Roots(b *testing.B) {
	baseBenchmarkWithMassive(tenThousandRootsP, b)
}

// func BenchmarkOutput_pipeline_20000Roots(b *testing.B) {
// 	baseBenchmarkWithMassive(twentyThousandRootsP, b)
// }

func baseBenchmarkWithMassive(roots string, b *testing.B) {
	r := strings.NewReader(roots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r, gtree.WithMassive()); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(roots)
	}
}

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
