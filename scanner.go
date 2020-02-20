package main

import (
	"strconv"
	"unicode"
)

type scanner struct {
	index  int
	source []rune
}

func (scan *scanner) init(source string) {
	scan.index = 0
	scan.source = []rune(source)
}

func (scan *scanner) back() {
	if scan.index != len(scan.source)-1 {
		scan.index--
	}
}

func (scan *scanner) nextChar() rune {
	if scan.index == len(scan.source)-1 {
		return 0
	}
	ch := scan.source[scan.index]
	scan.index++
	return ch
}

func (scan *scanner) nextToken() token {
	for ch := scan.nextChar(); ch != 0; ch = scan.nextChar() {

		if unicode.IsSpace(ch) {
			continue
		}

		if unicode.IsDigit(ch) {
			return scan.lexNumber(ch)
		}

		if unicode.IsLetter(ch) {
			// handle identifiers/keywords
		}

		switch ch {
		// case '"': // string
		// case '<': // comment?
		// case '&': // operators
		case '!', '/', '=', '>', '(', ')':
			return token(ch)
		}

	}
	return tokEOF
}

func (scan *scanner) lexNumber(ch rune) token {
	number := "" + string(ch)

	for {
		ch = scan.nextChar()
		if unicode.IsNumber(ch) {
			number += string(ch)
		} else {
			scan.back()
			break
		}
	}

	i, err := strconv.Atoi(number)

	if err != nil {
		panic(err)
	}

	intVal = i
	return tokIntConst
}
