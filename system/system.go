package system

import "isso0424/go_number_place/board"

func(s *GameSystem) Start() {
	b := &s.board
	for (isGameOver(b)) {
		// TODO: input and output
	}
}

func isGameOver(b *board.Board) bool {
	return !b.CheckConsistNumberPlace() || len(b.GetEmptyCoordinates()) == 0
}
