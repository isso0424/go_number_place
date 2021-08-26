package system

import (
	"isso0424/go_number_place/board"
	"isso0424/go_number_place/input"
	"isso0424/go_number_place/output"
)

type GameSystem struct {
	board     board.Board
	inputter  input.Inputter
	outputter output.Outputter
}
