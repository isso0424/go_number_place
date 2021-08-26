package system

import (
	"fmt"
	"strconv"
)

func(sys *GameSystem) OutputBoard() {
	numbers := sys.board.GetNumbers()

	sys.outputter.Output("-------------------\n")
	for _, arr := range numbers {
		str := "|"
		for _, x := range arr {
			if x == 0 {
				str += " |"
			}
			str += fmt.Sprintf("%d|", x)
		}
		sys.outputter.Output(str)
		sys.outputter.Output("\n-------------------\n")
	}
}

func(s *GameSystem) inputNumber(message, validateErrorMessage string, validate func(input int) bool) int {
		nonError := false
		validated := false
		for !nonError || !validated {
			s.outputter.Output(message)
			value, err := strconv.Atoi(s.inputter.Input())
			if err != nil {
				s.outputter.Output("Please input number\n")
				nonError = false
				validated = false
				continue
			} else {
				nonError = true
			}
			validated = validate(value)
			if !validated {
				s.outputter.Output(validateErrorMessage)
				continue
			}

			return value
		}

		// UNREACHABLE
		return 0
}
