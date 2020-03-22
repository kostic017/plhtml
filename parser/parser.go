package parser

import (
    "fmt"

    "../ast"
    "../logger"
    "../scanner"
)

type Token = scanner.Token
type TokenType = scanner.TokenType

var myLogger = logger.New("PARSER")

func SetLogLevel(level logger.LogLevel) {
    myLogger.SetLevel(level)
}

type Parser struct {
    index  int
    tokens []Token
}

func New() *Parser {
    parser := new(Parser)
    return parser
}

func (parser *Parser) Parse(tokens []Token) ast.ProgramNode {
    parser.index = -1
    parser.tokens = tokens
    return parser.parseProgram()
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
    myLogger.Debug("'%s'", string(actual))

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
    myLogger.Debug("<%s> expected", string(expected))
    parser.expect(TokenType('<'))
    parser.expect(expected)
    parser.expect(TokenType('>'))
}

func (parser *Parser) parseCloseTag(expected TokenType) {
    myLogger.Debug("</%s> expected", string(expected))
    parser.expect(TokenType('<'))
    parser.expect(TokenType('/'))
    parser.expect(expected)
    parser.expect(TokenType('>'))
}
