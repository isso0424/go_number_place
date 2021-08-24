package board

func New() Board {
	return Board{
		numbers: [][]int{
			{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
			{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
			{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
			{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
			{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
			{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
			{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
			{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
			{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
		},
	}
}

func NewWithValue(plane [][]int) Board {
	if len(plane) != 9 {
		panic("WHY DO YOU WANNA PLAY NUMBER PLACE WITHOUT 9 X 9 BOARD????")
	}

	for _, arr := range(plane) {
		if len(arr) != 9 {
			panic("WHY DO YOU WANNA PLAY NUMBER PLACE WITHOUT 9 X 9 BOARD????")
		}
	}

	return Board{ numbers: plane }
}
