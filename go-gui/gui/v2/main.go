package main

import (
	"log"
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/ddddddO/work/go-gui/v2/db"
	"github.com/pkg/errors"
)

func main() {
	sqlite, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite.CloseSQLite()

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

				result, err := sqlite.Select(searchWord.Text)
				if err != nil {
					dialog.ShowError(err, window)
					return
				}

				dialog.ShowInformation("search", result, window)
			},
		},
	)
	insertWord := widget.NewEntry()
	storageContent := widget.NewVBox(
		widget.NewLabel("XXXX"),
		&widget.Form{
			Items: []*widget.FormItem{
				{Text: "Insert", Widget: insertWord},
			},
			OnSubmit: func() {
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
	historyContent := widget.NewVBox(widget.NewLabel("Activity"), widget.NewSeparator())

	// NOTE: https://developer.fyne.io/tour/basics/canvas.html
	// 上記リンクな感じで表示を追加できる
	go func() {
		for {
			time.Sleep(2 * time.Second)
			historyContent.Append(widget.NewLabel("aa"))
		}
	}()

	tabs := []*widget.TabItem{
		{Text: "menu", Icon: theme.MenuIcon(), Content: menuContent},
		{Text: "home", Icon: theme.HomeIcon(), Content: homeContent},
		{Text: "settings", Icon: theme.SettingsIcon(), Content: settingsContent},
		{Text: "folder", Icon: theme.FolderIcon(), Content: folderContent},
		{Text: "search", Icon: theme.SearchIcon(), Content: searchContent},
		{Text: "storage", Icon: theme.StorageIcon(), Content: storageContent},
		{Text: "history", Icon: theme.HistoryIcon(), Content: historyContent},
	}
	tabContainer := widget.NewTabContainer()
	for _, tab := range tabs {
		tabContainer.Append(widget.NewTabItemWithIcon(tab.Text, tab.Icon, tab.Content))
	}

	// tabバーの位置
	tabContainer.SetTabLocation(widget.TabLocationLeading)

	window.SetContent(tabContainer)

	log.Println(window)

	window.ShowAndRun()
}
