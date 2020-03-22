package main

import (
    "./parser"
    "./scanner"
    "./semantic"
    "./utility"
)

func main() {

    source := utility.ReadFile("tests/examples/fibonacci.html")

    myScanner := scanner.New()
    tokens := myScanner.Scan(source)

    myParser := parser.New()
    prgNode := myParser.Parse(tokens)

    semanticAnalyzer := semantic.NewAnalyzer()
    prgNode.Accept(semanticAnalyzer)

}
