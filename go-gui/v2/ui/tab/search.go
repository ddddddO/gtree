package tab

import (
	"errors"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/ddddddO/work/go-gui/v2/db"
)

func NewSearch(window fyne.Window, sqlite *db.Sqlite, activityCh chan<- string) *widget.TabItem {
	searchWord := widget.NewEntry()
	searchContent := widget.NewVBox(
		widget.NewLabel("XXXX"),
		&widget.Form{
			Items: []*widget.FormItem{
				{Text: "Search word", Widget: searchWord},
			},
			OnSubmit: func() {
				activityCh <- "Enter serch"
				if searchWord.Text == "" {
					dialog.ShowError(errors.New("Please input search word"), window)
					return
				}

				result, err := sqlite.Select(searchWord.Text)
				if err != nil {
					dialog.ShowError(err, window)
					return
				}

				dialog.ShowInformation("search", result, window)
			},
		},
	)
	return &widget.TabItem{Text: "search", Icon: theme.SearchIcon(), Content: searchContent}
}
