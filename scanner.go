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

	scan.keywords = map[string]token{
		"doctype": tokDoctype,
		"lang":    tokLang,
		"html":    tokHTML,
		"head":    tokHead,
		"title":   tokTitle,
		"body":    tokBody,
		"main":    tokMain,
		"var":     tokVar,
		"class":   tokClass,
		"output":  tokOutput,
		"input":   tokInput,
		"name":    tokName,
		"data":    tokData,
		"value":   tokValue,
		"div":     tokDiv,
		"if":      tokIf,
		"while":   tokWhile,
		"integer": tokIntType,
		"real":    tokRealType,
		"boolean": tokBoolType,
		"string":  tokStringType,
	}

	scan.operators = map[string]token{
		"&plus;":   tokAddOp,
		"&minus;":  tokSubOp,
		"&times;":  tokMulOp,
		"&divide;": tokDivOp,
		"&lt;":     tokLtOp,
		"&gt;":     tokGtOp,
		"&leq;":    tokLeqOp,
		"&geq;":    tokGeqOp,
		"&Equal;":  tokEqOp,
		"&ne;":     tokNeqOp,
		"&Not;":    tokNotOp,
	}
}

func (scan *scanner) goBack() {
	if scan.index < len(scan.source) {
		scan.index--
		scannerLog.Debug("Unread char: %s\n", strconv.Quote(string(scan.source[scan.index])))
	}
}

func (scan *scanner) nextChar() (rune, bool) {
	if scan.index != len(scan.source)-1 {
		ch := scan.source[scan.index]
		scan.index++
		scannerLog.Debug("Read char: %s\n", strconv.Quote(string(ch)))
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
			return scan.lexWord(ch)
		}

		switch ch {
		case '"', '!', '/', '=', '<', '>', '(', ')', '-', '.':
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
		check(err)

		realVal = f
		return tokRealConst
	}

	i, err := strconv.Atoi(number)
	check(err)

	intVal = i
	return tokIntConst
}

func (scan *scanner) lexWord(ch rune) token {
	// &[a-zA-Z];            operators
	// [a-zA-Z][a-zA-Z0-9]*  identifiers/keywords

	word := string(ch)

	for ch, ok := scan.nextChar(); ok; ch, ok = scan.nextChar() {

		valid := unicode.IsLetter(ch) ||
			(word[0] == '&' && ch == ';') ||
			(word[0] != '&' && unicode.IsNumber(ch))

		if valid {
			word += string(ch)
		} else {
			scan.goBack()
			break
		}

	}

	if tok, ok := scan.lexOperator(word); ok {
		return tok
	}

	if tok, ok := scan.keywords[strings.ToLower(word)]; ok {
		return tok
	}

	strVal = word
	return tokIdentifier
}

func (scan *scanner) lexOperator(word string) (token, bool) {

	firstChar := word[0]
	lastChar := word[len(word)-1:]

	if firstChar != '&' {
		return tokEOF, false
	}

	if lastChar != ";" {
		panic(fmt.Sprintf("Unterminated operator %s.", word))
	}

	if tok, ok := scan.operators[word]; ok {
		return tok, ok
	}

	panic(fmt.Sprintf("Operator %s is not valid.", word))

}
