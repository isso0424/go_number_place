package system

import (
	"fmt"
	"strconv"
)

func (sys *GameSystem) OutputBoard() {
	numbers := sys.board.GetNumbers()

	sys.outputter.Output("  0 1 2 3 4 5 6 7 8\n")
	sys.outputter.Output("  ===== ===== =====\n")
	for i, arr := range numbers {
		str := fmt.Sprintf("%d ", i)
		for j, x := range arr {
			if x == 0 {
				str += " "
			} else {
				str += fmt.Sprintf("%d", x)
			}
			if j == 8 {
				str += ""
			} else if j % 3 == 2 {
				str += " "
			} else {
				str += "|"
			}
		}
		sys.outputter.Output(str)
		if i % 3 == 2 {
			sys.outputter.Output("\n  ===== ===== =====\n")
			continue
		}
		sys.outputter.Output("\n  ----- ----- -----\n")
	}
}

func (s *GameSystem) inputNumber(message, validateErrorMessage string, validate func(input int) bool) int {
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
