package windows

import (
	"log"
	"main/backend/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateEntry(myWindow fyne.Window) {
	button := widget.NewButton("Create a new entry", func() {
		input := widget.NewEntry()
		input.SetPlaceHolder("Enter your username")

		input2 := widget.NewEntry()
		input2.SetPlaceHolder("Enter your password")

		input3 := widget.NewEntry()
		input3.SetPlaceHolder("Enter the website")

		input4 := widget.NewEntry()
		input4.SetPlaceHolder("Notes")

		button2 := widget.NewButton("Submit", func() {
			if input.Text == "" || input2.Text == "" || input3.Text == "" {
				log.Println("Fields cannot be empty")
			} else {
				helpers.CreateEntry(input.Text, input2.Text, input3.Text, input4.Text)
			}
		})

		content := container.New(layout.NewVBoxLayout(), input, input2, input3, input4, button2)
		myWindow.SetContent(content)

	})

	myWindow.SetContent(button)

}

func ShowEntries(myWindow fyne.Window) {
	entries := helpers.GetEntries(helpers.GetFromMemory().Password, helpers.GetFromMemory().Filename)

	list := widget.NewList(
		func() int {
			return len(entries)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(entries[i].Username)
		},
	)

	myWindow.SetContent(list)
}
