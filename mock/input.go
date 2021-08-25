package mock

type MockInputter struct {
	results []string
	index int
}

func NewMockInputter(results []string) MockInputter {
	return MockInputter{ results: results, index: 0 }
}
