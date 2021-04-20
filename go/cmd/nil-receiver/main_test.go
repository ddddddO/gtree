// ref: https://shogo82148.github.io/blog/2017/04/13/go1-8-allocation/

package main

import (
	"runtime"
	"testing"
)

func BenchmarkNewHumanEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewHumanEmpty()
	}
}

// Result:
// BenchmarkNewHumanEmpty-4                1000000000               1.11 ns/op           0 B/op                    0 allocs/op

func BenchmarkMakeNewHumanNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewHumanNil()
	}
}

// Result:
// BenchmarkMakeNewHumanNil-4              1000000000               0.451 ns/op          0 B/op                    0 allocs/op

/*----------------------------------------------------------------*/

func BenchmarkNewHumanEmpty_KeepAlive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(NewHumanEmpty())
	}
}

// Result:
// BenchmarkNewHumanEmpty_KeepAlive-4      161403067                9.31 ns/op           0 B/op                    0 allocs/op

func BenchmarkMakeNewHumanNil_KeepAlive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(NewHumanNil())
	}
}

// Result:
// BenchmarkMakeNewHumanNil_KeepAlive-4    1000000000               0.838 ns/op          0 B/op                    0 allocs/op
