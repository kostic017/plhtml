package parser

import (
	"fmt"

	"../scanner"
)

func (parser Parser) current() Token {
	return parser.tokens[parser.index]
}

func (parser Parser) peek() Token {
	current := parser.current()
	if current.Type == scanner.TokEOF {
		return current
	}
	return parser.tokens[parser.index+1]
}

func (parser *Parser) next() Token {
	current := parser.current()
	if current.Type != scanner.TokEOF {
		parser.index++
	}
	return current
}

func (parser *Parser) expect(expected ...TokenType) TokenType {
	actual := parser.next().Type
	for _, exp := range expected {
		if actual == exp {
			return actual
		}
	}
	panic(fmt.Sprintf("Expected one of %s, got %s.", expected, actual))
}
