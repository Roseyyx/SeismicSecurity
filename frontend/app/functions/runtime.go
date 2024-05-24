package functions

import (
	"encoding/json"
	"fmt"
	"main/backend/helpers"
	"main/backend/models"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetAllEntries() ([]models.Entry, error) {
	resp, err := http.Get("http://localhost:4545/api/v1/entries")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var entries []models.Entry

	err = json.NewDecoder(resp.Body).Decode(&entries)

	if err != nil {
		return nil, err
	}

	return entries, nil
}

func Setup() {
	// entries, err := GetAllEntries()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	myWindow.Resize(fyne.NewSize(800, 600))

	if !helpers.CheckFile() {

		input := widget.NewEntry()
		input.SetPlaceHolder("Enter your password")

		label := widget.NewLabel("Create a new file")

		button := widget.NewButton("Submit", func() {
			if input.Text == "" {
				label.SetText("Password cannot be empty")
			}
			helpers.Register(input.Text)
		})

		// how do i make it so not everything is centered but instead the label is on top and the input and button are centered
		content := container.New(layout.NewVBoxLayout(), label, input, button)

		myWindow.SetContent(content)

	} else {
		input := widget.NewEntry()
		input.SetPlaceHolder("Enter your password")

		label := widget.NewLabel("Open an existing file")

		button := widget.NewButton("Submit", func() {
			if input.Text == "" {
				label.SetText("Password cannot be empty")
			}
			helpers.Login(input.Text, helpers.FilePath)
		})

		content := container.New(layout.NewVBoxLayout(), label, input, button)

		myWindow.SetContent(content)

	}

	// listView := widget.NewList(func() int {
	// 	return 10 //len(entries)
	// }, func() fyne.CanvasObject {
	// 	return widget.NewLabel("Template")
	// }, func(i widget.ListItemID, o fyne.CanvasObject) {
	// 	//o.(*widget.Label).SetText(entries[i].Username)
	// 	o.(*widget.Label).SetText("Test")
	// })

	// split := container.NewHSplit(
	// 	listView,
	// 	container.NewStack(),
	// )

	// split.Offset = 0.2

	// myWindow.SetContent(split)

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
