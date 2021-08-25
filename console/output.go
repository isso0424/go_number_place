package console

import "fmt"

type consoleOutputter struct {}

func(outputter *consoleOutputter) Output(message string) {
	fmt.Print(message)
}
