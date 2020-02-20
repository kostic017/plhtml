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
				fmt.Printf("%d\t\t%s\n", intVal, tok)
			} else if tok == tokIdentifier || tok == tokStringConst {
				fmt.Printf("%s\t\t%s\n", strVal, tok)
			} else {
				fmt.Printf("%s\n", tok)
			}
		}
	}
}
