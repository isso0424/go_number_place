package mock_test

import (
	"isso0424/go_number_place/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	channel := make(chan string, 3)
	outputter := mock.NewMockOutputter(channel)
	outputter.Output("hoge")

	assert.Equal(t, "hoge", <-channel)
	outputter.Output("fuga")
	outputter.Output("foo")
	outputter.Output("bar")

	assert.Equal(t, "fuga", <-channel)
	assert.Equal(t, "foo", <-channel)
	assert.Equal(t, "bar", <-channel)
}
