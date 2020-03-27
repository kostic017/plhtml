package main

import (
    "os"
    "plhtml/interpreter"
    "plhtml/parser"
    "plhtml/scanner"
    "plhtml/semantic"
    "plhtml/util"
)

func main() {
    source := util.ReadFile(os.Args[1])

    myScanner := scanner.New()
    tokens := myScanner.Scan(source)

    myParser := parser.New()
    prgNode := myParser.Parse(tokens)

    analyzer := semantic.NewAnalyzer()
    prgNode.AcceptAnalyzer(analyzer)

    interp := interpreter.New()
    prgNode.AcceptInterpreter(interp)
}
