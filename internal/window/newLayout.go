package window

import (
	"bufio"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"game_with_Nikita/internal/fileWork"
	"image/color"
	"log"
	"os"
)

type State int

const (
	first State = iota + 1
	second
	third
)

func NewLayout(content *fyne.Container, login string) {
	content.RemoveAll()
	example := canvas.NewText("test", color.RGBA{R: 255, A: 255})
	content.Add(example)

	currentState := third

	filename := "log.txt"
	_, statErr := os.Stat(filename)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			_, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr)
			}
		}
	}
	logFile := fileWork.OpenWrite(filename)
	defer fileWork.CloseFunc(logFile)
	writer := bufio.NewWriter(logFile)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}(writer)
	_, err := fmt.Fprintln(writer, login)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(writer, currentState)
	if err != nil {
		log.Fatal(err)
	}
}
