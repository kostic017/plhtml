package main

import (
	"strconv"
	"strings"
	"unicode"
)

type scanner struct {
	index    int
	source   []rune
	keywords map[string]token
}

func (scan *scanner) init(source string) {
	scan.index = 0
	scan.source = []rune(source)

	scan.keywords = make(map[string]token)
	scan.keywords["doctype"] = tokDoctype
	scan.keywords["lang"] = tokLang
	scan.keywords["html"] = tokHTML
	scan.keywords["head"] = tokHead
	scan.keywords["title"] = tokTitle
	scan.keywords["body"] = tokBody
	scan.keywords["main"] = tokMain
	scan.keywords["var"] = tokVar
	scan.keywords["class"] = tokClass
	scan.keywords["output"] = tokOutput
	scan.keywords["input"] = tokInput
	scan.keywords["name"] = tokName
	scan.keywords["data"] = tokData
	scan.keywords["value"] = tokValue
	scan.keywords["if"] = tokIf
	scan.keywords["while"] = tokWhile
	scan.keywords["for"] = tokFor
}

func (scan *scanner) goBack() {
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
			return scan.lexIdentifier(ch)
		}

		switch ch {
		// case '"': // string
		case '<':
			return token('<') // TODO comments
		// case '&': // operators
		case '!', '/', '=', '>', '(', ')', '-':
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
			scan.goBack()
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

func (scan *scanner) lexIdentifier(ch rune) token {
	// [a-zA-Z][a-zA-Z0-9]*

	identifier := string(ch)

	for {
		ch = scan.nextChar()
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			identifier += string(ch)
		} else {
			scan.goBack()
			break
		}
	}

	if tok, ok := scan.keywords[strings.ToLower(identifier)]; ok {
		return tok
	}

	strVal = identifier
	return tokIdentifier
}
