package main

import (
	"fmt"

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
	fmt.Print(prgNode.ToString())

	symbolTable := semantic.NewSymbolTable()
	semanticVisitor := semantic.NewSemanticVisitor(symbolTable)
	prgNode.Accept(semanticVisitor)

}
