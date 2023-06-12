package buttonLogic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"os"
)

func LogIn(inputLogin, inputPassword *widget.Entry, window fyne.Window, content *fyne.Container) bool {
	login := inputLogin.Text
	password := inputPassword.Text

	wrongDataMessage := canvas.NewText("Incorrect login/password", color.RGBA{R: 255, A: 255})
	wrongDataMessage.TextStyle = fyne.TextStyle{Italic: true}

	if len(content.Objects) > 3 {
		content.Objects = content.Objects[:len(content.Objects)-1]
	}

	if !(sendLoginLogIn(login) && sendPasswordLogIn(login, password)) {
		content.Add(wrongDataMessage)
		window.SetContent(content) // Обновление содержимого окна
		window.Resize(window.Canvas().Size())
		return false
	}

	return true
}

func SignUp(inputLogin, inputPassword *widget.Entry) {
	login := inputLogin.Text
	password := inputPassword.Text
	filename := login + "-log.txt"

	_, statErr := os.Stat(filename)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			_, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr)
			}
			sendLoginSignUp(login)
			sendPasswordSignUp(login, password)
		}
	}
	log.SetOutput(os.Stderr)
	log.Println("This username already exists")
}
