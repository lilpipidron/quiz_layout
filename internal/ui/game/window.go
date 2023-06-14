package game

import (
	"bufio"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"game_with_Nikita/internal/file_work"
	"image/color"
	"log"
	"strconv"
)

const (
	First int = iota + 1
)

var currentState int

func LayoutSetter(stage int, content *fyne.Container, taskText, login string) {
	content.RemoveAll()
	currentState = stage
	stageText := canvas.NewText(strconv.Itoa(stage)+" Stage", color.RGBA{R: 150, G: 150, B: 150, A: 150})
	stageText.TextStyle = fyne.TextStyle{Bold: true}
	stageText.TextSize = 48
	content.Add(stageText)

	newText := widget.NewLabel(taskText)
	newText.Wrapping = fyne.TextWrapWord
	content.Add(newText)

	filename := login + "-log.txt"
	logFile := file_work.OpenWrite(filename)
	defer file_work.CloseFunc(logFile)
	writer := bufio.NewWriter(logFile)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}(writer)
	_, err := fmt.Fprintln(writer, currentState)
	if err != nil {
		log.Fatal(err)
	}
}

func FirstStageLayout(content *fyne.Container, login string) {
	stage := First

	LayoutSetter(stage, content, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", login)
}
