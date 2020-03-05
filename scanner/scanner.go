package scanner

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"../logger"
)

type Scanner struct {
	line       int
	index      int
	column     int
	tabSize    int
	prevColumn int
	source     []rune
	logger     *logger.MyLogger
	keywords   map[string]TokenType
	operators  map[string]TokenType
}

func NewScanner() *Scanner {
	scanner := new(Scanner)

	scanner.tabSize = 4

	scanner.logger = logger.New("SCANNER")
	scanner.logger.SetLevel(logger.Info)

	scanner.keywords = map[string]TokenType{
		"doctype": TokDoctype,
		"lang":    TokLang,
		"html":    TokHTML,
		"head":    TokHead,
		"title":   TokTitle,
		"body":    TokBody,
		"main":    TokMain,
		"var":     TokVar,
		"class":   TokClass,
		"output":  TokOutput,
		"input":   TokInput,
		"name":    TokName,
		"data":    TokData,
		"value":   TokValue,
		"div":     TokDiv,
		"if":      TokIf,
		"while":   TokWhile,
		"integer": TokIntType,
		"real":    TokRealType,
		"boolean": TokBoolType,
		"string":  TokStringType,
	}

	scanner.operators = map[string]TokenType{
		"&lt;":     TokLtOp,
		"&gt;":     TokGtOp,
		"&leq;":    TokLeqOp,
		"&geq;":    TokGeqOp,
		"&equals;": TokEqOp,
		"&ne;":     TokNeqOp,
	}

	return scanner
}

func (scanner *Scanner) SetTabSize(tabSize int) {
	scanner.tabSize = tabSize
}

func (scanner *Scanner) SetLogLevel(level logger.LogLevel) {
	scanner.logger.SetLevel(level)
}

func (scanner *Scanner) Scan(source string) []Token {
	scanner.line = 1
	scanner.index = 0
	scanner.column = 0
	scanner.source = []rune(source)

	var tokens []Token

	for {
		tok := scanner.nextToken()
		tokens = append(tokens, tok)
		if tok.Type == TokEOF {
			break
		}
	}

	return tokens
}

func (scanner *Scanner) nextToken() Token {

	for {

		ch, ok := scanner.nextChar()

		if !ok {
			break
		}

		if unicode.IsSpace(ch) {
			continue
		}

		if ch == '`' {
			return scanner.lexString(ch, scanner.line, scanner.column)
		}

		// after lexString because <!-- could be part of string literal
		if ch == '<' && scanner.lexComment(scanner.line, scanner.column) {
			continue
		}

		if unicode.IsDigit(ch) {
			return scanner.lexNumber(ch)
		}

		if ch == '&' || unicode.IsLetter(ch) {
			return scanner.lexWord(ch, scanner.line, scanner.column)
		}

		switch ch {
		case '+', '-', '*', '/', '(', ')', '!', '"', '=', '<', '>', '.':
			return Token{Type: TokenType(ch)}
		}

		panic(fmt.Sprintf("Illegal character %c at %d:%d.", ch, scanner.line, scanner.column))

	}

	return Token{Type: TokEOF}
}

func (scanner *Scanner) nextChar() (rune, bool) {
	if scanner.index != len(scanner.source)-1 {
		ch := scanner.source[scanner.index]
		scanner.logger.Debug("Read char: %s\n", strconv.Quote(string(ch)))
		scanner.incColLine(ch)
		scanner.index++
		return ch, true
	}

	return 0, false
}

func (scanner *Scanner) lookahead(i int) (rune, bool) {
	if scanner.index+i < len(scanner.source)-1 {
		return scanner.source[scanner.index+i], true
	}
	return 0, false
}

func (scanner *Scanner) goBack() {
	if scanner.index < len(scanner.source) {
		scanner.index--
		ch := scanner.source[scanner.index]
		scanner.logger.Debug("Unread char: %s\n", strconv.Quote(string(ch)))
		scanner.decColLine(ch)
	}
}

func (scanner *Scanner) decColLine(ch rune) {
	if ch == '\n' {
		scanner.column = scanner.prevColumn
		scanner.line--
	} else if ch == '\t' {
		scanner.column -= scanner.tabSize
	} else {
		scanner.column--
	}
}

func (scanner *Scanner) incColLine(ch rune) {
	if ch == '\n' {
		scanner.prevColumn = scanner.column
		scanner.column = 0
		scanner.line++
	} else if ch == '\t' {
		scanner.column += scanner.tabSize
	} else {
		scanner.column++
	}
}

func (scanner *Scanner) lexString(ch rune, line int, column int) Token {
	var ok bool
	str := string(ch)

	for {
		ch, ok = scanner.nextChar()

		if ok {

			if ch == '`' {
				break
			}

			str += string(ch)

		} else {
			panic(fmt.Sprintf("Unterminated string at %d:%d.", line, column))
		}
	}

	return Token{Type: TokStringConst, StrVal: str[1:]}
}

func (scanner *Scanner) lexComment(line int, column int) bool {
	// <!--.*-->

	ch1, ok1 := scanner.lookahead(0)
	ch2, ok2 := scanner.lookahead(1)
	ch3, ok3 := scanner.lookahead(2)

	if ok1 && ok2 && ok3 && ch1 == '!' && ch2 == '-' && ch3 == '-' {
		for {
			ch, ok := scanner.nextChar()
			if !ok {
				panic(fmt.Sprintf("End of file inside comment at %d:%d.", line, column))
			}
			if ch == '-' {
				ch1, ok1 := scanner.lookahead(0)
				ch2, ok2 := scanner.lookahead(1)
				if ok1 && ok2 && ch1 == '-' && ch2 == '>' {
					scanner.nextChar()
					scanner.nextChar()
					return true
				}
			}
		}
	}

	return false
}

func (scanner *Scanner) lexNumber(ch rune) Token {
	var ok bool
	real := false
	number := "" + string(ch)

	for ch, ok = scanner.nextChar(); ok; ch, ok = scanner.nextChar() {

		if ch == '.' {
			real = true
		}

		if ch == '.' || unicode.IsNumber(ch) {
			number += string(ch)
		} else {
			scanner.goBack()
			break
		}

	}

	if real {
		realVal, _ := strconv.ParseFloat(number, 64)
		return Token{Type: TokRealConst, RealVal: realVal}
	}

	intVal, _ := strconv.Atoi(number)
	return Token{Type: TokIntConst, IntVal: intVal}
}

func (scanner *Scanner) lexWord(ch rune, line int, column int) Token {
	// &[a-zA-Z];            operators
	// [a-zA-Z][a-zA-Z0-9]*  identifiers/keywords

	word := string(ch)

	for ch, ok := scanner.nextChar(); ok; ch, ok = scanner.nextChar() {

		valid := unicode.IsLetter(ch) ||
			(word[0] == '&' && ch == ';') ||
			(word[0] != '&' && unicode.IsNumber(ch))

		if valid {
			word += string(ch)
		} else {
			scanner.goBack()
			break
		}

	}

	if tok, ok := scanner.lexOperator(word, line, column); ok {
		return tok
	}

	word = strings.ToLower(word)

	if word == "true" || word == "false" {
		return Token{Type: TokBoolConst, BoolVal: word == "true"}
	}

	if tok, ok := scanner.keywords[word]; ok {
		return Token{Type: tok}
	}

	return Token{Type: TokIdentifier, StrVal: word}
}

func (scanner *Scanner) lexOperator(word string, line int, column int) (Token, bool) {

	firstChar := word[0]
	lastChar := word[len(word)-1:]

	if firstChar != '&' {
		return Token{Type: TokEOF}, false
	}

	if lastChar != ";" {
		panic(fmt.Sprintf("Unterminated operator %s at %d:%d.", word, line, column))
	}

	if tok, ok := scanner.operators[word]; ok {
		return Token{Type: tok}, true
	}

	panic(fmt.Sprintf("Operator %s is not valid at %d:%d.", word, line, column))

}
