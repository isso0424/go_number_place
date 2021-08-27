package generator

import (
	"isso0424/go_number_place/board"
	"isso0424/go_number_place/coordinate"
	"math/rand"
	"time"
)

type EasyGenerator struct {}

func(g *EasyGenerator) Generate() (problem, answer board.Board) {
	rand.Seed(time.Now().UnixNano())
	answer = createAnswer()
	selected := []coordinate.Coordinate{}
	used := []coordinate.Coordinate{}

	for !isCoverAllCoordinate(selected, used) {
		coor := generateRandomCoordinate()
		if isIncludeCoor(connectCoors(selected, used), coor) {
			continue
		}

		selected = append(selected, coor)
		used = append(used, g.usedCoors(answer, coor, answer.GetNumber(coor))...)
	}

	tmp := board.New()
	probNum := tmp.GetNumbers()
	ansNum := answer.GetNumbers()
	for _, coor := range selected {
		probNum[coor.GetY()][coor.GetX()] = ansNum[coor.GetY()][coor.GetX()]
	}

	problem = board.NewWithValue(probNum)

	return
}

func(g *EasyGenerator) usedCoors(b board.Board, coor coordinate.Coordinate, number int) []coordinate.Coordinate {
	coors := []coordinate.Coordinate{}
	x := coor.GetX()
	y := coor.GetY()
	xgroup := x / 3
	ygroup := y / 3

	numbers := b.GetNumbers()

	isInTarget := func (c coordinate.Coordinate) bool {
		if c.GetX() == x || c.GetY() == y {
			return false
		}
		return c.GetX() / 3 == xgroup || c.GetY() == ygroup
	}

	for yIndex, arr := range numbers {
		for xIndex, value := range arr {
			coor := coordinate.New(xIndex, yIndex)
			if number == value && isInTarget(coor) {
				coors = append(coors, coor)
			}
		}
	}

	return coors
}
