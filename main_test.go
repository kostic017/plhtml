package main

import (
	"fmt"
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

	scan := scanner.NewScanner()
	tokens := scan.Scan(source)
	assertThatScannerWorks(t, tokens)

	parser := parser.NewParser()
	prgNode := parser.Parse(tokens)
	assertThatParserWorks(t, prgNode)

}

func assertThatScannerWorks(t *testing.T, tokens []scanner.Token) {
	assert(t, "fibonacci.scanner", tokensToString(tokens))
}

func assertThatParserWorks(t *testing.T, prgNode ast.ProgramNode) {
	assert(t, "fibonacci.parser", prgNode.ToString())
}

func assert(t *testing.T, test string, actual string) {
	expected := utility.ReadFile("./tests/" + test + ".expected")

	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		utility.WriteFile("./tests/"+test+".actual", actual)
		fmt.Printf("Failed: %s\n", test)
		t.Fail()
	}
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

		result += fmt.Sprintf("(%s,%s)\n", tok.Type, value)

	}

	return result
}
