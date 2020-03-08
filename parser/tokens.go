package parser

import (
	"fmt"

	"../scanner"
)

func (parser Parser) current() Token {
	return parser.tokens[parser.index]
}

func (parser Parser) peek() Token {
	if parser.current().Type == scanner.TokEOF {
		return parser.current()
	}
	return parser.tokens[parser.index+1]
}

func (parser *Parser) goBack() {
	parser.index--
}

func (parser *Parser) next() Token {
	if parser.index < 0 || parser.current().Type != scanner.TokEOF {
		parser.index++
	}
	return parser.current()
}

func (parser *Parser) expect(expected ...TokenType) TokenType {
	actual := parser.next().Type
	parser.logger.Debug("got '%s'", string(actual))

	for _, exp := range expected {
		if actual == exp {
			return actual
		}
	}

	panic(fmt.Sprintf("Unexpected token %s.", string(actual)))
}
