package main

import (
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/pkg/errors"
)

var registedSearchWord string

func main() {
	log.Println("a")

	application := app.New()
	window := application.NewWindow("GUI APP")

	homeContent := widget.NewVBox(widget.NewLabel("Your name?"), widget.NewEntry())
	menuContent := widget.NewVBox(widget.NewLabel("XXXX"), widget.NewEntry())
	settingsContent := widget.NewVBox(widget.NewLabel("XXXX"), widget.NewEntry())
	folderContent := widget.NewVBox(widget.NewLabel("XXXX"), widget.NewEntry())
	searchWord := widget.NewEntry()
	searchContent := widget.NewVBox(
		widget.NewLabel("XXXX"),
		&widget.Form{
			Items: []*widget.FormItem{
				{Text: "Search word", Widget: searchWord},
			},
			OnSubmit: func() {
				if searchWord.Text == "" {
					dialog.ShowError(errors.New("Please input search word"), window)
					return
				}
				registedSearchWord = searchWord.Text
				log.Println(registedSearchWord)
			},
		},
	)
	storageContent := widget.NewVBox(
		widget.NewLabel("XXXX"),
		widget.NewButton("Result", func() {
			log.Println(registedSearchWord)
		}),
	)

	tabs := []*widget.TabItem{
		{Text: "menu", Icon: theme.MenuIcon(), Content: menuContent},
		{Text: "home", Icon: theme.HomeIcon(), Content: homeContent},
		{Text: "settings", Icon: theme.SettingsIcon(), Content: settingsContent},
		{Text: "folder", Icon: theme.FolderIcon(), Content: folderContent},
		{Text: "search", Icon: theme.SearchIcon(), Content: searchContent},
		{Text: "storage", Icon: theme.StorageIcon(), Content: storageContent},
	}
	tabContainer := widget.NewTabContainer()
	for _, tab := range tabs {
		tabContainer.Append(widget.NewTabItemWithIcon("", tab.Icon, tab.Content))
	}

	window.SetContent(tabContainer)
	window.ShowAndRun()
}
