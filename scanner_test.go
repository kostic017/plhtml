package main

import (
	"fmt"
	"strconv"
	"testing"

	"./utility"
)

func TestScanner(t *testing.T) {
	expected := utility.ReadFile("tests/fibonacci.scanner.expected")
	actual := scan("tests/examples/fibonacci.html")

	if expected != actual {
		utility.WriteFile("tests/fibonacci.scanner.actual", actual)
		t.Fail()
	}
}

func scan(file string) string {
	scan := NewScanner(utility.ReadFile(file))

	result := ""
	for tok := scan.NextToken(); tok.Type != TokEOF; tok = scan.NextToken() {

		var value string
		switch tok.Type {
		case TokIdentifier, TokStringConst:
			value = tok.StrVal
		case TokIntConst:
			value = strconv.Itoa(tok.IntVal)
		case TokRealConst:
			value = strconv.FormatFloat(tok.RealVal, 'E', -1, 64)
		case TokBoolConst:
			value = strconv.FormatBool(tok.BoolVal)
		}

		if value == "" {
			result += fmt.Sprintf("%s\n", tok.Type)
		} else {
			result += fmt.Sprintf("%s->%s\n", tok.Type, value)
		}
	}

	return result
}
