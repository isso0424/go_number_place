package board

import "isso0424/go_number_place/coordinate"

func(b *Board) inputNumber(value int, coor coordinate.Coordinate) BoardError {
	prev := b.getNumber(coor)
	if prev != 0 {
		return AlreadyInputted{ coor }
	}

	b.numbers[coor.GetY()][coor.GetX()] = value

	return nil
}
