package functions

import (
	"fmt"
	"main/frontend/app/functions/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Setup() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	myWindow.Resize(fyne.NewSize(800, 600))

	windows.Check(myWindow)

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
