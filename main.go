package main

import (
	"io/ioutil"
	"path/filepath"

	"./logging"
)

var (
	intVal  int
	boolVal bool
	realVal float64
	strVal  string
)

var (
	scannerLog = logging.New("SCANNER")
)

func main() {
	scannerLog.SetLevel(logging.Info)
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
			scannerLog.Info("%s, %d\n", tok, intVal)
		} else if tok == tokIdentifier || tok == tokStringConst {
			scannerLog.Info("%s, %s\n", tok, strVal)
		} else {
			scannerLog.Info("%s\n", tok)
		}
	}
}
