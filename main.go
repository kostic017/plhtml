package main

import (
	"./parser"
	"./scanner"
	"./utility"
)

func main() {
	source := utility.ReadFile("tests/examples/fibonacci.html")
	scanner := scanner.NewScanner()
	tokens := scanner.Scan(source)

	parser := parser.NewParser()
	parser.Parse(tokens)
}
