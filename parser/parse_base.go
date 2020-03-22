package parser

import (
    "../ast"
    "../scanner"
)

func (parser Parser) parseProgram() ast.ProgramNode {
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

func (parser *Parser) parseHTML() ast.ProgramNode {
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
    return ast.ProgramNode{Title: programTitle, Body: programBody}
}

func (parser *Parser) parseProgramHeader() ast.StringConstNode {
    parser.logger.Debug("=BEG= Prg Header")
    parser.parseOpenTag(scanner.TokHead)
    programTitle := parser.parseProgramTitle()
    parser.parseCloseTag(scanner.TokHead)
    parser.logger.Debug("=END= Prg Header")
    return programTitle
}

func (parser *Parser) parseProgramTitle() ast.StringConstNode {
    parser.logger.Debug("=BEG= Prg Title")
    parser.parseOpenTag(scanner.TokTitle)
    programTitle := parser.parseStringConst()
    parser.parseCloseTag(scanner.TokTitle)
    parser.logger.Debug("=END= Prg Title")
    return programTitle
}

func (parser *Parser) parseProgramBody() ast.ProgramBodyNode {
    parser.logger.Debug("=BEG= Prg Body")
    parser.parseOpenTag(scanner.TokBody)
    mainFunc := parser.parseMainFunc()
    parser.parseCloseTag(scanner.TokBody)
    parser.logger.Debug("=END= Prg Title")
    return ast.ProgramBodyNode{MainFunc: mainFunc}
}

func (parser *Parser) parseMainFunc() ast.MainFuncNode {
    parser.logger.Debug("=BEG= Main")
    parser.parseOpenTag(scanner.TokMain)
    statements := parser.parseStatements()
    parser.parseCloseTag(scanner.TokMain)
    parser.logger.Debug("=END= Main")
    return ast.MainFuncNode{Statements: statements}
}
