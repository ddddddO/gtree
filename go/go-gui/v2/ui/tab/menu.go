package tab

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func NewMenu() *widget.TabItem {
	menuContent := widget.NewVBox(widget.NewLabel("XXXX"), widget.NewEntry())
	return &widget.TabItem{Text: "menu", Icon: theme.MenuIcon(), Content: menuContent}
}
