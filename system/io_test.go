package system_test

import (
	"isso0424/go_number_place/board"
	"isso0424/go_number_place/mock"
	"isso0424/go_number_place/system"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputBoard(t *testing.T) {
	numbers := [][]int{
		{
			1,
			2,
			3,
			4,
			5,
			6,
			7,
			8,
			9,
		},
		{
			2,
			3,
			4,
			5,
			6,
			7,
			8,
			9,
			1,
		},
		{
			3,
			4,
			5,
			6,
			7,
			8,
			9,
			1,
			2,
		},
		{
			4,
			5,
			6,
			7,
			8,
			9,
			1,
			2,
			3,
		},
		{
			5,
			6,
			7,
			8,
			9,
			1,
			2,
			3,
			4,
		},
		{
			6,
			7,
			8,
			9,
			1,
			2,
			3,
			4,
			5,
		},
		{
			7,
			8,
			9,
			1,
			2,
			3,
			4,
			5,
			6,
		},
		{
			8,
			9,
			1,
			2,
			3,
			4,
			5,
			6,
			7,
		},
		{
			9,
			1,
			2,
			3,
			4,
			5,
			6,
			7,
			8,
		},
	}

	b := board.NewWithValue(numbers)
	c := make(chan string, 20)
	inputter := mock.NewMockInputter([]string{})
	outputter := mock.NewMockOutputter(c)
	sys := system.New(b, &inputter, &outputter)
	sys.OutputBoard()
	str := ""

	for i := 0; i < 20; i++ {
		str += <-c
	}

	expectedStr := `  0 1 2 3 4 5 6 7 8
  ===== ===== =====
0 1|2|3 4|5|6 7|8|9
  ----- ----- -----
1 2|3|4 5|6|7 8|9|1
  ----- ----- -----
2 3|4|5 6|7|8 9|1|2
  ===== ===== =====
3 4|5|6 7|8|9 1|2|3
  ----- ----- -----
4 5|6|7 8|9|1 2|3|4
  ----- ----- -----
5 6|7|8 9|1|2 3|4|5
  ===== ===== =====
6 7|8|9 1|2|3 4|5|6
  ----- ----- -----
7 8|9|1 2|3|4 5|6|7
  ----- ----- -----
8 9|1|2 3|4|5 6|7|8
  ===== ===== =====
`
	assert.Equal(t, expectedStr, str)
}
