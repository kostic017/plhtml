package parser

import (
	"fmt"

	"../ast"
	"../logger"
	"../scanner"
)

type Token = scanner.Token
type TokenType = scanner.TokenType

type Parser struct {
	index  int
	tokens []Token
	logger *logger.MyLogger
}

func NewParser() *Parser {
	parser := new(Parser)
	parser.logger = logger.New("PARSER")
	parser.logger.SetLevel(logger.Info)
	return parser
}

func (parser *Parser) Parse(tokens []Token) ast.ProgramNode {
	parser.index = -1
	parser.tokens = tokens
	return parser.parseProgram()
}

func (parser *Parser) SetLogLevel(level logger.LogLevel) {
	parser.logger.SetLevel(level)
}

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

func (parser *Parser) parseOpenTag(expected TokenType) {
	parser.logger.Debug("<%s> expected", string(expected))
	parser.expect(TokenType('<'))
	parser.expect(expected)
	parser.expect(TokenType('>'))
}

func (parser *Parser) parseCloseTag(expected TokenType) {
	parser.logger.Debug("</%s> expected", string(expected))
	parser.expect(TokenType('<'))
	parser.expect(TokenType('/'))
	parser.expect(expected)
	parser.expect(TokenType('>'))
}
