package board

import "isso0424/go_number_place/coordinate"

func checkAllValueDiffirence(arr []int) bool {
	exist := make([]int, 9)
	for  v := range(arr) {
		if v == 0 {
			continue
		}
		if exist[v - 1] == 1{
			return false
		}

		exist[v - 1] = 1
	}

	return true
}
