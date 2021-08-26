package system

import (
	"isso0424/go_number_place/board"
	"isso0424/go_number_place/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputNumber(t *testing.T) {
	c := make(chan string, 10)
	inputter := mock.NewMockInputter([]string{"-5", "hoge", "0"})
	outputter := mock.NewMockOutputter(c)
	sys := GameSystem{
		board: board.New(),
		inputter: &inputter,
		outputter: &outputter,
	}

	result := sys.inputNumber("msg1", "msg2", func(input int) bool { return input == 0 })

	assert.Equal(t, "msg1", <-c)
	assert.Equal(t, "msg2", <-c)
	assert.Equal(t, "msg1", <-c)
	assert.Equal(t, "Please input number\n", <-c)
	assert.Equal(t, "msg1", <-c)

	assert.Equal(t, 0, result)
}
