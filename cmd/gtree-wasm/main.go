package main

import (
	"html/template"
	"strings"
	"syscall/js" // nolint

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
	getElementByID := getElementByIDFunc(document)

	parts1 := getElementByID("parts1").Get("value").String()
	parts2 := getElementByID("parts2").Get("value").String()
	parts3 := getElementByID("parts3").Get("value").String()
	parts4 := getElementByID("parts4").Get("value").String()

	lastNodeBranchDirectly := parts1 + parts3
	lastNodeBranchIndirectly := "    "
	options := []gt.Option{gt.WithIndentTwoSpaces()}
	options = append(options, gt.WithBranchFormatLastNode(lastNodeBranchDirectly, lastNodeBranchIndirectly))

	intermedialNodeBranchDirectly := parts2 + parts3
	intermedialNodeBranchIndirectly := parts4 + "   "
	options = append(options, gt.WithBranchFormatIntermedialNode(intermedialNodeBranchDirectly, intermedialNodeBranchIndirectly))

	rawInput := getElementByID("in").Get("value").String()

	w := &strings.Builder{}
	r := strings.NewReader(rawInput)
	err := gt.Output(w, r, options...)
	if err != nil {
		js.Global().Call("alert", err.Error())
		return nil
	}

	div := getElementByID("result")
	prePre := getElementByID("treeView")
	if !prePre.IsNull() {
		div.Call("removeChild", prePre)
	}

	pre := document.Call("createElement", "pre")
	pre.Set("id", "treeView")
	pre.Set("innerHTML", template.HTMLEscapeString(w.String()))
	div.Call("appendChild", pre)

	mainContainer := getElementByID("main")
	mainContainer.Call("appendChild", div)

	return nil
}

func getElementByIDFunc(document js.Value) func(id string) js.Value {
	return func(id string) js.Value {
		return document.Call("getElementById", id)
	}
}
