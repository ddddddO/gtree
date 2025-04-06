package gtree_test

import (
	"testing"

	tu "github.com/ddddddO/gtree/testutil"
)

func BenchmarkOutput_pipeline_singleRoot(b *testing.B) {
	tu.BenchmarkWithMassive(b, tu.SingleRoot)
}

func BenchmarkOutput_pipeline_tenRoots(b *testing.B) {
	tu.BenchmarkWithMassive(b, tu.TenRoots)
}

func BenchmarkOutput_pipeline_fiftyRoots(b *testing.B) {
	tu.BenchmarkWithMassive(b, tu.FiftyRoots)
}

func BenchmarkOutput_pipeline_hundredRoots(b *testing.B) {
	tu.BenchmarkWithMassive(b, tu.HundredRoots)
}

func BenchmarkOutput_pipeline_fiveHundredsRoots(b *testing.B) {
	tu.BenchmarkWithMassive(b, tu.FiveHundredsRoots)
}

func BenchmarkOutput_pipeline_thousandRoots(b *testing.B) {
	tu.BenchmarkWithMassive(b, tu.ThousandRoots)
}

func BenchmarkOutput_pipeline_3000Roots(b *testing.B) {
	tu.BenchmarkWithMassive(b, tu.ThreeThousandRoots)
}

func BenchmarkOutput_pipeline_6000Roots(b *testing.B) {
	tu.BenchmarkWithMassive(b, tu.SixThousandRoots)
}

func BenchmarkOutput_pipeline_10000Roots(b *testing.B) {
	tu.BenchmarkWithMassive(b, tu.TenThousandRoots)
}

func BenchmarkOutput_pipeline_20000Roots(b *testing.B) {
	tu.BenchmarkWithMassive(b, tu.TwentyThousandRoots)
}
