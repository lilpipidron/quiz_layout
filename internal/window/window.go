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
	logIn := widget.NewButton("Log in", func() {
		if buttonLogic.LogIn(login, password, myWindow, content) {
			newLayout(content)
		}
	})

	signUp := widget.NewButton("Sign up", func() {
		buttonLogic.SignUp(login, password)
	})

	buttons := container.NewGridWithColumns(10, logIn, signUp)
	content = container.NewVBox(login, password, buttons)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
