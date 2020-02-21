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
	if scan.index != len(scan.source)-1 {
		scan.index--
	}
}

func (scan *scanner) nextChar() (rune, bool) {
	if scan.index == len(scan.source)-1 {
		return 0, false
	}
	ch := scan.source[scan.index]
	scan.index++
	return ch, true
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

		if ch == '<' {
			if scan.lexComment() {
				// <!--.*-->
				continue
			}
		}

		if unicode.IsDigit(ch) {
			return scan.lexNumber(ch)
		}

		if ch == '&' || unicode.IsLetter(ch) {
			// &[a-zA-Z];            operators
			// [a-zA-Z][a-zA-Z0-9]*  identifiers/keywords
			return scan.lexIdentifier(ch)
		}

		switch ch {
		case '!', '/', '=', '>', '(', ')', '-', '<':
			return token(ch)
		}

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
	if scan.lookahead("!--") {
		for {
			ch, ok := scan.nextChar()
			if !ok {
				panic("Unterminated comment.")
			}
			if ch == '-' && scan.lookahead("->") {
				return true
			}
		}
	}

	return false
}

func (scan *scanner) lookahead(expected string) bool {
	got := ""
	counter := 0

	for i := 0; i < len(expected); i++ {
		counter++
		ch, ok := scan.nextChar()
		if !ok {
			break
		}
		got += string(ch)
	}

	if got != expected {
		for i := 1; i <= counter; i++ {
			scan.goBack()
		}
		return false
	}

	return true
}

func (scan *scanner) lexNumber(ch rune) token {
	var ok bool
	number := "" + string(ch)

	for {
		ch, ok = scan.nextChar()
		if ok && unicode.IsNumber(ch) {
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
	var ok bool
	identifier := string(ch)

	start := identifier
	end := identifier

	for {
		ch, ok = scan.nextChar()

		if ok && (unicode.IsLetter(ch) || (start == "&" && ch == ';') || (start != "&" && unicode.IsNumber(ch))) {
			end = string(ch)
			identifier += end
		} else {
			scan.goBack()
			break
		}

	}

	identifier = strings.ToLower(identifier)

	if start == "&" {
		if end == ";" {
			if tok, ok := scan.operators[identifier]; ok {
				return tok
			}
			panic(fmt.Sprintf("Operator %s is not valid.", identifier))
		} else {
			identifier = identifier[1:]
		}
	}

	if tok, ok := scan.keywords[identifier]; ok {
		return tok
	}

	strVal = identifier
	return tokIdentifier
}
