package gtree_test

import (
	"testing"

	tu "github.com/ddddddO/gtree/testutil"
)

func BenchmarkOutput_pipeline_singleRoot(b *testing.B) {
	tu.BaseBenchmarkWithMassive(tu.SingleRoot, b)
}

func BenchmarkOutput_pipeline_tenRoots(b *testing.B) {
	tu.BaseBenchmarkWithMassive(tu.TenRoots, b)
}

func BenchmarkOutput_pipeline_fiftyRoots(b *testing.B) {
	tu.BaseBenchmarkWithMassive(tu.FiftyRoots, b)
}

func BenchmarkOutput_pipeline_hundredRoots(b *testing.B) {
	tu.BaseBenchmarkWithMassive(tu.HundredRoots, b)
}

func BenchmarkOutput_pipeline_fiveHundredsRoots(b *testing.B) {
	tu.BaseBenchmarkWithMassive(tu.FiveHundredsRoots, b)
}

func BenchmarkOutput_pipeline_thousandRoots(b *testing.B) {
	tu.BaseBenchmarkWithMassive(tu.ThousandRoots, b)
}

func BenchmarkOutput_pipeline_3000Roots(b *testing.B) {
	tu.BaseBenchmarkWithMassive(tu.ThreeThousandRoots, b)
}

func BenchmarkOutput_pipeline_6000Roots(b *testing.B) {
	tu.BaseBenchmarkWithMassive(tu.SixThousandRoots, b)
}

func BenchmarkOutput_pipeline_10000Roots(b *testing.B) {
	tu.BaseBenchmarkWithMassive(tu.TenThousandRoots, b)
}

// func BenchmarkOutput_pipeline_20000Roots(b *testing.B) {
// 	tu.BaseBenchmarkWithMassive(tu.TwentyThousandRoots, b)
// }
