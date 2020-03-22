package main

import (
    "./logger"
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

    semantic.SetLogLevel(logger.Debug)
    semanticAnalyzer := semantic.NewAnalyzer()
    prgNode.Accept(semanticAnalyzer)

}
