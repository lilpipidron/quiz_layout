package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

const stageAmount = 2

type StageData struct {
	number             int
	taskText           string
	nextStageAvailable bool
}

func layoutSetter(content *fyne.Container, stages []StageData, login string, currentStage *int) {
	content.RemoveAll()

	if *currentStage != stageAmount+1 {
		stageText := canvas.NewText(strconv.Itoa(stages[*currentStage-1].number)+" Stage", color.RGBA{R: 150, G: 150, B: 150, A: 150})
		stageText.TextStyle = fyne.TextStyle{Bold: true}
		stageText.TextSize = 48
		content.Add(stageText)

		newText := widget.NewLabel(stages[*currentStage-1].taskText)
		newText.Wrapping = fyne.TextWrapWord
		content.Add(newText)

		if !stages[*currentStage-1].nextStageAvailable {
			answerEntry := widget.NewEntry()
			answerEntry.SetPlaceHolder("Enter answer")
			content.Add(answerEntry)
			nextStageButton := widget.NewButton("Next stage", func() {
				*currentStage = *currentStage + 1
				*currentStage = getSaveData(login, *currentStage)
				layoutSetter(content, stages, login, currentStage)
			})

			submitButton := widget.NewButton("Submit", func() {
				if validate(stages[*currentStage-1].number, answerEntry.Text) && !stages[*currentStage-1].nextStageAvailable {
					content.Add(nextStageButton)
					stages[*currentStage-1].nextStageAvailable = true
				} else {
					if len(content.Objects) == 4 { // fixme: number will vary depending on stage content
						wrongAnswerMessage := canvas.NewText("Wrong answer!", color.RGBA{R: 255, A: 255})
						wrongAnswerMessage.TextStyle = fyne.TextStyle{Italic: true}
						content.Add(wrongAnswerMessage)
					}
				}
			})

			content.Add(submitButton)
		}
	} else {
		stageText := canvas.NewText("Congratulations!", color.RGBA{R: 150, G: 150, B: 150, A: 150})
		stageText.TextStyle = fyne.TextStyle{Bold: true}
		stageText.TextSize = 48
		content.Add(stageText)

		newText := widget.NewLabel("Congratulations! You answered all the questions correctly!\nWe are impressed by your outstanding performance and knowledge. You are clearly well-informed and very intelligent.\nYou should feel proud of this amazing achievement.")
		newText.Wrapping = fyne.TextWrapWord
		content.Add(newText)
	}
}

func DefaultLayout(content *fyne.Container, login string) {
	stages := make([]StageData, stageAmount)
	maxAvailableStage := 0
	maxAvailableStage = getSaveData(login, maxAvailableStage)
	assignTasks(stages, maxAvailableStage)
	layoutSetter(content, stages, login, &maxAvailableStage)
}
