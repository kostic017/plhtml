package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/kylelemons/godebug/diff"
)

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

	return strings.TrimSpace(result)
}

func TestScanner(t *testing.T) {
	expected := readFromFile("tests/fibonacci.scanner.out")
	actual := scan("examples/fibonacci.html")

	if expected != actual {
		t.Error(diff.Diff(expected, actual))
	}
}
