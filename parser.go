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
     | '<' TokData TokValue '=' '"' expr '"' '>' TokIdentifier '<' '/' TokData '>'
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

func (parser Parser) nextTokenCheck(expected TokenType) TokenType {
	actual := parser.nextToken().Type
	if actual != expected {
		panic(fmt.Sprintf("'%s' expected.", expected))
	}
	return actual
}

func (parser Parser) nextTokenCheckMore(expected ...TokenType) TokenType {
	actual := parser.nextToken().Type
	for _, exp := range expected {
		if actual == exp {
			return actual
		}
	}
	panic(fmt.Sprintf("Expected one of %s, got %s.", expected, actual))
}

func (parser *Parser) nextToken() Token {
	parser.curTok = parser.scanner.NextToken()
	return parser.curTok
}

func (parser *Parser) parseOpenTag(expected TokenType) {
	parser.nextTokenCheck(TokenType('<'))
	parser.nextTokenCheck(expected)
	parser.nextTokenCheck(TokenType('>'))
}

func (parser *Parser) parseCloseTag(expected TokenType) {
	parser.nextTokenCheck(TokenType('<'))
	parser.nextTokenCheck(TokenType('/'))
	parser.nextTokenCheck(expected)
	parser.nextTokenCheck(TokenType('>'))
}

func (parser Parser) parseProgram() ProgramNode {
	parser.parseDoctype()
	return parser.parseHTML()
}

func (parser Parser) parseDoctype() {
	parser.nextTokenCheck(TokenType('<'))
	parser.nextTokenCheck(TokenType('!'))
	parser.nextTokenCheck(TokDoctype)
	parser.nextTokenCheck(TokHTML)
	parser.nextTokenCheck(TokenType('>'))
}

func (parser Parser) parseHTML() ProgramNode {
	parser.nextTokenCheck(TokenType('<'))
	parser.nextTokenCheck(TokHTML)
	parser.nextTokenCheck(TokLang)
	parser.nextTokenCheck(TokenType('='))
	parser.nextTokenCheck(TokenType('"'))
	parser.parseIdentifier()
	parser.nextTokenCheck(TokenType('"'))
	parser.nextTokenCheck(TokenType('>'))
	programTitle := parser.parseProgramHeader()
	programBody := parser.parseProgramBody()
	parser.parseCloseTag(TokHTML)
	return ProgramNode{Title: programTitle, Body: programBody}
}

func (parser *Parser) parseProgramHeader() StringConstNode {
	parser.parseOpenTag(TokHead)
	programTitle := parser.parseProgramTitle()
	parser.parseCloseTag(TokHead)
	return programTitle
}

func (parser *Parser) parseProgramTitle() StringConstNode {
	parser.parseOpenTag(TokTitle)
	programTitle := parser.parseStringConst()
	parser.parseCloseTag(TokTitle)
	return programTitle
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
	parser.nextTokenCheck(TokenType('<'))

	switch parser.nextToken().Type {
	case TokVar:
		parser.parseVarDecl()
	case TokData:
		parser.parseVarAssign()
	case TokOutput:
		parser.parseWriteStmt()
	case TokInput:
		parser.parseReadStmt()
	case TokDiv:
		parser.parseControlFlowStmt()
	}

	return nil // TODO
}

func (parser *Parser) parseVarDecl() VarDeclNode {
	parser.nextTokenCheck(TokClass)
	parser.nextTokenCheck(TokenType('='))
	parser.nextTokenCheck(TokenType('"'))
	varType := parser.nextTokenCheckMore(TokIntType, TokRealType, TokBoolType, TokStringType)
	parser.nextTokenCheck(TokenType('"'))
	parser.nextTokenCheck(TokenType('>'))
	identifier := parser.parseIdentifier()
	parser.parseCloseTag(TokVar)
	return VarDeclNode{Type: varType, Identifier: identifier}
}

func (parser *Parser) parseVarAssign() VarAssignNode {
	parser.nextTokenCheck(TokValue)
	parser.nextTokenCheck(TokenType('='))
	parser.nextTokenCheck(TokenType('"'))
	value := parser.parseExpression()
	parser.nextTokenCheck(TokenType('"'))
	parser.nextTokenCheck(TokenType('>'))
	identifier := parser.parseIdentifier()
	parser.parseCloseTag(TokData)
	return VarAssignNode{Identifier: identifier, Value: value}
}

func (parser *Parser) parseWriteStmt() WriteStmtNode {
	parser.nextTokenCheck(TokenType('>'))
	value := parser.parseExpression()
	parser.parseCloseTag(TokOutput)
	return WriteStmtNode{Value: value}
}

func (parser *Parser) parseReadStmt() ReadStmtNode {
	parser.nextTokenCheck(TokName)
	parser.nextTokenCheck(TokenType('='))
	parser.nextTokenCheck(TokenType('"'))
	identifier := parser.parseIdentifier()
	parser.nextTokenCheck(TokenType('"'))
	parser.nextTokenCheck(TokenType('>'))
	return ReadStmtNode{Identifier: identifier}
}

func (parser *Parser) parseControlFlowStmt() StatementNode {
	// TODO if, if-else
	parser.nextTokenCheck(TokData)
	parser.nextTokenCheck(TokenType('-'))
	parser.nextTokenCheck(TokWhile)
	parser.nextTokenCheck(TokenType('='))
	parser.nextTokenCheck(TokenType('"'))
	condition := parser.parseExpression()
	parser.nextTokenCheck(TokenType('"'))
	parser.nextTokenCheck(TokenType('>'))
	statements := parser.parseStatements()
	parser.parseCloseTag(TokDiv)
	return WhileStmtNode{Condition: condition, Statements: statements}
}

func (parser *Parser) parseIdentifier() IdentifierNode {
	parser.nextTokenCheck(TokIdentifier)
	return IdentifierNode{Name: parser.curTok.StrVal}
}

func (parser *Parser) parseExpression() ExpressionNode {
	return nil // TODO
}

func (parser *Parser) parseStringConst() StringConstNode {
	parser.nextTokenCheck(TokStringConst)
	return StringConstNode{Value: parser.curTok.StrVal}
}

func (parser *Parser) parseIntConst() IntConstNode {
	parser.nextTokenCheck(TokStringConst)
	return IntConstNode{Value: parser.curTok.IntVal}
}

func (parser *Parser) parseRealConst() RealConstNode {
	parser.nextTokenCheck(TokStringConst)
	return RealConstNode{Value: parser.curTok.RealVal}
}

func (parser *Parser) parseBoolConst() BoolConstNode {
	parser.nextTokenCheck(TokStringConst)
	return BoolConstNode{Value: parser.curTok.BoolVal}
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
*/
