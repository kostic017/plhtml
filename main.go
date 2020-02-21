package main

import (
	"log"
	"os"
)

var (
	intVal  int
	boolVal bool
	realVal float32
	strVal  string
)

var (
	scannerLog = log.New(os.Stdout, "SCANNER ", 0)
)

func main() {
	scan()
}

func scan() {
	source, err := readFile("test/comments.html")

	if err == nil {
		var scan scanner
		scan.init(source)
		for tok := scan.nextToken(); tok != tokEOF; tok = scan.nextToken() {
			if tok == tokIntConst {
				scannerLog.Printf("(%s, %d)\n", tok, intVal)
			} else if tok == tokIdentifier || tok == tokStringConst {
				scannerLog.Printf("(%s, %s)\n", tok, strVal)
			} else {
				scannerLog.Printf("%s\n", tok)
			}
		}
	}
}
