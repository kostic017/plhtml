package parser

import (
	"../ast"
	"../scanner"
)

func (parser Parser) parseProgram() ast.ProgramNode {
	myLogger.Debug("=BEG= Program")
	parser.parseDoctype()
	prg := parser.parseHTML()
	myLogger.Debug("=END= Program")
	return prg
}

func (parser *Parser) parseDoctype() {
	myLogger.Debug("=BEG= Doctype")
	parser.eat(TokenType('<'))
	parser.eat(TokenType('!'))
	parser.eat(scanner.TokDoctype)
	parser.eat(scanner.TokHTML)
	parser.eat(TokenType('>'))
	myLogger.Debug("=END= Doctype")
}

func (parser *Parser) parseHTML() ast.ProgramNode {
	myLogger.Debug("=BEG= HTML")
	parser.eat(TokenType('<'))
	parser.eat(scanner.TokHTML)
	parser.eat(scanner.TokLang)
	parser.eat(TokenType('='))
	parser.eat(TokenType('"'))
	parser.parseIdentifier()
	parser.eat(TokenType('"'))
	parser.eat(TokenType('>'))
	programTitle := parser.parseProgramHeader()
	programBody := parser.parseProgramBody()
	parser.parseCloseTag(scanner.TokHTML)
	myLogger.Debug("=END= HTML")
	return ast.ProgramNode{Title: programTitle, Body: programBody}
}

func (parser *Parser) parseProgramHeader() ast.StringConstNode {
	myLogger.Debug("=BEG= Prg Header")
	parser.parseOpenTag(scanner.TokHead)
	programTitle := parser.parseProgramTitle()
	parser.parseCloseTag(scanner.TokHead)
	myLogger.Debug("=END= Prg Header")
	return programTitle
}

func (parser *Parser) parseProgramTitle() ast.StringConstNode {
	myLogger.Debug("=BEG= Prg Title")
	parser.parseOpenTag(scanner.TokTitle)
	programTitle := parser.parseStringConst()
	parser.parseCloseTag(scanner.TokTitle)
	myLogger.Debug("=END= Prg Title")
	return programTitle
}

func (parser *Parser) parseProgramBody() ast.ProgramBodyNode {
	myLogger.Debug("=BEG= Prg Body")
	parser.parseOpenTag(scanner.TokBody)
	mainFunc := parser.parseMainFunc()
	parser.parseCloseTag(scanner.TokBody)
	myLogger.Debug("=END= Prg Title")
	return ast.ProgramBodyNode{MainFunc: mainFunc}
}

func (parser *Parser) parseMainFunc() ast.MainFuncNode {
	myLogger.Debug("=BEG= Main")
	parser.parseOpenTag(scanner.TokMain)
	statements := parser.parseStatements()
	parser.parseCloseTag(scanner.TokMain)
	myLogger.Debug("=END= Main")
	return ast.MainFuncNode{Statements: statements}
}
