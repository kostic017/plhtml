package parser

import (
	"../logger"
	"../scanner"
)

type Token = scanner.Token
type TokenType = scanner.TokenType

type Parser struct {
	index      int
	tokens     []Token
	logger     *logger.MyLogger
	binOpsPrec map[TokenType]int
}

func NewParser() *Parser {
	parser := new(Parser)

	parser.logger = logger.New("PARSER")
	parser.logger.SetLevel(logger.Info)

	parser.setBinOpsPrec([][]TokenType{
		[]TokenType{
			scanner.TokEqOp,
			scanner.TokNeqOp,
		},
		[]TokenType{
			scanner.TokLtOp,
			scanner.TokGtOp,
			scanner.TokLeqOp,
			scanner.TokGeqOp,
		},
		[]TokenType{
			TokenType('+'),
			TokenType('-'),
		},
		[]TokenType{
			TokenType('*'),
			TokenType('/'),
		},
	})

	return parser
}

func (parser *Parser) Parse(tokens []Token) ProgramNode {
	parser.index = -1
	parser.tokens = tokens
	return parser.parseProgram()
}

func (parser *Parser) SetLogLevel(level logger.LogLevel) {
	parser.logger.SetLevel(level)
}

func (parser *Parser) setBinOpsPrec(operators [][]TokenType) {
	parser.binOpsPrec = make(map[TokenType]int)
	for i, group := range operators {
		for _, op := range group {
			parser.binOpsPrec[op] = i
		}
	}
}

func (parser *Parser) parseOpenTag(expected TokenType) {
	parser.logger.Debug("'%s' begin tag expected", string(expected))
	parser.expect(TokenType('<'))
	parser.expect(expected)
	parser.expect(TokenType('>'))
}

func (parser *Parser) parseCloseTag(expected TokenType) {
	parser.logger.Debug("'%s' close tag expected", string(expected))
	parser.expect(TokenType('<'))
	parser.expect(TokenType('/'))
	parser.expect(expected)
	parser.expect(TokenType('>'))
}

func (parser Parser) parseProgram() ProgramNode {
	parser.logger.Debug("=BEG= Program")
	parser.parseDoctype()
	prg := parser.parseHTML()
	parser.logger.Debug("=END= Program")
	return prg
}

func (parser *Parser) parseDoctype() {
	parser.logger.Debug("=BEG= Doctype")
	parser.expect(TokenType('<'))
	parser.expect(TokenType('!'))
	parser.expect(scanner.TokDoctype)
	parser.expect(scanner.TokHTML)
	parser.expect(TokenType('>'))
	parser.logger.Debug("=END= Doctype")
}

func (parser *Parser) parseHTML() ProgramNode {
	parser.logger.Debug("=BEG= HTML")
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
	parser.logger.Debug("=END= HTML")
	return ProgramNode{Title: programTitle, Body: programBody}
}

func (parser *Parser) parseProgramHeader() StringConstNode {
	parser.logger.Debug("=BEG= Prg Header")
	parser.parseOpenTag(scanner.TokHead)
	programTitle := parser.parseProgramTitle()
	parser.parseCloseTag(scanner.TokHead)
	parser.logger.Debug("=END= Prg Header")
	return programTitle
}

func (parser *Parser) parseProgramTitle() StringConstNode {
	parser.logger.Debug("=BEG= Prg Title")
	parser.parseOpenTag(scanner.TokTitle)
	programTitle := parser.parseStringConst()
	parser.parseCloseTag(scanner.TokTitle)
	parser.logger.Debug("=END= Prg Title")
	return programTitle
}

func (parser *Parser) parseProgramBody() ProgramBodyNode {
	parser.logger.Debug("=BEG= Prg Body")
	parser.parseOpenTag(scanner.TokBody)
	mainFunc := parser.parseMainFunc()
	parser.parseCloseTag(scanner.TokBody)
	parser.logger.Debug("=END= Prg Title")
	return ProgramBodyNode{MainFunc: mainFunc}
}

func (parser *Parser) parseMainFunc() MainFuncNode {
	parser.logger.Debug("=BEG= Main")
	parser.parseOpenTag(scanner.TokMain)
	statements := parser.parseStatements()
	parser.parseCloseTag(scanner.TokMain)
	parser.logger.Debug("=END= Main")
	return MainFuncNode{Statements: statements}
}
