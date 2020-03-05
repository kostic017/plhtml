package scanner

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"../utility"
)

func TestScanner(t *testing.T) {
	expected := utility.ReadFile("../tests/fibonacci.scanner.expected")
	actual := scan("../tests/examples/fibonacci.html")

	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		utility.WriteFile("../tests/fibonacci.scanner.actual", actual)
		t.Fail()
	}
}

func scan(file string) string {
	source := utility.ReadFile("../tests/examples/fibonacci.html")
	scanner := NewScanner()
	tokens := scanner.Scan(source)

	result := ""
	for _, tok := range tokens {

		if tok.Type == TokEOF {
			break
		}

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

		result += fmt.Sprintf("(%s,%s)\n", tok.Type, value)

	}

	return result
}
