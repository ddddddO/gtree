package ui

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/ddddddO/work/go-gui/v2/db"
	"github.com/ddddddO/work/go-gui/v2/ui/tab"
)

func Run(sqlite *db.Sqlite) {
	application := app.New()
	window := application.NewWindow("GUI APP")

	activityCh := make(chan string)

	tabs := []*widget.TabItem{
		tab.NewMenu(),
		tab.NewHome(),
		tab.NewSettings(),
		tab.NewFolder(),
		tab.NewSearch(window, sqlite, activityCh),
		tab.NewStorage(window, sqlite, activityCh),
		tab.NewHistory(activityCh),
	}
	tabContainer := widget.NewTabContainer()
	for _, tab := range tabs {
		tabContainer.Append(widget.NewTabItemWithIcon(tab.Text, tab.Icon, tab.Content))
	}

	// tabバーの位置
	tabContainer.SetTabLocation(widget.TabLocationLeading)

	window.SetContent(tabContainer)
	window.ShowAndRun()
}
