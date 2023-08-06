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

	intermedialNodeBranchDirectly := parts2 + parts3
	intermedialNodeBranchIndirectly := parts4 + "   "

	rawInput := getElementByID("in").Get("value").String()
	r := strings.NewReader(rawInput)
	w := &strings.Builder{}
	options := []gt.Option{
		gt.WithBranchFormatLastNode(lastNodeBranchDirectly, lastNodeBranchIndirectly),
		gt.WithBranchFormatIntermedialNode(intermedialNodeBranchDirectly, intermedialNodeBranchIndirectly),
	}
	if err := gt.Output(w, r, options...); err != nil {
		alert(err.Error())
		return nil
	}

	div := getElementByID("result")
	if prePre := getElementByID("treeView"); !prePre.IsNull() {
		removeChildFunc(div)(prePre)
	}

	pre := createElementFunc(document)("pre")
	pre.Set("id", "treeView")
	pre.Set("innerHTML", template.HTMLEscapeString(w.String()))
	appendChildFunc(div)(pre)

	mainContainer := getElementByID("main")
	appendChildFunc(mainContainer)(div)

	return nil
}

func getElementByIDFunc(document js.Value) func(id string) js.Value {
	return func(id string) js.Value {
		return document.Call("getElementById", id)
	}
}

func createElementFunc(document js.Value) func(element string) js.Value {
	return func(element string) js.Value {
		return document.Call("createElement", element)
	}
}

func removeChildFunc(element js.Value) func(target js.Value) {
	return func(target js.Value) {
		element.Call("removeChild", target)
	}
}

func appendChildFunc(element js.Value) func(target js.Value) {
	return func(target js.Value) {
		element.Call("appendChild", target)
	}
}

func alert(msg string) {
	js.Global().Call("alert", msg)
}
