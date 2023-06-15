package game

import (
	"bufio"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"game_with_Nikita/internal/fileWork"
	"image/color"
	"log"
	"os"
	"strconv"
)

var currentStage int

func LayoutSetter(content *fyne.Container, taskText, login string) {
	content.RemoveAll()
	filename := login + "-log.txt"

	if currentStage == 1 {
		if _, err := os.Stat(filename); err == nil {
			file, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}
			scanner := bufio.NewScanner(file)
			var lines []string
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			const stringsToKeep = 3
			if len(lines) >= stringsToKeep {
				if tmp, _ := strconv.Atoi(lines[len(lines)-1]); tmp > currentStage {
					currentStage = tmp
				}
				if lines[len(lines)-1] == "" {
					lines = lines[:len(lines)-1]
				}
				lines = lines[:len(lines)-1]
			}
			lines = append(lines, strconv.Itoa(currentStage))
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
			if err := os.Remove(filename); err != nil {
				log.Fatal(err)
			}
			file, err = os.Create(filename)
			if err != nil {
				log.Fatal(err)
			}
			writer := bufio.NewWriter(file)
			for _, line := range lines {
				if _, err := writer.WriteString(line + "\n"); err != nil {
					log.Fatal(err)
				}
			}
			if err := writer.Flush(); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		logFile := fileWork.OpenWrite(filename)
		defer fileWork.CloseFunc(logFile)
		writer := bufio.NewWriter(logFile)
		defer func(writer *bufio.Writer) {
			err := writer.Flush()
			if err != nil {
				log.Fatal(err)
			}
		}(writer)
		_, err := fmt.Fprintln(writer, currentStage)
		if err != nil {
			log.Fatal(err)
		}
	}

	stageText := canvas.NewText(strconv.Itoa(currentStage)+" Stage", color.RGBA{R: 150, G: 150, B: 150, A: 150})
	stageText.TextStyle = fyne.TextStyle{Bold: true}
	stageText.TextSize = 48
	content.Add(stageText)

	newText := widget.NewLabel(taskText)
	newText.Wrapping = fyne.TextWrapWord
	content.Add(newText)
}

func FirstStageLayout(content *fyne.Container, login string) {
	currentStage = 1

	LayoutSetter(content, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", login)
}
