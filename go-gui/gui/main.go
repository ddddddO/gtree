package main

import (
	"log"
	"strconv"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"

	"github.com/pkg/errors"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("GUI APP")

	listURLsTab := newListURLsTab(myWindow)
	settingsTab := newSettingsTab(myWindow, listURLsTab)

	settingsTabItem := widget.NewTabItem("SETTINGS", settingsTab)
	listURLsTabItem := widget.NewTabItem("URL LIST", listURLsTab)

	myWindow.SetContent(
		widget.NewTabContainer(
			settingsTabItem,
			listURLsTabItem,
		),
	)

	myWindow.ShowAndRun()
}

func newSettingsTab(parent fyne.Window, listURLsTab fyne.CanvasObject) fyne.CanvasObject {
	serviceHostEntry := widget.NewEntry()
	serviceHostEntry.SetPlaceHolder("Enter Service Host name or IP")

	secretEntry := widget.NewPasswordEntry()
	secretEntry.SetPlaceHolder("Enter Your Secret ....")

	saveButton := widget.NewButton("Save", func() {
		host := serviceHostEntry.Text
		secret := secretEntry.Text
		if !isValidSettings(host, secret, parent) {
			return
		}
		log.Println(host, secret)

		listURLsBox := listURLsTab.(*widget.Box)
		listURLsBox.Children = nil
		listURLsBox.Append(widget.NewLabel("Target Host: http://" + host))
		listURLsBox.Append(widget.NewLabel("Your Secret: " + secret))
		listURLsBox.Append(widget.NewButton("Search", func() {
			listURLsBox.Children = listURLsBox.Children[:3]
			log.Println("GET http://" + host + ":8888")

			listURLsBox.Append(widget.NewLabel("Please Check Extra URLs"))
			extraDeleteTargetURLs := map[int]string{}

			totalCount := 5
			for i := 0; i < totalCount; i++ {
				ii := i // NOTE: 注意！

				check := widget.NewCheck(strconv.Itoa(ii), func(checked bool) {
					if checked {
						if _, ok := extraDeleteTargetURLs[ii]; !ok {
							extraDeleteTargetURLs[ii] = "/" + strconv.Itoa(ii)
						}
						log.Println("checked: " + strconv.Itoa(ii))
					} else {
						if _, ok := extraDeleteTargetURLs[ii]; ok {
							delete(extraDeleteTargetURLs, ii)
						}
						log.Println("unchecked: " + strconv.Itoa(ii))
					}
				})
				listURLsBox.Append(check)
			}

			listURLsBox.Append(widget.NewButton("Print Delete targets", func() {
				log.Println("Print Extra Delete targets")
				log.Println(extraDeleteTargetURLs)
			}))

			listURLsBox.Append(widget.NewButton("Exec Delete", func() {
				progressBar := widget.NewProgressBar()
				progressBar.Min = 0
				progressBar.Max = float64(totalCount - len(extraDeleteTargetURLs))
				listURLsBox.Append(progressBar)

				c := 0
				for range time.Tick(0.3 * 1000 * time.Millisecond) {
					if float64(c) > progressBar.Max {
						break
					}
					c++
					progressBar.SetValue(float64(c))
				}

				listURLsBox.Append(widget.NewLabel("Completed!"))
			}))

		}))
	})

	vbox := widget.NewVBox(
		widget.NewLabel("Service Host:"),
		serviceHostEntry,
		widget.NewLabel("Your Secret:"),
		secretEntry,
		saveButton,
	)

	return vbox
}

func isValidSettings(host, secret string, parent fyne.Window) bool {
	if host == "" && secret == "" {
		dialog.ShowError(errors.New("Please input Host & Secret."), parent)
		return false
	}
	if host == "" {
		dialog.ShowError(errors.New("Please input Host."), parent)
		return false
	}
	if secret == "" {
		dialog.ShowError(errors.New("Please input Secret."), parent)
		return false
	}

	return true
}

func newListURLsTab(parent fyne.Window) fyne.CanvasObject {
	vbox := widget.NewVBox()

	return vbox
}
