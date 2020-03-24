package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"

	"./parser"
	"./scanner"
	"./token"
	"./util"
)

func TestScanner(t *testing.T) {
	tokens := scan("./tests/examples/fibonacci.html")
	compare(t, "fibonacci.scanner", tokensToString(tokens))
}

func TestParser(t *testing.T) {
	tokens := scan("./tests/examples/fibonacci.html")
	myParser := parser.New()
	prgNode := myParser.Parse(tokens)
	compare(t, "fibonacci.parser", prgNode.ToString())
}

func scan(file string) []scanner.Token {
	source := util.ReadFile(file)
	myScanner := scanner.New()
	return myScanner.Scan(source)
}

func compare(t *testing.T, testName string, actual string) {
	expected := util.ReadFile("./tests/" + testName + ".expected")
	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual), "These two should be the same.")
}

func tokensToString(tokens []scanner.Token) string {

	result := ""
	for _, tok := range tokens {

		if tok.Type == token.EOF {
			break
		}

		var value string
		switch tok.Type {
		case token.Identifier, token.StringConst:
			value = tok.StrVal
		case token.IntConst:
			value = strconv.Itoa(tok.IntVal)
		case token.RealConst:
			value = util.FloatToString(tok.RealVal)
		case token.BoolConst:
			value = strconv.FormatBool(tok.BoolVal)
		}

		result += fmt.Sprintf("(%s,%d,%d,%s)\n", token.TypeToStr[tok.Type], tok.Line, tok.Column, value)

	}

	return result
}
