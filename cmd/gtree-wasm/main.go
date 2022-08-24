package main

import (
	"strings"
	"syscall/js" // nolint
	"text/template"

	gt "github.com/ddddddO/gtree"
)

func main() {
	c := make(chan struct{}, 0)
	println("gtree WebAssembly Initialized")
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("gtree", js.FuncOf(gtree))
}

func gtree(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	getElementByIdFunc := getElementById(document)

	parts1 := getElementByIdFunc("parts1").Get("value").String()
	parts2 := getElementByIdFunc("parts2").Get("value").String()
	parts3 := getElementByIdFunc("parts3").Get("value").String()
	parts4 := getElementByIdFunc("parts4").Get("value").String()

	lastNodeBranchDirectly := parts1 + parts3
	lastNodeBranchIndirectly := "    "
	options := []gt.Option{gt.WithIndentTwoSpaces()}
	options = append(options, gt.WithBranchFormatLastNode(lastNodeBranchDirectly, lastNodeBranchIndirectly))

	intermedialNodeBranchDirectly := parts2 + parts3
	intermedialNodeBranchIndirectly := parts4 + "   "
	options = append(options, gt.WithBranchFormatIntermedialNode(intermedialNodeBranchDirectly, intermedialNodeBranchIndirectly))

	rawInput := getElementByIdFunc("in").Get("value").String()

	w := &strings.Builder{}
	r := strings.NewReader(rawInput)
	err := gt.Output(w, r, options...)
	if err != nil {
		js.Global().Call("alert", err.Error())
		return nil
	}

	prePre := getElementByIdFunc("treeView")
	if !prePre.IsNull() {
		document.Get("body").Call("removeChild", prePre)
	}

	pre := document.Call("createElement", "pre")
	pre.Set("id", "treeView")
	pre.Set("innerHTML", template.HTMLEscapeString(w.String()))
	document.Get("body").Call("appendChild", pre)

	return nil
}

func getElementById(document js.Value) func(id string) js.Value {
	return func(id string) js.Value {
		return document.Call("getElementById", id)
	}
}
