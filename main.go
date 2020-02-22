package main

import (
	"log"
	"os"
	"path/filepath"
)

var (
	intVal  int
	boolVal bool
	realVal float64
	strVal  string
)

var (
	scannerLog = log.New(os.Stdout, "SCANNER ", 0)
)

func main() {
	scan(filepath.Join("test", "ex1.html"))
}

func scan(file string) {
	source, err := readFile(file)

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
