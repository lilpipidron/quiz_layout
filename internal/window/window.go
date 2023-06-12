package window

import (
	"fyne.io/fyne/v2/app"
)

func StartWindow() {
	a := app.New()
	w := a.NewWindow("qwe")

	w.ShowAndRun()

}
