package scanner

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"../logging"
)

type Scanner struct {
	line      int
	index     int
	column    int
	source    []rune
	keywords  map[string]TokenType
	operators map[string]TokenType
}

var (
	logger = logging.New("SCANNER")
)

func (scan *Scanner) init(source string) {
	logger.SetLevel(logging.Info)

	scan.line = 0
	scan.index = 0
	scan.column = 0
	scan.source = []rune(source)

	scan.keywords = map[string]TokenType{
		"doctype": Doctype,
		"lang":    Lang,
		"html":    HTML,
		"head":    Head,
		"title":   Title,
		"body":    Body,
		"main":    Main,
		"var":     Var,
		"class":   Class,
		"output":  Output,
		"input":   Input,
		"name":    Name,
		"data":    Data,
		"value":   Value,
		"div":     Div,
		"if":      If,
		"while":   While,
		"integer": IntType,
		"real":    RealType,
		"boolean": BoolType,
		"string":  StringType,
	}

	scan.operators = map[string]TokenType{
		"&plus;":   AddOp,
		"&minus;":  SubOp,
		"&times;":  MulOp,
		"&divide;": DivOp,
		"&lt;":     LtOp,
		"&gt;":     GtOp,
		"&leq;":    LeqOp,
		"&geq;":    GeqOp,
		"&Equal;":  EqOp,
		"&ne;":     NeqOp,
		"&Not;":    NotOp,
	}
}

func (scan *Scanner) goBack() {
	if scan.index < len(scan.source) {
		scan.index--
		logger.Debug("Unread char: %s\n", strconv.Quote(string(scan.source[scan.index])))
	}
}

func (scan *Scanner) nextChar() (rune, bool) {
	if scan.index != len(scan.source)-1 {
		ch := scan.source[scan.index]
		scan.index++
		logger.Debug("Read char: %s\n", strconv.Quote(string(ch)))
		return ch, true
	}
	return 0, false
}

func (scan *Scanner) lookahead(i int) (rune, bool) {
	if scan.index+i < len(scan.source)-1 {
		return scan.source[scan.index+i], true
	}
	return 0, false
}

func (scan *Scanner) nextToken() Token {

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
			return Token{Type: TokenType(ch)}
		}

		panic(fmt.Sprintf("Illegal character %c.", ch))

	}

	return Token{Type: EOF}
}

func (scan *Scanner) lexString(ch rune) Token {
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

	return Token{Type: StringConst, Value: str[1:]}
}

func (scan *Scanner) lexComment() bool {
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

func (scan *Scanner) lexNumber(ch rune) Token {
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
		return Token{Type: RealConst, Value: number}
	}

	return Token{Type: IntConst, Value: number}
}

func (scan *Scanner) lexWord(ch rune) Token {
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
		return Token{Type: tok}
	}

	return Token{Type: Identifier, Value: word}
}

func (scan *Scanner) lexOperator(word string) (Token, bool) {

	firstChar := word[0]
	lastChar := word[len(word)-1:]

	if firstChar != '&' {
		return Token{Type: EOF}, false
	}

	if lastChar != ";" {
		panic(fmt.Sprintf("Unterminated operator %s.", word))
	}

	if tok, ok := scan.operators[word]; ok {
		return Token{Type: tok}, true
	}

	panic(fmt.Sprintf("Operator %s is not valid.", word))

}
