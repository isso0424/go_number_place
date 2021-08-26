package mock

type mockOutputter struct {
	dist chan string
}

func NewMockOutputter(dist chan string) mockOutputter {
	return mockOutputter{dist}
}

func (outputter *mockOutputter) Output(message string) {
	outputter.dist <- message
}
