package main

import (
	"fmt"

	"./parser"
	"./scanner"
	"./utility"
)

func main() {
	source := utility.ReadFile("tests/examples/fibonacci.html")
	scanner := scanner.NewScanner()
	tokens := scanner.Scan(source)

	parser := parser.NewParser()
	prgNode := parser.Parse(tokens)
	fmt.Print(prgNode.ToString())
}
