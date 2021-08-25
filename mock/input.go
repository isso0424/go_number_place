package mock

type mockInputter struct {
	results []string
	index int
}

func NewMockInputter(results []string) mockInputter {
	return mockInputter{ results: results, index: 0 }
}

func(inputter *mockInputter) Input() string {
	if len(inputter.results) == inputter.index {
		panic("Results over flow")
	}

	result := inputter.results[inputter.index]
	inputter.index++;

	return result
}
