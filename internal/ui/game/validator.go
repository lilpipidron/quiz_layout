package game

import "game_with_Nikita/consts"

func validate(stage int, answer string) bool {
	switch stage {
	case 1:
		return answer == consts.FirstAnswer
	case 2:
		return answer == consts.SecondAnswer
	default:
		return false
	}
}
