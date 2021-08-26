package mock_test

import (
	"isso0424/go_number_place/input"
	"isso0424/go_number_place/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	values := []string{"hoge", "fuga", "foo", "bar"}

	inputter := mock.NewMockInputter(values)

	assert.Equal(t, "hoge", inputter.Input())
	assert.Equal(t, "fuga", inputter.Input())
	assert.Equal(t, "foo", inputter.Input())
	assert.Equal(t, "bar", inputter.Input())

	validate := func(i input.Inputter) bool {
		return true
	}

	assert.Equal(t, true, validate(&inputter))

	defer func() {
		err := recover()
		s := "Results over flow"
		if err != s {
			t.Errorf("want: %s actual: %s\n", s, err)
		}
	}()

	inputter.Input()
}
