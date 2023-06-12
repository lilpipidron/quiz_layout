package window

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"game_with_Nikita/internal/buttonLogic"
)

func StartWindow() {
	myApp := app.New()
	myWindow := myApp.NewWindow("qwe")

	login := widget.NewEntry()
	login.SetPlaceHolder("Enter login")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Enter password")
	content := container.NewVBox()
	signIn := widget.NewButton("sign in", func() {
		buttonLogic.SignIn(login, password, myWindow, content)
	})

	registration := widget.NewButton("registration", func() {
		buttonLogic.Registration(login, password)
	})

	buttons := container.NewGridWithColumns(10, signIn, registration)
	content = container.NewVBox(login, password, buttons)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
