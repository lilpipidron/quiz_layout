package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func newLayout(content *fyne.Container) {
	content.RemoveAll()
	example := canvas.NewText("test", color.RGBA{R: 255, A: 255})
	content.Add(example)
}
