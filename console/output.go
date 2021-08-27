package console

import "fmt"

type consoleOutputter struct{}

func NewOutputter() consoleOutputter {
	return consoleOutputter{}
}

func (outputter *consoleOutputter) Output(message string) {
	fmt.Print(message)
}
