package main

import (
    "./interpreter"
    "./parser"
    "./scanner"
    "./semantic"
    "./util"
)

func main() {

    source := util.ReadFile("tests/examples/fibonacci.html")

    myScanner := scanner.New()
    tokens := myScanner.Scan(source)

    myParser := parser.New()
    prgNode := myParser.Parse(tokens)

    //semantic.SetLogLevel(logger.Debug)
    semanticAnalyzer := semantic.NewAnalyzer()
    prgNode.Accept(semanticAnalyzer)

    //interpreter.SetLogLevel(logger.Debug)
    interp := interpreter.New()
    prgNode.Accept(interp)
}
