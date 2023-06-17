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

type LayoutData struct {
	stage     int
	username  string
	taskText  string
	nextStage bool
}

func LayoutSwitcher(content *fyne.Container) {
	content.RemoveAll()
	//todo
}

func LayoutSetter(content *fyne.Container, data LayoutData) {
	LayoutSwitcher(content)
	filename := data.username + "-log.txt"

	if data.stage == 1 {
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
				if tmp, _ := strconv.Atoi(lines[len(lines)-1]); tmp > data.stage {
					data.stage = tmp
				}
				if lines[len(lines)-1] == "" {
					lines = lines[:len(lines)-1]
				}
				lines = lines[:len(lines)-1]
			}
			lines = append(lines, strconv.Itoa(data.stage))
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
		_, err := fmt.Fprintln(writer, data.stage)
		if err != nil {
			log.Fatal(err)
		}
	}

	stageText := canvas.NewText(strconv.Itoa(data.stage)+" Stage", color.RGBA{R: 150, G: 150, B: 150, A: 150})
	stageText.TextStyle = fyne.TextStyle{Bold: true}
	stageText.TextSize = 48
	content.Add(stageText)

	newText := widget.NewLabel(data.taskText)
	newText.Wrapping = fyne.TextWrapWord
	content.Add(newText)

	answerEntry := widget.NewEntry()
	answerEntry.SetPlaceHolder("Enter answer")
	content.Add(answerEntry)

	nextStageButton := widget.NewButton("Next stage", func() {
		if data.nextStage {
		}
	})
	submitButton := widget.NewButton("Submit", func() {
		if validate(data.stage, answerEntry.Text) && !data.nextStage {
			content.Add(nextStageButton)
			data.nextStage = true
		}
	})

	content.Add(submitButton)
}

func Layout(content *fyne.Container, login string) {
	task := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	firstStageData := LayoutData{1, login, task, false}

	LayoutSetter(content, firstStageData)
}
