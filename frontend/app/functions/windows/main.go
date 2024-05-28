package windows

import (
	"main/backend/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateEntry() *fyne.Container {
	username := widget.NewEntry()
	password := widget.NewPasswordEntry()
	generatePassword := widget.NewButton("Generate Password", func() {
		password.Text = "random"
		password.Refresh()
	})
	website := widget.NewEntry()
	notes := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username", Widget: username},
			{Text: "Password", Widget: password},
			{Text: "", Widget: generatePassword},
			{Text: "Website", Widget: website},
			{Text: "Notes", Widget: notes},
		},
		OnSubmit: func() {
			helpers.CreateEntry(username.Text, password.Text, website.Text, notes.Text)
		},
	}

	return container.NewVBox(form)
}

func GetEntries() *fyne.Container {
	entries := helpers.GetEntries("rose")
	entriesList := widget.NewList(
		func() int {
			return len(entries)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(entries[i].Username)
		},
	)

	return container.NewVBox(entriesList)
}

func MainLayout(myWindow fyne.Window, myApp fyne.App) {
	// Create a tab container
	tabs := container.NewAppTabs(
		container.NewTabItem("Create Entry", CreateEntry()),
		container.NewTabItem("Show Entries", GetEntries()),
		container.NewTabItem("Database", widget.NewLabel("Database")),
		container.NewTabItem("Settings", widget.NewLabel("Settings")),
		container.NewTabItem("About", widget.NewLabel("About")),
	)

	// Set the content of the window to the tab container
	myWindow.SetContent(tabs)

}
