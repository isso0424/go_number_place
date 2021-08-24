package board_test

import (
	"isso0424/go_number_place/board"
	"isso0424/go_number_place/coordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumber(t *testing.T) {
	board1 := board.NewWithValue(
		[][]int{
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
		},
	)

	num := board1.GetNumber(coordinate.New(0, 0))
	assert.Equal(t, 1, num)
}

func TestGetEmpty(t *testing.T) {
	board1 := board.NewWithValue(
		[][]int{
			{
				0,
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
				4,
				5,
				6,
				0,
				8,
				9,
				1,
				2,
				3,
			},
			{
				7,
				8,
				9,
				1,
				2,
				3,
				4,
				0,
				6,
			},
			{
				0,
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
				5,
				6,
				7,
				8,
				9,
				1,
				2,
				3,
				0,
			},
			{
				0,
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
				3,
				4,
				5,
				6,
				7,
				8,
				0,
				1,
				2,
			},
			{
				6,
				7,
				8,
				0,
				1,
				2,
				3,
				4,
				5,
			},
			{
				9,
				1,
				2,
				0,
				4,
				5,
				0,
				7,
				8,
			},
		},
	)

	coordinates := board1.GetEmptyCoordinates()
	assert.Equal(t, 10, len(coordinates))
}

func TestNewBoard(t *testing.T) {
	board1 := board.New()

	coordinates := board1.GetEmptyCoordinates()
	assert.Equal(t, 81, len(coordinates))
}

func TestInputBoard(t *testing.T) {
	board1 := board.NewWithValue(
		[][]int{
			{
				0,
				0,
				3,
				4,
				5,
				6,
				7,
				8,
				9,
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
		},
	)

	err := board1.InputNumber(2, coordinate.New(0, 0))
	assert.Equal(t, 2, board1.GetNumber(coordinate.New(0, 0)))
	assert.Equal(t, nil, err)

	err = board1.InputNumber(3, coordinate.New(0, 0))
	assert.Equal(t, 2, board1.GetNumber(coordinate.New(0, 0)))
	assert.NotEqual(t, nil, err)
}

func TestConsistNumberPlace(t *testing.T) {
	board1 := board.New()
	assert.Equal(t, true, board1.CheckConsistNumberPlace())

	board1.InputNumber(1, coordinate.New(0, 0))
	assert.Equal(t, true, board1.CheckConsistNumberPlace())

	board2 := board.NewWithValue(board1.GetNumbers())
	board2.InputNumber(1, coordinate.New(1, 1))
	assert.Equal(t, false, board2.CheckConsistNumberPlace())

	board3 := board.NewWithValue(board1.GetNumbers())
	board3.InputNumber(1, coordinate.New(3, 0))
	assert.Equal(t, false, board3.CheckConsistNumberPlace())

	board4 := board.NewWithValue(board1.GetNumbers())
	board4.InputNumber(1, coordinate.New(0, 3))
	assert.Equal(t, false, board4.CheckConsistNumberPlace())
}
