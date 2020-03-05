package parser

import (
	"../logger"
	"../scanner"
)

type Token = scanner.Token
type TokenType = scanner.TokenType

type Parser struct {
	index  int
	tokens []Token
	logger *logger.MyLogger
	opPrec map[TokenType]int
}

func NewParser() *Parser {
	parser := new(Parser)

	parser.logger = logger.New("PARSER")
	parser.logger.SetLevel(logger.Info)

	parser.setOpPrec([]TokenType{
		scanner.TokLtOp,
		scanner.TokGtOp,
		scanner.TokLeqOp,
		scanner.TokGeqOp,
		scanner.TokEqOp,
		scanner.TokNeqOp,
		TokenType('+'),
		TokenType('-'),
		TokenType('*'),
		TokenType('/'),
	})

	return parser
}

func (parser *Parser) SetLogLevel(level logger.LogLevel) {
	parser.logger.SetLevel(level)
}

func (parser *Parser) Parse(tokens []Token) {
	parser.index = 0
	parser.tokens = tokens
	parser.parseProgram()
}

func (parser *Parser) setOpPrec(operators []TokenType) {
	parser.opPrec = make(map[TokenType]int)
	for i, v := range operators {
		parser.opPrec[v] = i
	}
}

// func (parser *Parser) getOpPrec(operator TokenType) int {

// }

func (parser *Parser) parseOpenTag(expected TokenType) {
	parser.expect(TokenType('<'))
	parser.expect(expected)
	parser.expect(TokenType('>'))
}

func (parser *Parser) parseCloseTag(expected TokenType) {
	parser.expect(TokenType('<'))
	parser.expect(TokenType('/'))
	parser.expect(expected)
	parser.expect(TokenType('>'))
}

func (parser Parser) parseProgram() ProgramNode {
	parser.parseDoctype()
	return parser.parseHTML()
}

func (parser *Parser) parseDoctype() {
	parser.expect(TokenType('<'))
	parser.expect(TokenType('!'))
	parser.expect(scanner.TokDoctype)
	parser.expect(scanner.TokHTML)
	parser.expect(TokenType('>'))
}

func (parser *Parser) parseHTML() ProgramNode {
	parser.expect(TokenType('<'))
	parser.expect(scanner.TokHTML)
	parser.expect(scanner.TokLang)
	parser.expect(TokenType('='))
	parser.expect(TokenType('"'))
	parser.parseIdentifier()
	parser.expect(TokenType('"'))
	parser.expect(TokenType('>'))
	programTitle := parser.parseProgramHeader()
	programBody := parser.parseProgramBody()
	parser.parseCloseTag(scanner.TokHTML)
	return ProgramNode{Title: programTitle, Body: programBody}
}

func (parser *Parser) parseProgramHeader() StringConstNode {
	parser.parseOpenTag(scanner.TokHead)
	programTitle := parser.parseProgramTitle()
	parser.parseCloseTag(scanner.TokHead)
	return programTitle
}

func (parser *Parser) parseProgramTitle() StringConstNode {
	parser.parseOpenTag(scanner.TokTitle)
	programTitle := parser.parseStringConst()
	parser.parseCloseTag(scanner.TokTitle)
	return programTitle
}

func (parser *Parser) parseProgramBody() ProgramBodyNode {
	parser.parseOpenTag(scanner.TokBody)
	mainFunc := parser.parseMainFunc()
	parser.parseCloseTag(scanner.TokBody)
	return ProgramBodyNode{MainFunc: mainFunc}
}

func (parser *Parser) parseMainFunc() MainFuncNode {
	parser.parseOpenTag(scanner.TokMain)
	statements := parser.parseStatements()
	parser.parseCloseTag(scanner.TokMain)
	return MainFuncNode{Statements: statements}
}
