package main

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "strconv"
    "strings"
    "testing"

    "./ast"
    "./parser"
    "./scanner"
    "./utility"
)

func TestCompiler(t *testing.T) {

    source := utility.ReadFile("./tests/examples/fibonacci.html")

    myScanner := scanner.New()
    tokens := myScanner.Scan(source)
    testScanner(t, tokens)

    myParser := parser.New()
    prgNode := myParser.Parse(tokens)
    testParser(t, prgNode)

}

func testScanner(t *testing.T, tokens []scanner.Token) {
    test(t, "fibonacci.scanner", tokensToString(tokens))
}

func testParser(t *testing.T, prgNode ast.ProgramNode) {
    test(t, "fibonacci.parser", prgNode.ToString())
}

func test(t *testing.T, testName string, actual string) {
    expected := utility.ReadFile("./tests/" + testName + ".expected")
    assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual), "These two should be the same.")
}

func tokensToString(tokens []scanner.Token) string {

    result := ""
    for _, tok := range tokens {

        if tok.Type == scanner.TokEOF {
            break
        }

        var value string
        switch tok.Type {
        case scanner.TokIdentifier, scanner.TokStringConst:
            value = tok.StrVal
        case scanner.TokIntConst:
            value = strconv.Itoa(tok.IntVal)
        case scanner.TokRealConst:
            value = strconv.FormatFloat(tok.RealVal, 'E', -1, 64)
        case scanner.TokBoolConst:
            value = strconv.FormatBool(tok.BoolVal)
        }

        result += fmt.Sprintf("(%s,%d,%d,%s)\n", tok.Type, tok.Line, tok.Column, value)

    }

    return result
}
