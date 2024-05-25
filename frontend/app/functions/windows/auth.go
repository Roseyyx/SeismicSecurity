package windows

import (
	"log"
	"main/backend/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Check(myWindow fyne.Window) {
	if !helpers.CheckFile() {

		input := widget.NewEntry()
		input.SetPlaceHolder("Enter your password")

		label := widget.NewLabel("Create a new file")

		button := widget.NewButton("Submit", func() {
			if input.Text == "" {
				label.SetText("Password cannot be empty")
			} else {
				log.Println(input.Text)
				helpers.Register(input.Text)
				CreateEntry(myWindow)
			}
		})

		content := container.New(layout.NewVBoxLayout(), label, input, button)

		myWindow.SetContent(content)
	} else {
		input := widget.NewEntry()
		input.SetPlaceHolder("Enter your password")

		label := widget.NewLabel("Open an existing file")

		button := widget.NewButton("Submit", func() {
			if input.Text == "" {
				label.SetText("Password cannot be empty")
			} else {
				if helpers.Login(input.Text, "database.Seismic") {
					ShowEntries(myWindow)
				} else {
					label.SetText("Password is incorrect")
				}
			}
		})

		content := container.New(layout.NewVBoxLayout(), label, input, button)

		myWindow.SetContent(content)

	}
}
