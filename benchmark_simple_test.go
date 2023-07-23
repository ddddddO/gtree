package gtree_test

import (
	"testing"

	tu "github.com/ddddddO/gtree/testutil"
)

func BenchmarkOutput_singleRoot(b *testing.B) {
	tu.BaseBenchmark(b, tu.SingleRoot)
}

func BenchmarkOutput_tenRoots(b *testing.B) {
	tu.BaseBenchmark(b, tu.TenRoots)
}

func BenchmarkOutput_fiftyRoots(b *testing.B) {
	tu.BaseBenchmark(b, tu.FiftyRoots)
}

func BenchmarkOutput_hundredRoots(b *testing.B) {
	tu.BaseBenchmark(b, tu.HundredRoots)
}

func BenchmarkOutput_fiveHundredsRoots(b *testing.B) {
	tu.BaseBenchmark(b, tu.FiveHundredsRoots)
}

func BenchmarkOutput_thousandRoots(b *testing.B) {
	tu.BaseBenchmark(b, tu.ThousandRoots)
}

func BenchmarkOutput_3000Roots(b *testing.B) {
	tu.BaseBenchmark(b, tu.ThreeThousandRoots)
}

func BenchmarkOutput_6000Roots(b *testing.B) {
	tu.BaseBenchmark(b, tu.SixThousandRoots)
}

func BenchmarkOutput_10000Roots(b *testing.B) {
	tu.BaseBenchmark(b, tu.TenThousandRoots)
}

// NOTE: ベンチマーク取得でタイムアウトになったため。条件を合わせるためpipeline側もコメントアウト
// func BenchmarkOutput_20000Roots(b *testing.B) {
// 	tu.BaseBenchmark(b, tu.TwentyThousandRoots)
// }
