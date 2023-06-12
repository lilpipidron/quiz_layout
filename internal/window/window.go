package window

import (
	"GolandProjects/game_with_Nikita/internal/buttonLogic"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func StartWindow() {
	a := app.New()
	w := a.NewWindow("qwe")
	click := widget.NewButton("Test", buttonLogic.Test)
	w.SetContent(click)
	w.ShowAndRun()
}
