package scanner

import (
	"fmt"
	"testing"

	"../utility"
)

func TestScanner(t *testing.T) {
	expected := utility.ReadFile("../tests/fibonacci.scanner.expected")
	actual := scan("../tests/examples/fibonacci.html")

	if expected != actual {
		utility.WriteFile("../tests/fibonacci.scanner.actual", actual)
		t.Fail()
	}
}

func scan(file string) string {
	var scan Scanner
	scan.init(utility.ReadFile(file))

	result := ""
	for tok := scan.nextToken(); tok.Type != EOF; tok = scan.nextToken() {
		if tok.Type == Identifier || tok.Type == IntConst || tok.Type == RealConst || tok.Type == BoolConst || tok.Type == StringConst {
			result += fmt.Sprintf("%s->%s\n", tok.Type, tok.Value)
		} else {
			result += fmt.Sprintf("%s\n", tok.Type)
		}
	}

	return result
}
