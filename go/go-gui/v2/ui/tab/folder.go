package tab

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func NewFolder() *widget.TabItem {
	folderContent := widget.NewVBox(widget.NewLabel("XXXX"), widget.NewEntry())
	return &widget.TabItem{Text: "folder", Icon: theme.FolderIcon(), Content: folderContent}
}
