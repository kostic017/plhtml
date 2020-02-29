package main

import (
	"fmt"

	"./logging"
)

/*

prog = '<' '!' TokDoctype TokHtml '>' '<' TokHTML TokLang '=' '"' TokIdentifier '"' '>' prog_header prog_body '<' '/' TokHTML '>';
prog_header = '<' TokHead '>' prog_title '<' '/' TokHEAD '>';
prog_title  = '<' TokTitle '>' TokStringConst '<' '/' TokTitle '>';
prog_body   = '<' TokBody '>' func_main '<' '/' TokBody '>';
func_main   = '<' TokMain '>' stmts '<' '/' TokMain '>';

stmt = '<' TokVar TokClass '=' '"' TokIntType '"' '>' TokIdentifier '<' '/' TokVar '>'
     | '<' TokVar TokClass '=' '"' TokRealType '"' '>' TokIdentifier '<' '/' TokVar '>'
     | '<' TokVar TokClass '=' '"' TokBoolType '"' '>' TokIdentifier '<' '/' TokVar '>'
     | '<' TokVar TokClass '=' '"' TokStringType '"' '>' TokIdentifier '<' '/' TokVar '>'
     | '<' TokData TokValue '=' expr '>' TokIdentifier '<' '/' TokData '>'
     | '<' TokOutput '>' expr '<' '/' TokOutput '>'
     | '<' TokInput TokName '=' '"' TokIdentifier '"' '>'
     | '<' TokDiv TokData '-' TokWhile '=' '"' expr '"' '>' ... '<' '/' TokDiv '>'
     ;

stmts = stmt
      | stmt stmts
      ;

expr = TokIntConst
     | TokRealConst
     | TokBoolConst
     | TokStringConst
     | TokIdentifier
     | expr '+' expr
     | expr '-' expr
     | expr '*' expr
     | expr '/' expr
     | expr TokLtOp expr
     | expr TokGtOp expr
     | expr TokLeqOp expr
     | expr TokGeqOp expr
     | expr TokEqOp expr
     | expr TokNeqOp expr
     | "(" expr ")"
     ;
*/

type Parser struct {
	scanner *Scanner
	curTok  Token
	logger  *logging.MyLogger
}

func NewParser(scanner *Scanner) *Parser {
	parser := new(Parser)
	parser.logger = logging.New("PARSER")
	parser.logger.SetLevel(logging.Info)
	parser.scanner = scanner
	return parser
}

func (parser Parser) Parse() {
	parser.parseProgram()
}

func (parser Parser) parseProgram() ProgramNode {
	parser.parseDoctype()
	return parser.parseHTML()
}

func (parser Parser) parseDoctype() {
	if parser.nextToken().Type != TokenType('<') {
		panic("'<' expected.")
	}
	if parser.nextToken().Type != TokenType('!') {
		panic("'!' expected.")
	}
	if parser.nextToken().Type != TokDoctype {
		panic("'DOCTYPE' expected.")
	}
	if parser.nextToken().Type != TokHTML {
		panic("'html' expected.")
	}
	if parser.nextToken().Type != TokenType('>') {
		panic("'>' expected.")
	}
}

func (parser Parser) parseHTML() ProgramNode {
	if parser.nextToken().Type != TokenType('<') {
		panic("'<' expected.")
	}
	if parser.nextToken().Type != TokHTML {
		panic("'html' expected.")
	}
	if parser.nextToken().Type != TokLang {
		panic("'lang' expected.")
	}
	if parser.nextToken().Type != TokenType('=') {
		panic("'=' expected.")
	}
	if parser.nextToken().Type != TokenType('"') {
		panic("'\"' expected.")
	}

	parser.parseIdentifier()

	if parser.nextToken().Type != TokenType('"') {
		panic("'\"' expected.")
	}
	if parser.nextToken().Type != TokenType('>') {
		panic("'>' expected.")
	}

	programTitle := parser.parseProgramHeader()
	programBody := parser.parseProgramBody()

	parser.parseCloseTag(TokHTML)

	return ProgramNode{Title: programTitle, Body: programBody}
}

func (parser *Parser) parseProgramHeader() ProgramTitleNode {
	parser.parseOpenTag(TokHead)
	programTitle := parser.parseProgramTitle()
	parser.parseCloseTag(TokHead)
	return programTitle
}

func (parser *Parser) parseProgramTitle() ProgramTitleNode {
	parser.parseOpenTag(TokTitle)
	value := parser.parseStringConst()
	parser.parseCloseTag(TokTitle)
	return ProgramTitleNode{Value: value}
}

func (parser *Parser) parseProgramBody() ProgramBodyNode {
	parser.parseOpenTag(TokBody)
	mainFunc := parser.parseMainFunc()
	parser.parseCloseTag(TokBody)
	return ProgramBodyNode{MainFunc: mainFunc}
}

func (parser *Parser) parseMainFunc() MainFuncNode {
	parser.parseOpenTag(TokMain)
	statements := parser.parseStatements()
	parser.parseCloseTag(TokMain)
	return MainFuncNode{Statements: statements}
}

func (parser *Parser) parseStatements() []StatementNode {
	return nil // TODO
}

func (parser *Parser) parseIdentifier() IdentifierNode {
	parser.nextToken()
	if parser.curTok.Type != TokIdentifier {
		panic("Expected identifier.")
	}
	return IdentifierNode{Name: parser.curTok.StrVal}
}

func (parser *Parser) parseStringConst() StringConstNode {
	parser.nextToken()
	if parser.curTok.Type != TokStringConst {
		panic("Expected string constant.")
	}
	return StringConstNode{Value: parser.curTok.StrVal}
}

func (parser *Parser) nextToken() Token {
	parser.curTok = parser.scanner.NextToken()
	return parser.curTok
}

func (parser *Parser) parseOpenTag(expected TokenType) {
	if parser.nextToken().Type != TokenType('<') {
		panic("'<' expected.")
	}
	if parser.nextToken().Type != expected {
		panic(fmt.Sprintf("'%s' expected.", expected))
	}
	if parser.nextToken().Type != TokenType('>') {
		panic("'>' expected.")
	}
}

func (parser *Parser) parseCloseTag(expected TokenType) {
	if parser.nextToken().Type != TokenType('<') {
		panic("'<' expected.")
	}
	if parser.nextToken().Type != TokenType('/') {
		panic("'/' expected.")
	}
	if parser.nextToken().Type != expected {
		panic(fmt.Sprintf("'%s' expected.", expected))
	}
	if parser.nextToken().Type != TokenType('>') {
		panic("'>' expected.")
	}
}

/*
func (parser *Parser) parsePrimary() ExpressionNode {
	switch parser.curTok.Type {
	case TokIntConst:
		return parser.parseIntConst()
	case TokRealConst:
		return parser.parseRealConst()
	case TokBoolConst:
		return parser.parseBoolConst()
	case TokStringConst:
		return parser.parseStringConst()
	}
	panic("Unknown token.")
}

// IntConstNode -> TokIntConst
func (parser *Parser) parseIntConst() IntConstNode {
	node := IntConstNode{Value: parser.curTok.IntVal}
	parser.nextToken()
	return node
}

// RealConst -> TokRealConst
func (parser *Parser) parseRealConst() RealConstNode {
	node := RealConstNode{Value: parser.curTok.RealVal}
	parser.nextToken()
	return node
}

// BoolConstNode -> TokBoolConst
func (parser *Parser) parseBoolConst() BoolConstNode {
	node := BoolConstNode{Value: parser.curTok.BoolVal}
	parser.nextToken()
	return node
}

// StringConstNode -> TokStringConst
func (parser *Parser) parseStringConst() StringConstNode {
	node := StringConstNode{Value: parser.curTok.StrVal}
	parser.nextToken()
	return node
}

// Identifier -> TokIdentifier

*/
