package generator

import (
	"isso0424/go_number_place/board"
	"isso0424/go_number_place/coordinate"
	"math/rand"
	"time"
)

func createAnswer() board.Board {
	return board.NewWithValue(
		[][]int {
			{
				4,
				7,
				5,
				2,
				6,
				3,
				9,
				8,
				1,
			},
			{
				1,
				6,
				8,
				7,
				9,
				5,
				4,
				2,
				3,
			},
			{
				3,
				9,
				2,
				8,
				1,
				4,
				6,
				5,
				7,
			},
			{
				7,
				2,
				6,
				1,
				3,
				8,
				5,
				9,
				4,
			},
			{
				8,
				1,
				9,
				5,
				4,
				7,
				2,
				3,
				6,
			},
			{
				5,
				4,
				3,
				6,
				2,
				9,
				1,
				7,
				8,
			},
			{
				9,
				8,
				1,
				3,
				5,
				6,
				7,
				4,
				2,
			},
			{
				2,
				5,
				7,
				4,
				8,
				1,
				3,
				6,
				9,
			},
			{
				6,
				3,
				4,
				9,
				7,
				2,
				8,
				1,
				5,
			},
		},
	)
}

func generateRandomCoordinate() coordinate.Coordinate {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(9)
	y := rand.Intn(9)

	return coordinate.New(x, y)
}

func isCoverAllCoordinate(a, b []coordinate.Coordinate) bool {
	coors := []coordinate.Coordinate{}
	for _, x := range a {
		coors = append(coors, x)
	}

	for _, y := range b {
		coors = append(coors, y)
	}

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if !isIncludeCoor(coors, coordinate.New(x, y)) {
				return false
			}
		}
	}

	return true
}

func connectCoors(a, b []coordinate.Coordinate) []coordinate.Coordinate {
	coors := []coordinate.Coordinate{}
	coors = append(coors, a...)
	coors = append(coors, b...)

	return coors
}

func isIncludeCoor(coors []coordinate.Coordinate, coor coordinate.Coordinate) bool {
	for _, c := range coors {
		if c.GetX() == coor.GetX() && c.GetY() == coor.GetY() {
			return true
		}
	}

	return false
}
