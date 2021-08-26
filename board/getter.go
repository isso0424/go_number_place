package board

import "isso0424/go_number_place/coordinate"

func (b *Board) GetNumbers() [][]int {
	numbers := [][]int{}
	for _, arr := range b.numbers {
		tmp := []int{}
		for _, v := range arr {
			tmp = append(tmp, v)
		}
		numbers = append(numbers, tmp)
	}

	return numbers
}

func (b *Board) GetNumber(coordinate coordinate.Coordinate) int {
	return b.numbers[coordinate.GetY()][coordinate.GetX()]
}

func (b *Board) GetEmptyCoordinates() []coordinate.Coordinate {
	results := []coordinate.Coordinate{}
	for outerIndex, arr := range b.numbers {
		for innerIndex, value := range arr {
			if value == 0 {
				results = append(results, coordinate.New(innerIndex, outerIndex))
			}
		}
	}

	return results
}
