package main

import (
	"io/ioutil"
	"path/filepath"
)

var (
	intVal  int
	boolVal bool
	realVal float64
	strVal  string
)

var (
	scannerLog = newLogger("SCANNER")
)

func main() {
	scannerLog.setLevel(lvlInfo)
	scan(filepath.Join("test", "ex1.html"))
}

func scan(file string) {
	filebuffer, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var scan scanner
	scan.init(string(filebuffer))
	for tok := scan.nextToken(); tok != tokEOF; tok = scan.nextToken() {
		if tok == tokIntConst {
			scannerLog.info("%s, %d\n", tok, intVal)
		} else if tok == tokIdentifier || tok == tokStringConst {
			scannerLog.info("%s, %s\n", tok, strVal)
		} else {
			scannerLog.info("%s\n", tok)
		}
	}
}
