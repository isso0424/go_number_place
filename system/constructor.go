package system

import "isso0424/go_number_place/board"

func New(board board.Board) GameSystem {
	return GameSystem{ board }
}
