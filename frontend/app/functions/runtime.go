package functions

import (
	"fmt"
	"log"
	"main/backend/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func check(myWindow fyne.Window) {
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
				check(myWindow)
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
				helpers.Login(input.Text, helpers.FilePath)
			}
		})

		content := container.New(layout.NewVBoxLayout(), label, input, button)

		myWindow.SetContent(content)

	}
}

func Setup() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	myWindow.Resize(fyne.NewSize(800, 600))

	check(myWindow)

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
