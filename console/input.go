package console

import (
	"bufio"
	"os"
)

type consoleInputter struct{}

func NewInputter() consoleInputter {
	return consoleInputter{}
}

func (inputter *consoleInputter) Input() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	return stdin.Text()
}
