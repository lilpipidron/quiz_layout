package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"game_with_Nikita/internal/logs"
	"image/color"
	"log"
	"os"
)

func LogIn(loginEntry, passwordEntry *widget.Entry, window fyne.Window, content *fyne.Container) (bool, string) {
	login := loginEntry.Text
	password := passwordEntry.Text

	wrongDataMessage := canvas.NewText("Incorrect login or password", color.RGBA{R: 255, A: 255})
	wrongDataMessage.TextStyle = fyne.TextStyle{Italic: true}

	if len(content.Objects) > 3 {
		content.Objects = content.Objects[:len(content.Objects)-1]
	}

	invalidLogin := false
	filename := login + "-log.txt"
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			invalidLogin = true
		} else {
			log.Fatal(err)
		}
	}

	if invalidLogin || !(logs.SendLoginLogIn(login) && logs.SendPasswordLogIn(filename, password)) {
		content.Add(wrongDataMessage)
		window.SetContent(content)
		window.Resize(window.Canvas().Size())
		return false, ""
	}

	return true, login
}

func SignUp(loginEntry, passwordEntry *widget.Entry, window fyne.Window, content *fyne.Container) {
	login := loginEntry.Text
	password := passwordEntry.Text
	filename := login + "-log.txt"

	if len(content.Objects) > 3 {
		content.Objects = content.Objects[:len(content.Objects)-1]
	}

	var wrongDataMessage *canvas.Text

	if len(password) < 8 {
		wrongDataMessage = canvas.NewText("Use 8 characters or more for your password", color.RGBA{R: 255, A: 255})
		wrongDataMessage.TextStyle = fyne.TextStyle{Italic: true}
	} else {
		_, statErr := os.Stat(filename)
		if statErr != nil {
			if os.IsNotExist(statErr) {
				_, createErr := os.Create(filename)
				if createErr != nil {
					log.Fatal(createErr)
				}
				logs.SendLoginSignUp(login)
				logs.SendPasswordSignUp(login, password)
			}
		} else {
			wrongDataMessage = canvas.NewText("This username already exists", color.RGBA{R: 255, A: 255})
			wrongDataMessage.TextStyle = fyne.TextStyle{Italic: true}
		}
	}
	if wrongDataMessage != nil {
		content.Add(wrongDataMessage)
		window.SetContent(content)
		window.Resize(window.Canvas().Size())
	}
}

func NextStage() {
	//todo
}
