package main

import (
	"fmt"
)

var (
	intVal  int
	boolVal bool
	realVal float32
	strVal  string
)

func main() {
	source, err := readFile("test/ex1.html")

	if err == nil {
		var scan scanner
		scan.init(source)
		for tok := scan.nextToken(); tok != tokEOF; tok = scan.nextToken() {
			if tok == tokIntConst {
				fmt.Printf("%s %d\n", tok, intVal)
			} else if tok == tokIdentifier {
				fmt.Printf("%s %s\n", tok, strVal)
			} else {
				fmt.Printf("%s\n", tok)
			}
		}
	}
}
