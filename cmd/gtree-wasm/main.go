package main

import (
	"strings"
	"syscall/js" // nolint

	gt "github.com/ddddddO/gtree"
)

func gtree(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	options := []gt.Option{gt.WithIndentTwoSpaces()}

	parts1 := document.Call("getElementById", "parts1").Get("value").String()
	parts2 := document.Call("getElementById", "parts2").Get("value").String()
	parts3 := document.Call("getElementById", "parts3").Get("value").String()
	parts4 := document.Call("getElementById", "parts4").Get("value").String()

	lastNodeBranchDirectly := parts1 + parts3
	lastNodeBranchIndirectly := "    "
	options = append(options, gt.WithBranchFormatLastNode(lastNodeBranchDirectly, lastNodeBranchIndirectly))

	intermedialNodeBranchDirectly := parts2 + parts3
	intermedialNodeBranchIndirectly := parts4 + "   "
	options = append(options, gt.WithBranchFormatIntermedialNode(intermedialNodeBranchDirectly, intermedialNodeBranchIndirectly))

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
