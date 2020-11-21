package tab

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func NewSettings() *widget.TabItem {
	settingsContent := widget.NewVBox(widget.NewLabel("XXXX"), widget.NewEntry())
	return &widget.TabItem{Text: "settings", Icon: theme.SettingsIcon(), Content: settingsContent}
}
