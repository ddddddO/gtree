package gtree_test

import (
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

func BenchmarkOutput_singleRoot(b *testing.B) {
	r := strings.NewReader(singleRoot)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(singleRoot)
	}
}

func BenchmarkOutput_tenRoots(b *testing.B) {
	r := strings.NewReader(tenRoots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(tenRoots)
	}
}

func BenchmarkOutput_fiftyRoots(b *testing.B) {
	r := strings.NewReader(fiftyRoots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(fiftyRoots)
	}
}

func BenchmarkOutput_hundredRoots(b *testing.B) {
	r := strings.NewReader(hundredRoots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(hundredRoots)
	}
}

func BenchmarkOutput_fiveHundredsRoots(b *testing.B) {
	r := strings.NewReader(fiveHundredsRoots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(fiveHundredsRoots)
	}
}

func BenchmarkOutput_thousandRoots(b *testing.B) {
	r := strings.NewReader(thousandRoots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(thousandRoots)
	}
}

func BenchmarkOutput_3000Roots(b *testing.B) {
	r := strings.NewReader(threeThousandRoots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(threeThousandRoots)
	}
}

func BenchmarkOutput_6000Roots(b *testing.B) {
	r := strings.NewReader(sixThousandRoots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(sixThousandRoots)
	}
}

func BenchmarkOutput_10000Roots(b *testing.B) {
	r := strings.NewReader(tenThousandRoots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.Output(w, r); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(tenThousandRoots)
	}
}

// NOTE: ベンチマーク取得でタイムアウトになったため。条件を合わせるためpipeline側もコメントアウト
// func BenchmarkOutput_20000Roots(b *testing.B) {
// 	r := strings.NewReader(twentyThousandRoots)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		w := &strings.Builder{}
// 		b.StartTimer()
// 		if err := gtree.Output(w, r); err != nil {
// 			b.Fatal(err)
// 		}
// 		b.StopTimer()
// 		r.Reset(twentyThousandRoots)
// 	}
// }

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
