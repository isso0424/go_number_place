package coordinate

func New(x, y int) Coordinate {
	return Coordinate{x, y}
}

func CreateFromValue(plane [][]int, value int) ([]Coordinate, CoordinateError) {
	coordinates := []Coordinate{}
	for y, arr := range plane {
		for x, v := range arr {
			if value == v {
				coordinates = append(coordinates, New(x, y))
			}
		}
	}

	if len(coordinates) == 0 {
		return coordinates, ValueNotFound{value}
	}

	return coordinates, nil
}

func CreateFromValueOnce(plane [][]int, value int) (Coordinate, CoordinateError) {
	v, err := CreateFromValue(plane, value)
	if err != nil {
		return Coordinate{}, nil
	}

	return v[0], err
}
