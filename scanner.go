package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type scanner struct {
	index     int
	source    []rune
	keywords  map[string]token
	operators map[string]token
}

func (scan *scanner) init(source string) {
	scan.index = 0
	scan.source = []rune(source)
	scan.keywords = make(map[string]token)
	scan.operators = make(map[string]token)

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
	scan.keywords["integer"] = tokIntType
	scan.keywords["real"] = tokRealType
	scan.keywords["boolean"] = tokBoolType
	scan.keywords["string"] = tokStringType

	// operators
	scan.operators["&plus;"] = tokAddOp
	scan.operators["&minus;"] = tokSubOp
	scan.operators["&times;"] = tokMulOp
	scan.operators["&divide;"] = tokDivOp
}

func (scan *scanner) goBack() {
	if scan.index < len(scan.source) {
		scan.index--
		scannerLog.debug("Unread char: %s\n", strconv.Quote(string(scan.source[scan.index])))
	}
}

func (scan *scanner) nextChar() (rune, bool) {
	if scan.index != len(scan.source)-1 {
		ch := scan.source[scan.index]
		scan.index++
		scannerLog.debug("Read char: %s\n", strconv.Quote(string(ch)))
		return ch, true
	}
	return 0, false
}

func (scan *scanner) lookahead(i int) (rune, bool) {
	if scan.index+i < len(scan.source)-1 {
		return scan.source[scan.index+i], true
	}
	return 0, false
}

func (scan *scanner) nextToken() token {

	for {

		ch, ok := scan.nextChar()

		if !ok {
			break
		}

		if unicode.IsSpace(ch) {
			continue
		}

		if ch == '`' {
			return scan.lexString(ch)
		}

		// after lexString because <!-- could be part of string literal
		if ch == '<' && scan.lexComment() {
			continue
		}

		if unicode.IsDigit(ch) {
			return scan.lexNumber(ch)
		}

		if ch == '&' || unicode.IsLetter(ch) {
			return scan.lexIdentifier(ch)
		}

		switch ch {
		case '"', '!', '/', '=', '<', '>', '(', ')', '-':
			return token(ch)
		}

		panic(fmt.Sprintf("Illegal character %c.", ch))

	}

	return tokEOF
}

func (scan *scanner) lexString(ch rune) token {
	var ok bool
	str := string(ch)

	for {
		ch, ok = scan.nextChar()

		if ok {

			if ch == '`' {
				break
			}

			str += string(ch)

		} else {
			panic("Unterminated string.")
		}
	}

	strVal = str[1:]
	return tokStringConst
}

func (scan *scanner) lexComment() bool {
	// <!--.*-->

	ch1, ok1 := scan.lookahead(0)
	ch2, ok2 := scan.lookahead(1)
	ch3, ok3 := scan.lookahead(2)

	if ok1 && ok2 && ok3 && ch1 == '!' && ch2 == '-' && ch3 == '-' {
		for {
			ch, ok := scan.nextChar()
			if !ok {
				panic("End of file inside comment.")
			}
			if ch == '-' {
				ch1, ok1 := scan.lookahead(0)
				ch2, ok2 := scan.lookahead(1)
				if ok1 && ok2 && ch1 == '-' && ch2 == '>' {
					scan.nextChar()
					scan.nextChar()
					return true
				}
			}
		}
	}

	return false
}

func (scan *scanner) lexNumber(ch rune) token {
	var ok bool
	real := false
	number := "" + string(ch)

	for ch, ok = scan.nextChar(); ok; ch, ok = scan.nextChar() {

		if ch == '.' {
			real = true
		}

		if ch == '.' || unicode.IsNumber(ch) {
			number += string(ch)
		} else {
			scan.goBack()
			break
		}

	}

	if real {
		f, err := strconv.ParseFloat(number, 64)

		if err != nil {
			panic(err)
		}

		realVal = f
		return tokRealConst
	}

	i, err := strconv.Atoi(number)

	if err != nil {
		panic(err)
	}

	intVal = i
	return tokIntConst
}

func (scan *scanner) lexIdentifier(ch rune) token {
	// &[a-zA-Z];            operators
	// [a-zA-Z][a-zA-Z0-9]*  identifiers/keywords

	var ok bool

	identifier := string(ch)

	for ch, ok = scan.nextChar(); ok; ch, ok = scan.nextChar() {

		if unicode.IsLetter(ch) || (identifier[0] == '&' && ch == ';') || (identifier[0] != '&' && unicode.IsNumber(ch)) {
			identifier += string(ch)
		} else {
			scan.goBack()
			break
		}

	}

	identifier = strings.ToLower(identifier)

	if identifier[0] == '&' {
		if identifier[len(identifier)-1:] != ";" {
			panic(fmt.Sprintf("Unterminated operator %s.", identifier))
		}

		if tok, ok := scan.operators[identifier]; ok {
			return tok
		}

		panic(fmt.Sprintf("Operator %s is not valid.", identifier))
	}

	if tok, ok := scan.keywords[identifier]; ok {
		return tok
	}

	strVal = identifier
	return tokIdentifier
}
