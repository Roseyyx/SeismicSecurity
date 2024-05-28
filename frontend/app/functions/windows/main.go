package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainLayout(myWindow fyne.Window, myApp fyne.App) {
	Option1 := widget.NewButton("Database", func() {
		// drop down menu
		dropDown := widget.NewSelect([]string{"New Database", "Open Database", "Save Database"}, func(s string) {
			switch s {
			case "New Database":
				file := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
					if err != nil {
						dialog.ShowError(err, myWindow)
						return
					}
					if writer == nil {
						return
					}
					defer writer.Close()
				}, myWindow)
				file.SetFileName("database.Seismic")
				file.Show()
			case "Open Database":
				// open new window
				newWindow := myApp.NewWindow("Open Database")
				newWindow.Resize(fyne.NewSize(400, 300))
				newWindow.Show()
			case "Save Database":
				// open new window
				newWindow := myApp.NewWindow("Save Database")
				newWindow.Resize(fyne.NewSize(400, 300))
				newWindow.Show()
			}
		})

		canvasObject := container.New(layout.NewVBoxLayout(), widget.NewLabel("Database"), dropDown)
		myWindow.SetContent(canvasObject)
	})

	Option2 := widget.NewButton("Entries", func() {
		canvasObject := container.New(layout.NewVBoxLayout(), widget.NewLabel("Entries"))
		myWindow.SetContent(canvasObject)
	})

	Option3 := widget.NewButton("Settings", func() {
		canvasObject := container.New(layout.NewVBoxLayout(), widget.NewLabel("Settings"))
		myWindow.SetContent(canvasObject)
	})

	// make h split
	myWindow.SetContent(container.New(layout.NewVBoxLayout(),
		container.New(layout.NewHBoxLayout(),
			Option1,
			Option2,
			Option3,
		),
	))
}
