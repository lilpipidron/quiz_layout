package buttonLogic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"os"
)

func SignIn(l, p *widget.Entry, window fyne.Window, content *fyne.Container) bool {
	login := l.Text
	password := p.Text

	wrongData := canvas.NewText("Incorrect login/password", color.RGBA{R: 255, A: 255})
	wrongData.TextStyle = fyne.TextStyle{Italic: true}

	if len(content.Objects) > 3 {
		content.Objects = content.Objects[:len(content.Objects)-1]
	}

	if !(send_login_sign(login) && send_password_sign(password)) {
		content.Add(wrongData)
		window.SetContent(content) // Обновление содержимого окна
		window.Resize(window.Canvas().Size())
		return false
	}

	return true
}

func Registration(l, p *widget.Entry) {
	login := l.Text
	password := p.Text
	_, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}

	send_login_registation(login)
	send_password_registration(password)
}
