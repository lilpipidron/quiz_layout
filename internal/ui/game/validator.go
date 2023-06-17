package game

import "game_with_Nikita/consts"

func validate(stage int, answer string) bool {
	switch stage {
	case 1:
		return answer == consts.FirstAnswer
	default:
		return false
	}
}
