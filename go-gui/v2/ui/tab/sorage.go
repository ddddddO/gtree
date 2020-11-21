package tab

import (
	"errors"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/ddddddO/work/go-gui/v2/db"
)

func NewStorage(window fyne.Window, sqlite *db.Sqlite, activityCh chan<- string) *widget.TabItem {
	insertWord := widget.NewEntry()
	storageContent := widget.NewVBox(
		widget.NewLabel("XXXX"),
		&widget.Form{
			Items: []*widget.FormItem{
				{Text: "Insert", Widget: insertWord},
			},
			OnSubmit: func() {
				activityCh <- "Enter insert"
				if insertWord.Text == "" {
					dialog.ShowError(errors.New("Please input insert word"), window)
					return
				}
				if err := sqlite.Insert(insertWord.Text); err != nil {
					dialog.ShowError(err, window)
					return
				}
				dialog.ShowInformation("success", "success!", window)
			},
		},
	)
	return &widget.TabItem{Text: "storage", Icon: theme.StorageIcon(), Content: storageContent}
}
