package system

import (
	"fmt"
)

func(sys *GameSystem) OutputBoard() {
	numbers := sys.board.GetNumbers()

	sys.outputter.Output("-------------------\n")
	for _, arr := range numbers {
		sys.outputter.Output("|")
		for _, x := range arr {
			if x == 0 {
				sys.outputter.Output(" |")
			}
			sys.outputter.Output(fmt.Sprintf("%d|", x))
		}
		sys.outputter.Output("\n-------------------\n")
	}
}
