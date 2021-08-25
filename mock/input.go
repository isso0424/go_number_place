package mock

type MockInputter struct {
	results []string
	index int
}

func NewMockInputter(results []string) MockInputter {
	return MockInputter{ results: results, index: 0 }
}

func(inputter *MockInputter) Input() string {
	if len(inputter.results) == inputter.index {
		panic("Results over flow")
	}

	result := inputter.results[inputter.index]
	inputter.index++;

	return result
}
