package tab

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func NewHistory(activityCh <-chan string) *widget.TabItem {
	historyContent := widget.NewVBox(widget.NewLabel("Activity"), widget.NewSeparator())

	// NOTE: https://developer.fyne.io/tour/basics/canvas.html
	// 上記リンクな感じで表示を追加できる
	go func() {
		for {
			select {
			case activity := <-activityCh:
				historyContent.Append(widget.NewLabel(activity))
			}
		}
	}()

	return &widget.TabItem{Text: "history", Icon: theme.HistoryIcon(), Content: historyContent}
}
