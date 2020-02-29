package main

import "./utility"

func main() {
	scan := NewScanner(utility.ReadFile("tests/examples/fibonacci.html"))
	parser := NewParser(scan)
	parser.Parse()
}
