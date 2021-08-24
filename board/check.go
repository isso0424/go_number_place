package board

import (
	"isso0424/go_number_place/coordinate"
)

func (b *Board) CheckConsistNumberPlace() bool {
	for _, arr := range b.numbers {
		if !checkAllValueDiffirence(arr) {
			return false
		}
	}

	tmpPlane := [][]int{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	for _, arr := range b.numbers {
		for x, v := range arr {
			tmpPlane[x] = append(tmpPlane[x], v)
		}
	}

	for _, arr := range tmpPlane {
		if !checkAllValueDiffirence(arr) {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := []int{
				b.GetNumber(coordinate.New(i*3+0, j*3+0)),
				b.GetNumber(coordinate.New(i*3+1, j*3+0)),
				b.GetNumber(coordinate.New(i*3+2, j*3+0)),
				b.GetNumber(coordinate.New(i*3+0, j*3+1)),
				b.GetNumber(coordinate.New(i*3+1, j*3+1)),
				b.GetNumber(coordinate.New(i*3+2, j*3+1)),
				b.GetNumber(coordinate.New(i*3+0, j*3+2)),
				b.GetNumber(coordinate.New(i*3+1, j*3+2)),
				b.GetNumber(coordinate.New(i*3+2, j*3+2)),
			}

			if !checkAllValueDiffirence(tmp) {
				return false
			}
		}
	}

	return true
}

func checkAllValueDiffirence(arr []int) bool {
	exist := make([]int, 9)
	for _, v := range arr {
		if v == 0 {
			continue
		}
		if exist[v-1] == 1 {
			return false
		}

		exist[v-1] = 1
	}

	return true
}
