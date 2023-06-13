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
	First State = iota + 1
)

func SecondLayout(content *fyne.Container, login string) {
	content.RemoveAll()
	example := canvas.NewText("test", color.RGBA{R: 255, B: 255, A: 255})
	content.Add(example)

	currentState := First

	filename := "log.txt"
	err := os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	logFile := fileWork.OpenWrite(filename)
	defer fileWork.CloseFunc(logFile)
	writer := bufio.NewWriter(logFile)
	defer func(writer *bufio.Writer) {
		err = writer.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}(writer)
	_, err = fmt.Fprintln(writer, login)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(writer, currentState)
	if err != nil {
		log.Fatal(err)
	}
}
