package main

import (
	"fmt"
	"testing"
)

func TestScanner(t *testing.T) {
	expected := readFromFile("tests/fibonacci.scanner.expected")
	actual := scan("tests/examples/fibonacci.html")

	if expected != actual {
		writeToFile("tests/fibonacci.scanner.actual", actual)
		t.Fail()
	}
}

func scan(file string) string {
	var scan scanner
	scan.init(readFromFile(file))

	result := ""
	for tok := scan.nextToken(); tok != tokEOF; tok = scan.nextToken() {
		if tok == tokIntConst {
			result += fmt.Sprintf("%s->%d\n", tok, intVal)
		} else if tok == tokIdentifier || tok == tokStringConst {
			result += fmt.Sprintf("%s->%s\n", tok, strVal)
		} else {
			result += fmt.Sprintf("%s\n", tok)
		}
	}

	return result
}
