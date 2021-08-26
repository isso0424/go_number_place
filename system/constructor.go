package system

import (
	"isso0424/go_number_place/board"
	"isso0424/go_number_place/input"
	"isso0424/go_number_place/output"
)

func New(board board.Board, inputter input.Inputter, outputter output.Outputter) GameSystem {
	return GameSystem{board, inputter, outputter}
}
