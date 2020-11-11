package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	content := widget.NewVBox(input)
	content.Append(widget.NewButton("ADD", func() {
		content.Append(widget.NewLabel(input.Text))
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
