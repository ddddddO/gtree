package gtree_test

import (
	"testing"

	tu "github.com/ddddddO/gtree/testutil"
)

func BenchmarkOutput_simple_singleRoot(b *testing.B) {
	tu.BenchmarkOfSimpleOutput(b, tu.SingleRoot)
}

func BenchmarkOutput_simple_tenRoots(b *testing.B) {
	tu.BenchmarkOfSimpleOutput(b, tu.TenRoots)
}

func BenchmarkOutput_simple_fiftyRoots(b *testing.B) {
	tu.BenchmarkOfSimpleOutput(b, tu.FiftyRoots)
}

func BenchmarkOutput_simple_hundredRoots(b *testing.B) {
	tu.BenchmarkOfSimpleOutput(b, tu.HundredRoots)
}

func BenchmarkOutput_simple_fiveHundredsRoots(b *testing.B) {
	tu.BenchmarkOfSimpleOutput(b, tu.FiveHundredsRoots)
}

func BenchmarkOutput_simple_thousandRoots(b *testing.B) {
	tu.BenchmarkOfSimpleOutput(b, tu.ThousandRoots)
}

func BenchmarkOutput_simple_3000Roots(b *testing.B) {
	tu.BenchmarkOfSimpleOutput(b, tu.ThreeThousandRoots)
}

func BenchmarkOutput_simple_6000Roots(b *testing.B) {
	tu.BenchmarkOfSimpleOutput(b, tu.SixThousandRoots)
}

func BenchmarkOutput_simple_10000Roots(b *testing.B) {
	tu.BenchmarkOfSimpleOutput(b, tu.TenThousandRoots)
}

func BenchmarkOutput_simple_20000Roots(b *testing.B) {
	tu.BenchmarkOfSimpleOutput(b, tu.TwentyThousandRoots)
}
