package window

import (
	"GolandProjects/game_with_Nikita/internal/buttonLogic"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartWindow() {
	a := app.New()
	w := a.NewWindow("qwe")

	login := widget.NewEntry()
	login.SetPlaceHolder("Enter login")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Enter password")

	sign_in := widget.NewButton("sign in", func() {
		buttonLogic.Sign_in(login, password)
	})

	registration := widget.NewButton("registration", func() {
		buttonLogic.Registration(login, password)
	})

	buttons := container.NewGridWithColumns(10, sign_in, registration)
	content := container.NewVBox(login, password, buttons)
	w.SetContent(content)
	w.ShowAndRun()
}
