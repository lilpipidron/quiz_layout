package game

const (
	firstAnswer = "123"
)

func validate(stage int, answer string) bool {
	switch stage {
	case 1:
		return answer == firstAnswer
	default:
		return false
	}
}
