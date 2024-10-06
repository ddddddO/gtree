package main

import (
	"html/template"
	"strings"
	"syscall/js" // nolint

	tree "github.com/ddddddO/gtree"
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
	writer := &strings.Builder{}

	if err := tree.Output(
		writer,
		strings.NewReader(getElementByID("in").Get("value").String()),
		tree.WithBranchFormatLastNode(parts1+parts3, "    "),
		tree.WithBranchFormatIntermedialNode(parts2+parts3, parts4+"   "),
	); err != nil {
		// リアルタイムで tree 生成するようにし、ダイアログでエラー出したくないので、コメントアウトしておく
		// alert(err.Error())
		return nil
	}

	div := getElementByID("result")
	if prePre := getElementByID("treeView"); !prePre.IsNull() {
		removeChildFunc(div)(prePre)
	}

	pre := createElementFunc(document)("pre")
	pre.Set("id", "treeView")
	pre.Set("innerHTML", template.HTMLEscapeString(writer.String()))
	appendChildFunc(div)(pre)
	appendChildFunc(getElementByID("main"))(div)

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
