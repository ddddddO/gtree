package tab

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func NewHome() *widget.TabItem {
	homeContent := widget.NewVBox(widget.NewLabel("Your name?"), widget.NewEntry())
	return &widget.TabItem{Text: "home", Icon: theme.HomeIcon(), Content: homeContent}
}
