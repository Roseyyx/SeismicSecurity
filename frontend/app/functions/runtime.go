package functions

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Setup() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	myWindow.Resize(fyne.NewSize(800, 600))

	listView := widget.NewList(func() int {
		return 10
	}, func() fyne.CanvasObject {
		return widget.NewLabel("Template")
	}, func(i widget.ListItemID, o fyne.CanvasObject) {
		o.(*widget.Label).SetText("Template " + fmt.Sprint(i))
	})

	split := container.NewHSplit(
		listView,
		container.NewStack(),
	)

	split.Offset = 0.2

	myWindow.SetContent(split)

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
