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

	loginEntry := widget.NewEntry()
	loginEntry.SetPlaceHolder("Enter login")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")

	content := container.NewVBox()
	logIn := widget.NewButton("Log in", func() {
		loggedIn, loginName := buttonLogic.LogIn(loginEntry, passwordEntry, myWindow, content)
		if loggedIn {
			SecondLayout(content, loginName)
		}
	})

	signUp := widget.NewButton("Sign up", func() {
		buttonLogic.SignUp(loginEntry, passwordEntry, myWindow, content)
	})

	buttons := container.NewGridWithColumns(10, logIn, signUp)
	content = container.NewVBox(loginEntry, passwordEntry, buttons)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
