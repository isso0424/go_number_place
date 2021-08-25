package console

import (
	"bufio"
	"fmt"
	"os"
)

type consoleInputter struct {}

func NewInputter() consoleInputter {
	return consoleInputter{}
}

func(inputter *consoleInputter) Input(message string) string {
	fmt.Print(message)
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	return stdin.Text()
}
