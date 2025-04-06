package gtree_test

import (
	"testing"

	tu "github.com/ddddddO/gtree/testutil"
)

func BenchmarkOutput_iterator_singleRoot(b *testing.B) {
	tu.BaseBenchmark(b, tu.SingleRoot)
}

func BenchmarkOutput_iterator_tenRoots(b *testing.B) {
	tu.BaseBenchmark(b, tu.TenRoots)
}

func BenchmarkOutput_iterator_fiftyRoots(b *testing.B) {
	tu.BaseBenchmark(b, tu.FiftyRoots)
}

func BenchmarkOutput_iterator_hundredRoots(b *testing.B) {
	tu.BaseBenchmark(b, tu.HundredRoots)
}

func BenchmarkOutput_iterator_fiveHundredsRoots(b *testing.B) {
	tu.BaseBenchmark(b, tu.FiveHundredsRoots)
}

func BenchmarkOutput_iterator_thousandRoots(b *testing.B) {
	tu.BaseBenchmark(b, tu.ThousandRoots)
}

func BenchmarkOutput_iterator_3000Roots(b *testing.B) {
	tu.BaseBenchmark(b, tu.ThreeThousandRoots)
}

func BenchmarkOutput_iterator_6000Roots(b *testing.B) {
	tu.BaseBenchmark(b, tu.SixThousandRoots)
}

func BenchmarkOutput_iterator_10000Roots(b *testing.B) {
	tu.BaseBenchmark(b, tu.TenThousandRoots)
}

func BenchmarkOutput_iterator_20000Roots(b *testing.B) {
	tu.BaseBenchmark(b, tu.TwentyThousandRoots)
}
