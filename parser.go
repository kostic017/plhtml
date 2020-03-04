package main

import (
	"fmt"

	"./logging"
)

type Parser struct {
	scanner       *Scanner
	curTok        Token
	logger        *logging.MyLogger
	opPrecedences map[TokenType]int
}

func NewParser(scanner *Scanner) *Parser {
	parser := new(Parser)
	parser.scanner = scanner

	parser.logger = logging.New("PARSER")
	parser.logger.SetLevel(logging.Info)

	parser.setPrecedences([]TokenType{
		TokLtOp,
		TokGtOp,
		TokLeqOp,
		TokGeqOp,
		TokEqOp,
		TokNeqOp,
		TokenType('+'),
		TokenType('-'),
		TokenType('*'),
		TokenType('/'),
	})

	return parser
}

func (parser *Parser) setPrecedences(operators []TokenType) {
	parser.opPrecedences = make(map[TokenType]int)
	for i, v := range operators {
		parser.opPrecedences[v] = i
	}
}

func (parser *Parser) getOpPrecedence(operator TokenType) int {

}

func (parser *Parser) Parse() {
	parser.parseProgram()
}

func (parser *Parser) checkNextToken(expected TokenType) TokenType {
	actual := parser.nextToken().Type
	if actual != expected {
		panic(fmt.Sprintf("'%s' expected.", expected))
	}
	return actual
}

func (parser *Parser) checkNextTokenMore(expected ...TokenType) TokenType {
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
	parser.checkNextToken(TokenType('<'))
	parser.checkNextToken(expected)
	parser.checkNextToken(TokenType('>'))
}

func (parser *Parser) parseCloseTag(expected TokenType) {
	parser.checkNextToken(TokenType('<'))
	parser.checkNextToken(TokenType('/'))
	parser.checkNextToken(expected)
	parser.checkNextToken(TokenType('>'))
}

func (parser Parser) parseProgram() ProgramNode {
	parser.parseDoctype()
	return parser.parseHTML()
}

func (parser *Parser) parseDoctype() {
	parser.checkNextToken(TokenType('<'))
	parser.checkNextToken(TokenType('!'))
	parser.checkNextToken(TokDoctype)
	parser.checkNextToken(TokHTML)
	parser.checkNextToken(TokenType('>'))
}

func (parser *Parser) parseHTML() ProgramNode {
	parser.checkNextToken(TokenType('<'))
	parser.checkNextToken(TokHTML)
	parser.checkNextToken(TokLang)
	parser.checkNextToken(TokenType('='))
	parser.checkNextToken(TokenType('"'))
	parser.parseIdentifier()
	parser.checkNextToken(TokenType('"'))
	parser.checkNextToken(TokenType('>'))
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
	parser.checkNextToken(TokenType('<'))

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
	parser.checkNextToken(TokClass)
	parser.checkNextToken(TokenType('='))
	parser.checkNextToken(TokenType('"'))
	varType := parser.checkNextTokenMore(TokIntType, TokRealType, TokBoolType, TokStringType)
	parser.checkNextToken(TokenType('"'))
	parser.checkNextToken(TokenType('>'))
	identifier := parser.parseIdentifier()
	parser.parseCloseTag(TokVar)
	return VarDeclNode{Type: varType, Identifier: identifier}
}

func (parser *Parser) parseVarAssign() VarAssignNode {
	parser.checkNextToken(TokValue)
	parser.checkNextToken(TokenType('='))
	parser.checkNextToken(TokenType('"'))
	value := parser.parseExpression()
	parser.checkNextToken(TokenType('"'))
	parser.checkNextToken(TokenType('>'))
	identifier := parser.parseIdentifier()
	parser.parseCloseTag(TokData)
	return VarAssignNode{Identifier: identifier, Value: value}
}

func (parser *Parser) parseWriteStmt() WriteStmtNode {
	parser.checkNextToken(TokenType('>'))
	value := parser.parseExpression()
	parser.parseCloseTag(TokOutput)
	return WriteStmtNode{Value: value}
}

func (parser *Parser) parseReadStmt() ReadStmtNode {
	parser.checkNextToken(TokName)
	parser.checkNextToken(TokenType('='))
	parser.checkNextToken(TokenType('"'))
	identifier := parser.parseIdentifier()
	parser.checkNextToken(TokenType('"'))
	parser.checkNextToken(TokenType('>'))
	return ReadStmtNode{Identifier: identifier}
}

func (parser *Parser) parseControlFlowStmt() StatementNode {
	// TODO if, if-else
	parser.checkNextToken(TokData)
	parser.checkNextToken(TokenType('-'))
	parser.checkNextToken(TokWhile)
	parser.checkNextToken(TokenType('='))
	parser.checkNextToken(TokenType('"'))
	condition := parser.parseExpression()
	parser.checkNextToken(TokenType('"'))
	parser.checkNextToken(TokenType('>'))
	statements := parser.parseStatements()
	parser.parseCloseTag(TokDiv)
	return WhileStmtNode{Condition: condition, Statements: statements}
}

func (parser *Parser) parseIdentifier() IdentifierNode {
	parser.checkNextToken(TokIdentifier)
	return IdentifierNode{Name: parser.curTok.StrVal}
}

func (parser *Parser) parseExpression() ExpressionNode {
	// expr1 := parser.parseExpression()
	// op := parser.checkNextTokenMore()
	// expr2 := parser.parseExpression()
}

func (parser *Parser) parsePrimaryExpression() ExpressionNode {

	switch parser.nextToken().Type {
	case TokIntConst:
		return parser.parseIntConst()
	case TokRealConst:
		return parser.parseRealConst()
	case TokBoolConst:
		return parser.parseBoolConst()
	case TokStringConst:
		return parser.parseStringConst()
	case TokIdentifier:
		return parser.parseIdentifier()
	case TokenType('('):
		expr := parser.parseExpression()
		parser.checkNextToken(TokenType(')'))
		return expr
	case TokenType('-'), TokenType('!'):
		return parser.parseUnaryExpression(token)
	}

	panic("Invalid expression.")
}

func (parser *Parser) parseStringConst() StringConstNode {
	parser.checkNextToken(TokStringConst)
	return StringConstNode{Value: parser.curTok.StrVal}
}

func (parser *Parser) parseIntConst() IntConstNode {
	parser.checkNextToken(TokStringConst)
	return IntConstNode{Value: parser.curTok.IntVal}
}

func (parser *Parser) parseRealConst() RealConstNode {
	parser.checkNextToken(TokStringConst)
	return RealConstNode{Value: parser.curTok.RealVal}
}

func (parser *Parser) parseBoolConst() BoolConstNode {
	parser.checkNextToken(TokStringConst)
	return BoolConstNode{Value: parser.curTok.BoolVal}
}
