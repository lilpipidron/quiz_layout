package game

func getTaskText(stage int) string {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	return text
}

func assignTasks(stages []StageData, maxStage int) {
	for i := 0; i < len(stages); i++ {
		stages[i].number = i + 1
		stages[i].taskText = getTaskText(stages[i].number)
		stages[i].nextStageAvailable = false
	}

	for i := 0; i < maxStage-1; i++ {
		stages[i].nextStageAvailable = true
	}
}
