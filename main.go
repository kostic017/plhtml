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
	source, err := readFile("test/test1.html")

	if err == nil {
		var scan scanner
		scan.init(source)
		for tok := scan.nextToken(); tok != 0; tok = scan.nextToken() {
			if tok == tokIntConst {
				fmt.Printf("%s %d\n", tok, intVal)
			} else {
				fmt.Printf("%s\n", tok)
			}
		}
	}
}
