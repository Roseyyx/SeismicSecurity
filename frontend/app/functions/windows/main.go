package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainLayout(myWindow fyne.Window) {
	Option1 := widget.NewButton("Database", func() {
		canvasObject := container.New(layout.NewVBoxLayout(), widget.NewLabel("Database"))
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

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), Option1, Option2, Option3))
}
