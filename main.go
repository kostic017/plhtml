package main

import (
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
}
