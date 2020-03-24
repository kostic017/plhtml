package parser

import (
	"../ast"
	"../scanner"
)

func (parser *Parser) parseStatements() []ast.StatementNode {
	myLogger.Debug("=BEG= Statements")

	var stmts []ast.StatementNode
	for stmt := parser.parseStatement(); stmt != nil; stmt = parser.parseStatement() {
		stmts = append(stmts, stmt)
	}

	myLogger.Debug("=END= Statements")
	return stmts
}

func (parser *Parser) parseStatement() ast.StatementNode {
	parser.eat(TokenType('<'))

	switch parser.peek().Type {
	case scanner.TokVar:
		return parser.parseVarDecl()
	case scanner.TokData:
		return parser.parseVarAssign()
	case scanner.TokOutput:
		return parser.parseWriteStmt()
	case scanner.TokInput:
		return parser.parseReadStmt()
	case scanner.TokDiv:
		return parser.parseControlFlowStmt()
	}

	myLogger.Debug("'%s' not a statement", string(parser.peek().Type))
	parser.goBack() // '<'
	return nil
}

func (parser *Parser) parseVarDecl() ast.VarDeclNode {
	myLogger.Debug("=BEG= Var Declaration")
	parser.eat(scanner.TokVar)
	parser.eat(scanner.TokClass)
	parser.eat(TokenType('='))
	parser.eat(TokenType('"'))
	varType := parser.parseIdentifier()
	parser.eat(TokenType('"'))
	parser.eat(TokenType('>'))
	varName := parser.parseIdentifier()
	parser.parseCloseTag(scanner.TokVar)
	myLogger.Debug("=END= Var Declaration")
	return ast.VarDeclNode{Type: varType, Identifier: varName}
}

func (parser *Parser) parseVarAssign() ast.VarAssignNode {
	myLogger.Debug("=BEG= Var Assignment")
	parser.eat(scanner.TokData)
	parser.eat(scanner.TokValue)
	parser.eat(TokenType('='))
	parser.eat(TokenType('"'))
	value := parser.parseExpr()
	parser.eat(TokenType('"'))
	parser.eat(TokenType('>'))
	identifier := parser.parseIdentifier()
	parser.parseCloseTag(scanner.TokData)
	myLogger.Debug("=END= Var Assignment")
	return ast.VarAssignNode{Identifier: identifier, Value: value}
}

func (parser *Parser) parseWriteStmt() ast.WriteStmtNode {
	myLogger.Debug("=BEG= Write")
	parser.eat(scanner.TokOutput)
	parser.eat(TokenType('>'))
	value := parser.parseExpr()
	parser.parseCloseTag(scanner.TokOutput)
	myLogger.Debug("=END= Write")
	return ast.WriteStmtNode{Value: value}
}

func (parser *Parser) parseReadStmt() ast.ReadStmtNode {
	myLogger.Debug("=BEG= Read")
	parser.eat(scanner.TokInput)
	parser.eat(scanner.TokName)
	parser.eat(TokenType('='))
	parser.eat(TokenType('"'))
	identifier := parser.parseIdentifier()
	parser.eat(TokenType('"'))
	parser.eat(TokenType('>'))
	myLogger.Debug("=END= Read")
	return ast.ReadStmtNode{Identifier: identifier}
}

func (parser *Parser) parseControlFlowStmt() ast.ControlFlowStmtNode {
	myLogger.Debug("=BEG= Control Flow")
	parser.eat(scanner.TokDiv)
	parser.eat(scanner.TokData)
	parser.eat(TokenType('-'))
	stmtType := parser.eat(scanner.TokIf, scanner.TokWhile)
	parser.eat(TokenType('='))
	parser.eat(TokenType('"'))
	condition := parser.parseExpr()
	parser.eat(TokenType('"'))
	parser.eat(TokenType('>'))
	stmts := parser.parseStatements()
	parser.parseCloseTag(scanner.TokDiv)
	myLogger.Debug("=END= Control Flow")
	return ast.ControlFlowStmtNode{Type: stmtType, Condition: condition, Statements: stmts}
}
