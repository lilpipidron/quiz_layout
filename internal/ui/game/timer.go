package game

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"time"
)

func Timer(stop <-chan bool, minutes int, seconds int, container *fyne.Container) {
	clock := widget.NewLabel("")

	for i := minutes*60 + seconds; i > 0; i-- {
		if <-stop {
			break
		}
		min := i / 60
		sec := i % 60
		text := fmt.Sprintf("Time left: %d:%d", min, sec)
		clock.SetText(text)
		container.Add(clock)
		time.Sleep(time.Second)
		container.Remove(clock)
	}

}
