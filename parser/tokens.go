package parser

import (
	"fmt"

	"../scanner"
)

func (parser Parser) current() Token {
	return parser.tokens[parser.index]
}

func (parser *Parser) next() Token {
	if parser.index < 0 || parser.current().Type != scanner.TokEOF {
		parser.index++
	}
	return parser.current()
}

func (parser *Parser) goBack() {
	parser.index--
}

func (parser Parser) peek() Token {
	next := parser.next()
	parser.goBack()
	return next
}

func (parser *Parser) expectOpt(expected ...TokenType) bool {
	actual := parser.peek().Type
	parser.logger.Debug("'%s'", string(actual))

	for _, exp := range expected {
		if actual == exp {
			return true
		}
	}

	return false
}

func (parser *Parser) expect(expected ...TokenType) TokenType {
	ok := parser.expectOpt(expected...)
	if !ok {
		panic(fmt.Sprintf("Unexpected token '%s'.", string(parser.peek().Type)))
	}
	return parser.next().Type
}
