package system

import (
	"isso0424/go_number_place/board"
	"isso0424/go_number_place/coordinate"
	"strconv"
)

func validateOneToNine(value int) bool {
	return value >= 1 && value <= 9
}

func validateZeroToEight(value int) bool {
	return value >= 0 && value <= 8
}

func(s *GameSystem) Start() {
	b := &s.board
	for (isGameOver(b)) {
		x := s.inputNumber("Input x coordinate > ", "Please input between 0 and 8 > ", validateZeroToEight)
		y := s.inputNumber("Input y coordinate > ", "Please input between 0 and 8 > ", validateZeroToEight)
		v := s.inputNumber("Input number > ", "Please input between 1 and 9 > ", validateOneToNine)

		coor := coordinate.New(x, y)

		boardError := s.board.InputNumber(v, coor)
		if boardError != nil {
			s.outputter.Output(boardError.Error() + "\n")
		}
	}
}

func(s *GameSystem) inputNumber(message, validateErrorMessage string, validate func(input int) bool) int {
		nonError := false
		validated := false
		for nonError && validated {
			s.outputter.Output(message)
			value, err := strconv.Atoi(s.inputter.Input())
			if err != nil {
				s.outputter.Output("please input number\n")
				nonError = false
				validated = false
				continue
			} else {
				nonError = true
			}
			validated = validate(value)
			if !validated {
				s.outputter.Output(validateErrorMessage)
			}
			return value
		}

		// UNREACHABLE
		return 0
}

func isGameOver(b *board.Board) bool {
	return !b.CheckConsistNumberPlace() || len(b.GetEmptyCoordinates()) == 0
}
