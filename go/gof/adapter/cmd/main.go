package main

import (
	"fmt"

	"github.com/ddddddO/work/go/gof/adapter"
)

// Client役(Target役を使う側)
func main() {
	desc := "ADAPTER"

	// Targetインタフェースで宣言しておけば、
	// Targetインタフェースを満たすadapter.go以外の例えばadapter_a.goとかで別の実装が出来るから、
	// インタフェース越しにメソッドを呼び出すようにしているのかと
	var a adapter.Target
	a = adapter.NewAdapter(desc)

	a.OutputHyphenFrame()
	fmt.Println()
	a.OutputSharpFrame()

	// Output:
	// -------------
	// ADAPTER
	// -------------
	//
	// #############
	// ADAPTER
	// #############
}
