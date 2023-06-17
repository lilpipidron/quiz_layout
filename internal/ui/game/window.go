package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

const stageAmount = 1

type StageData struct {
	number             int
	taskText           string
	nextStageAvailable bool
}

func LayoutSwitcher(content *fyne.Container) {
	content.RemoveAll()
	//todo
}

func LayoutSetter(content *fyne.Container, stages []StageData, currentStage int) {
	LayoutSwitcher(content)

	stageText := canvas.NewText(strconv.Itoa(stages[currentStage].number)+" Stage", color.RGBA{R: 150, G: 150, B: 150, A: 150})
	stageText.TextStyle = fyne.TextStyle{Bold: true}
	stageText.TextSize = 48
	content.Add(stageText)

	newText := widget.NewLabel(stages[currentStage].taskText)
	newText.Wrapping = fyne.TextWrapWord
	content.Add(newText)

	answerEntry := widget.NewEntry()
	answerEntry.SetPlaceHolder("Enter answer")
	content.Add(answerEntry)

	nextStageButton := widget.NewButton("Next stage", func() {
		LayoutSetter(content, stages, currentStage+1)
	})

	submitButton := widget.NewButton("Submit", func() {
		if validate(stages[currentStage].number, answerEntry.Text) && !stages[currentStage].nextStageAvailable {
			content.Add(nextStageButton)
			stages[currentStage].nextStageAvailable = true
		}
	})

	content.Add(submitButton)
}

func Layout(content *fyne.Container, login string) {
	stages := make([]StageData, stageAmount)
	for i, stage := range stages {
		stage.number = i + 1
		stage.taskText = GetTaskText(stage.number)
		stage.nextStageAvailable = false
	}

	maxStage := GetSaveData(login)
	for i := 0; i < maxStage-1; i++ {
		stages[i].nextStageAvailable = true
	}

	LayoutSetter(content, stages, maxStage-1)
}
