package gtree_test

import (
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

func BenchmarkOutput_singleRoot(b *testing.B) {
	baseBenchmark(singleRoot, b)
}

func BenchmarkOutput_tenRoots(b *testing.B) {
	baseBenchmark(tenRoots, b)
}

func BenchmarkOutput_fiftyRoots(b *testing.B) {
	baseBenchmark(fiftyRoots, b)
}

func BenchmarkOutput_hundredRoots(b *testing.B) {
	baseBenchmark(hundredRoots, b)
}

func BenchmarkOutput_fiveHundredsRoots(b *testing.B) {
	baseBenchmark(fiveHundredsRoots, b)
}

func BenchmarkOutput_thousandRoots(b *testing.B) {
	baseBenchmark(thousandRoots, b)
}

func BenchmarkOutput_3000Roots(b *testing.B) {
	baseBenchmark(threeThousandRoots, b)
}

func BenchmarkOutput_6000Roots(b *testing.B) {
	baseBenchmark(sixThousandRoots, b)
}

func BenchmarkOutput_10000Roots(b *testing.B) {
	baseBenchmark(tenThousandRoots, b)
}

// NOTE: ベンチマーク取得でタイムアウトになったため。条件を合わせるためpipeline側もコメントアウト
// func BenchmarkOutput_20000Roots(b *testing.B) {
// 	baseBenchmark(twentyThousandRoots, b)
// }

func baseBenchmark(roots string, b *testing.B) {
	r := strings.NewReader(roots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(roots)
	}
}

var (
	singleRoot          = single
	tenRoots            = strings.Repeat(singleRoot, 10)
	fiftyRoots          = strings.Repeat(singleRoot, 50)
	hundredRoots        = strings.Repeat(singleRoot, 100)
	fiveHundredsRoots   = strings.Repeat(singleRoot, 500)
	thousandRoots       = strings.Repeat(singleRoot, 1000)
	threeThousandRoots  = strings.Repeat(singleRoot, 3000)
	sixThousandRoots    = strings.Repeat(singleRoot, 6000)
	tenThousandRoots    = strings.Repeat(singleRoot, 10000)
	twentyThousandRoots = strings.Repeat(singleRoot, 20000)
)

var single = strings.TrimPrefix(`
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
