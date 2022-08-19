package main

import (
	"strings"
	"syscall/js"

	gt "github.com/ddddddO/gtree"
)

func gtree(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")

	options := []gt.Option{gt.WithIndentTwoSpaces()}

	branch1Directly := document.Call("getElementById", "branch1Directly").Get("value").String()
	branch1Indirectly := document.Call("getElementById", "branch1Indirectly").Get("value").String()
	options = append(options, gt.WithBranchFormatLastNode(branch1Directly, branch1Indirectly))

	branch2Directly := document.Call("getElementById", "branch2Directly").Get("value").String()
	branch2Indirectly := document.Call("getElementById", "branch2Indirectly").Get("value").String()
	options = append(options, gt.WithBranchFormatIntermedialNode(branch2Directly, branch2Indirectly))

	rawInput := document.Call("getElementById", "in").Get("value").String()
	// console := js.Global().Get("console")
	// console.Call("log", rawInput)

	w := &strings.Builder{}
	r := strings.NewReader(rawInput)
	err := gt.Output(w, r, options...)
	if err != nil {
		js.Global().Call("alert", err.Error())
		return nil
	}

	prePre := document.Call("getElementById", "treeView")
	if !prePre.IsNull() {
		document.Get("body").Call("removeChild", prePre)
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
