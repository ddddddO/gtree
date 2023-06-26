package gtree_test

import (
	"testing"

	tu "github.com/ddddddO/gtree/testutil"
)

func BenchmarkOutput_singleRoot(b *testing.B) {
	tu.BaseBenchmark(tu.SingleRoot, b)
}

func BenchmarkOutput_tenRoots(b *testing.B) {
	tu.BaseBenchmark(tu.TenRoots, b)
}

func BenchmarkOutput_fiftyRoots(b *testing.B) {
	tu.BaseBenchmark(tu.FiftyRoots, b)
}

func BenchmarkOutput_hundredRoots(b *testing.B) {
	tu.BaseBenchmark(tu.HundredRoots, b)
}

func BenchmarkOutput_fiveHundredsRoots(b *testing.B) {
	tu.BaseBenchmark(tu.FiveHundredsRoots, b)
}

func BenchmarkOutput_thousandRoots(b *testing.B) {
	tu.BaseBenchmark(tu.ThousandRoots, b)
}

func BenchmarkOutput_3000Roots(b *testing.B) {
	tu.BaseBenchmark(tu.ThreeThousandRoots, b)
}

func BenchmarkOutput_6000Roots(b *testing.B) {
	tu.BaseBenchmark(tu.SixThousandRoots, b)
}

func BenchmarkOutput_10000Roots(b *testing.B) {
	tu.BaseBenchmark(tu.TenThousandRoots, b)
}

// NOTE: ベンチマーク取得でタイムアウトになったため。条件を合わせるためpipeline側もコメントアウト
// func BenchmarkOutput_20000Roots(b *testing.B) {
// 	tu.BaseBenchmark(tu.TwentyThousandRoots, b)
// }
