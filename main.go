package main

import (
	"./parser"
	"./scanner"
	"./semantic"
	"./utility"
)

func main() {

	source := utility.ReadFile("tests/examples/fibonacci.html")

	scanner := scanner.NewScanner()
	tokens := scanner.Scan(source)

	parser := parser.NewParser()
	prgNode := parser.Parse(tokens)

	semanticAnalyzer := semantic.NewSemanticAnalyzer()
	prgNode.Accept(semanticAnalyzer)

}
