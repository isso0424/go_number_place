package main

import (
	"isso0424/go_number_place/console"
	"isso0424/go_number_place/generator"
	"isso0424/go_number_place/system"
)

func main() {
	gen := generator.EasyGenerator{}
	problem, _ := gen.Generate()
	inputter := console.NewInputter()
	outputter := console.NewOutputter()
	sys := system.New(problem, &inputter, &outputter)
	sys.Start()
}
