package buttonLogic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"os"
)

func SignIn(inputLogin, inputPassword *widget.Entry, window fyne.Window, content *fyne.Container) bool {
	login := inputLogin.Text
	password := inputPassword.Text

	wrongDataMessage := canvas.NewText("Incorrect login/password", color.RGBA{R: 255, A: 255})
	wrongDataMessage.TextStyle = fyne.TextStyle{Italic: true}

	if len(content.Objects) > 3 {
		content.Objects = content.Objects[:len(content.Objects)-1]
	}

	if !(sendLoginSign(login) && sendPasswordSign(login, password)) {
		content.Add(wrongDataMessage)
		window.SetContent(content) // Обновление содержимого окна
		window.Resize(window.Canvas().Size())
		return false
	}

	return true
}

func Registration(l, p *widget.Entry) {
	login := l.Text
	password := p.Text
	filename := login + "-log.txt"
	_, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	sendLoginRegistration(login)
	sendPasswordRegistration(login, password)
}
