package generator

import (
	"isso0424/go_number_place/board"
	"math/rand"
	"time"
)

func createAnswer() board.Board {
	rand.Seed(time.Now().UnixNano())
	number := [][]int{}
	for i := 0; i < 9; i++ {
		tmp := []int{}
		for j := 0; j < 9; j++ {
			tmp = append(tmp, rand.Intn(9) + 1)
		}
		number = append(number, tmp)
	}

	return board.NewWithValue(number)
}
