package main

import (
	"strings"
	"syscall/js"

	gt "github.com/ddddddO/gtree"
)

func gtree(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")

	rawInput := document.Call("getElementById", "in").Get("value").String()
	console := js.Global().Get("console")
	console.Call("log", rawInput)

	w := &strings.Builder{}
	r := strings.NewReader(rawInput)
	options := []gt.Option{gt.WithIndentTwoSpaces()}

	err := gt.Output(w, r, options...)
	// TODO: error handling. 画面側にエラー内容を出力したい。
	if err != nil {
		panic(err)
	}

	pre := document.Call("createElement", "pre")
	pre.Set("id", "treeView")
	pre.Set("innerHTML", w.String())
	document.Get("body").Call("appendChild", pre)
	return nil
}

func registerCallbacks() {
	js.Global().Set("gtree", js.FuncOf(gtree))
}

func main() {
	c := make(chan struct{}, 0)
	println("gtree WebAssembly Initialized")
	registerCallbacks()
	<-c
}
