package testutil

import (
	"context"
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

func BenchmarkOfSimpleOutput(b *testing.B, roots string) {
	BaseBenchmark(b, roots, gtree.WithNoUseIterOfSimpleOutput())
}

func BenchmarkWithMassive(b *testing.B, roots string) {
	BaseBenchmark(b, roots, gtree.WithMassive(context.Background()))
}

func BaseBenchmark(b *testing.B, roots string, options ...gtree.Option) {
	r := strings.NewReader(roots)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := &strings.Builder{}
		b.StartTimer()
		if err := gtree.OutputFromMarkdown(w, r, options...); err != nil {
			b.Fatal(err)
		}
		b.StopTimer()
		r.Reset(roots)
	}
}

var (
	SingleRoot          = single
	TenRoots            = strings.Repeat(SingleRoot, 10)
	FiftyRoots          = strings.Repeat(SingleRoot, 50)
	HundredRoots        = strings.Repeat(SingleRoot, 100)
	FiveHundredsRoots   = strings.Repeat(SingleRoot, 500)
	ThousandRoots       = strings.Repeat(SingleRoot, 1000)
	ThreeThousandRoots  = strings.Repeat(SingleRoot, 3000)
	SixThousandRoots    = strings.Repeat(SingleRoot, 6000)
	TenThousandRoots    = strings.Repeat(SingleRoot, 10000)
	TwentyThousandRoots = strings.Repeat(SingleRoot, 20000)
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
