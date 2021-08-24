package board

import "isso0424/go_number_place/coordinate"

func(b *Board) getNumbers() [][]int {
	return b.numbers
}

func(b *Board) getEmptyCoordinates() []coordinate.Coordinate {
	results := []coordinate.Coordinate{}
	for outerIndex, arr := range(b.numbers) {
		for innerIndex, value := range(arr) {
			if value == 0 {
				results = append(results, coordinate.New(innerIndex, outerIndex))
			}
		}
	}

	return results
}
