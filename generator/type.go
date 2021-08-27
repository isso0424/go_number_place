package generator

import "isso0424/go_number_place/board"

type Generator interface {
	Generate() board.Board
}
